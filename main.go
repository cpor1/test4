package main

import (
	"fmt"
	"test4/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *database.Db = new(database.Db)
	db.ConnectDb()
	//now := time.Now().UnixNano()
	// var err error
	// engine, err = xorm.NewEngine("mysql", "root:@/m1test?charset=utf8")
	// fmt.Println(err)

	//bai 1: create table using struct
	//db.CreateTable()
	//db.Sync2Table()

	//bai 2:
	//insert
	// user := database.User{"4", "tungedit", 1592456481, now, now}
	// db.InsertUser(user)

	// update
	// user := database.User{"5", "tungtest", 11111111110, 159245648121312, now}
	// db.UpdateUser(user)

	//list
	// db.ListUser()

	// detail by id
	// user , _ := db.DetailUser("1")
	// fmt.Println(user)

	// tạo user thì insert user_id vào user_point với số điểm 10.
	//db.InsertUserAndPoint(user)
	// db.ListPoint()

	//bai 2
	err := db.SessionTest("k5", 1231231)
	if err != nil {
		fmt.Fprintln(nil)
	}
	//bai 3
	// for i := 0; i < 100; i++ {
	// 	user := database.User{"k" + strconv.Itoa(i), "user " + strconv.Itoa(i), now, now, now}
	// 	db.InsertUser(user)
	// }
	// db.ScanByRow()

}
