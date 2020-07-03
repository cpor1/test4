package database

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Birth     int64  `json:"birth"`
	Created   int64  `json:"created"`
	UpdatedAt int64  `json:"updated_at"`
}

type DataUser struct {
	Identity int
	User     User
}

// bai 1 cau 2: insert du lieu user
func (db *Db) InsertUser(u User) error {
	affected, err := db.engine.Insert(&u)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("can not insert")
	}
	return err
}

// bai 1 cau 2: update du lieu user
func (db *Db) UpdateUser(user, conditions *User) error {
	aff, err := db.engine.Update(user, conditions)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("cannot update")
	}
	return nil
}

// bai 1 cau 2: list danh sach User
func (db *Db) ListUser() ([]*User, error) {
	users := make([]*User, 0)
	err := db.engine.Find(&users)
	if err != nil {
		return nil, err
	}
	return users, err
}

// bai 1 cau 2: đọc user theo id
func (db *Db) DetailUser(id string) (*User, error) {
	user := &User{Id: id}
	has, err := db.engine.Get(user)
	if err != nil {
		log.Println("Fail")
		return nil, err
	}
	if !has {
		return nil, errors.New("Not Found")
	}
	return user, err
}

//cau 3
func (db *Db) ScanByRow(buffChannel chan *DataUser, wg *sync.WaitGroup) error {
	rows, err := db.engine.Rows(&User{})
	if err != nil {
		return err
	}
	defer rows.Close()
	user := new(User)
	i := 1
	for rows.Next() {
		err = rows.Scan(user)
		if err != nil {
			return err
		} else {
			dataUser := &DataUser{Identity: i, User: *user}
			buffChannel <- dataUser
			wg.Add(1)
			i++
		}
	}
	return nil
}

//cau 2
func (db *Db) SessionTest(id string, birth int64) error {
	session := db.engine.NewSession()
	defer session.Close()
	session.Begin()

	user := &User{Id: id}
	has, err := session.Get(user)
	if err != nil {
		session.Rollback()
		log.Println("Fail")
		return err
	}
	if !has {
		session.Rollback()
		log.Println("Not found user")
		return err
	}
	user.Name = user.Name + " Updated"
	user.Birth = birth
	user.UpdatedAt = time.Now().UnixNano()
	_, err = session.Update(user, &User{Id: id})
	if err != nil {
		session.Rollback()
		fmt.Println("Update user birth fail")
		return err
	}

	point := &Point{UserId: user.Id}
	_, err = session.Get(point)
	if err != nil {
		session.Rollback()
		return err
	}
	point.Points = point.Points + 10
	_, err = session.Update(point, &Point{UserId: user.Id})
	if err != nil {
		session.Rollback()
		fmt.Println("Update point fail")
		return err
	}

	session.Commit()
	return nil
}
