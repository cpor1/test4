package database

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Birth      int64  `json:"birth"`
	Created    int64  `json:"created"`
	Updated_at int64  `json:"updated_at"`
}

// bai 1 cau 2: insert du lieu user
func (db *Db) InsertUser(u User) {
	affected, err := db.engine.Insert(&u)
	fmt.Println(affected, err)
}

// bai 1 cau 2: update du lieu user
func (db *Db) UpdateUser(u User) {
	now := time.Now().UnixNano()
	id := u.Id
	affected, err := db.engine.Where("id = ?", id).Update(&User{Name: u.Name, Birth: u.Birth, Updated_at: now})
	fmt.Println(affected, err)
}

// bai 1 cau 2: list danh sach User
func (db *Db) ListUser() {
	var users []User
	err := db.engine.Find(&users)
	fmt.Println(err, users)
}

// bai 1 cau 2: đọc user theo id
func (db *Db) DetailUser(id string) (*User, error) {
	user := &User{Id: id}
	has, err := db.engine.Table(&user).Where("id = ?", id).Get(&user)
	if err != nil {
		log.Println("Fail")
		return nil, err
	}
	if !has {
		return nil, errors.New("Not Found")
	}
	return user, nil
}

//bai 1 cau 2:
func (db *Db) InsertUserAndPoint(u User) {
	affected, err := db.engine.Insert(&u)
	p := Point{u.Id, 10, 12}
	db.InsertPoint(p)
	fmt.Println(affected, err)
}

//update birth user
func (db *Db) UpdateBirth(u User) error {
	now := time.Now().UnixNano()
	id := u.Id
	_, err := db.engine.Where("id = ?", id).Update(&User{Name: u.Name + " Updated", Birth: u.Birth, Updated_at: now})
	return err
	// add Begin() before any action
}

// them diem vao user
func (db *Db) AddPointAfterUpdate(u User) error {
	var points int64
	affected, _ := db.engine.Table(&Point{}).Where("user_id = ?", u.Id).Cols("points").Get(&points)
	fmt.Println(affected)
	_, err := db.engine.Where("user_id = ?", u.Id).Update(&Point{Points: points + 10})
	return err
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

	now := time.Now().UnixNano()
	_, err = session.Where("id = ?", id).Update(&User{Name: user.Name + " Updated", Birth: birth, Updated_at: now})
	if err != nil {
		session.Rollback()
		fmt.Println("Update user birth fail")
		return err
	}

	var points int64
	_, err = session.Table(&Point{}).Where("user_id = ?", user.Id).Cols("points").Get(&points)
	if err != nil {
		session.Rollback()
		return err
	}
	_, err = session.Where("user_id = ?", user.Id).Update(&Point{Points: points + 10})
	if err != nil {
		session.Rollback()
		fmt.Println("Update point fail")
		return err
	}

	session.Commit()
	return nil
}

type dataUser struct {
	identity int
	user     User
}

func (db *Db) ScanByRow() error {
	buffScanData := make(chan *dataUser, 10)
	defer close(buffScanData)

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		go printData(buffScanData, &wg)
	}

	rows, err := db.engine.Rows(&User{})

	log.Println(err)

	defer rows.Close()

	user := new(User)

	i := 0
	for rows.Next() {
		rows.Scan(user)
		dUser := &dataUser{user: *user, identity: i}
		i++
		buffScanData <- dUser
		wg.Add(1)
	}
	wg.Wait()
	return nil
}

func printData(jobs chan *dataUser, wg *sync.WaitGroup) {

	for {
		select {
		case data := <-jobs:
			fmt.Printf("%v - %v - %v", data.identity, data.user.Id, data.user.Name)
			wg.Done()
		}
	}
}
