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

func b2() error {
	db.ConnectDb()
	err := db.SessionTest("1", 1231231)
	if err != nil {
		panic(err)
	}
	return nil
}

func b3() error {
	db.ConnectDb()
	now := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		user := database.User{"k" + strconv.Itoa(i), "user " + strconv.Itoa(i), now, now, now}
		err := db.InsertUser(user)
		if err != nil {
			panic(err)
		}
	}
	err := db.ScanByRow()
	if err != nil {
		panic(err)
	}
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
