package database

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type User struct {
	Id         string
	Name       string
	Birth      int64
	Created    int64
	Updated_at int64
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
func (db *Db) DetailUser() {
	user := User{}
	id := "2"
	ishas, err := db.engine.Table(&user).Where("id = ?", id).Get(&user)
	fmt.Println(ishas, err, user)
}

//bai 1 cau 2:
func (db *Db) InsertUserAndPoint(u User) {
	affected, err := db.engine.Insert(&u)
	db.InsertPoint(u)
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
func (db *Db) SessionTest(u User) error {
	session := db.engine.NewSession()
	defer session.Close()

	session.Begin()
	err := db.UpdateBirth(u)
	if err != nil {
		session.Rollback()
		return err
	}
	err = db.AddPointAfterUpdate(u)
	if err != nil {
		session.Rollback()
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
