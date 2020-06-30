package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Db struct {
	engine *xorm.Engine
}

var tables []interface{}

func (db *Db) ConnectDb() error {
	var err error
	db.engine, err = xorm.NewEngine("mysql", "root:@/m1test?charset=utf8")
	if err != nil {
		fmt.Println("false")
	} else {
		fmt.Println("success")
	}
	return err
}

//bai 1 cau 1: tao dababase su dung struct
func (db *Db) CreateTable() {
	db.engine.CreateTables(&User{})
	db.engine.CreateTables(&Point{})
}

func (db *Db) Sync2Table() error {
	tables = append(tables, new(User), new(Point))
	err := db.engine.Sync2(tables...)
	return err
}
