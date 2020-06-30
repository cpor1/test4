package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"test4/database"
	"time"

	"github.com/urfave/cli/v2"
)

var db *database.Db = new(database.Db)

func Bai1(c *cli.Context) error {
	now := time.Now().UnixNano()
	err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	err = db.Sync2Table()
	if err != nil {
		panic(err)
	}
	//insert
	user := database.User{"1", "tungedit", 1592456481, now, now}
	err = db.InsertUser(user)
	if err != nil {
		panic(err)
	}
	//update
	user2 := database.User{"1", "tungtest", 11111111110, 159245648121312, now}
	err = db.UpdateUser(user2)
	if err != nil {
		panic(err)
	}
	// list
	err = db.ListUser()
	if err != nil {
		panic(err)
	}
	// detail by id
	user3, err := db.DetailUser("1")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(user3)
	}

	//  tạo user thì insert user_id vào user_point với số điểm 10.
	err = db.InsertUserAndPoint(user)
	if err != nil {
		panic(err)
	}
	err = db.ListPoint()
	if err != nil {
		panic(err)
	}
	return nil
}

func Bai2(c *cli.Context) error {
	err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	err = db.SessionTest("1", 1231231)
	if err != nil {
		panic(err)
	}
	return nil
}

func Bai3(c *cli.Context) error {
	err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	now := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		user := database.User{"k" + strconv.Itoa(i), "user " + strconv.Itoa(i), now, now, now}
		err := db.InsertUser(user)
		if err != nil {
			panic(err)
		}
	}
	err = db.ScanByRow()
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "cli",
				Usage: "using cli",
			},
		},
	}

	app.Commands = []*cli.Command{
		{Name: "b1", Usage: "create table, insert, select", Action: Bai1},
		{Name: "b2", Usage: "update birth then add point", Action: Bai2},
		{Name: "b3", Usage: "add 100 user then scan row", Action: Bai3},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}