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

func b1() error {
	now := time.Now().UnixNano()
	db.ConnectDb()
	db.Sync2Table()

	//insert
	user := database.User{"1", "tungedit", 1592456481, now, now}
	db.InsertUser(user)

	//update
	user2 := database.User{"1", "tungtest", 11111111110, 159245648121312, now}
	db.UpdateUser(user2)

	// list
	db.ListUser()

	// detail by id
	user3, _ := db.DetailUser("1")
	fmt.Println(user3)

	//  tạo user thì insert user_id vào user_point với số điểm 10.
	db.InsertUserAndPoint(user)
	db.ListPoint()
	return nil
}

func b2() error {
	db.ConnectDb()
	err := db.SessionTest("1", 1231231)
	if err != nil {
		fmt.Fprintln(nil)
	}
	return nil
}

func b3() error {
	db.ConnectDb()
	now := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		user := database.User{"k" + strconv.Itoa(i), "user " + strconv.Itoa(i), now, now, now}
		db.InsertUser(user)
	}
	err := db.ScanByRow()
	return err
}

func UsingCli() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "cli",
				Usage: "using cli",
			},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "b1",
			Usage: "create table, insert, select",
			Action: func(c *cli.Context) error {
				return b1()
			},
		},
		{
			Name:  "b2",
			Usage: "update birth then add point",
			Action: func(c *cli.Context) error {
				return b2()
			},
		},
		{
			Name:  "b3",
			Usage: "add 100 user then scan row",
			Action: func(c *cli.Context) error {
				return b3()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
