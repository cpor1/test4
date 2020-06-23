package database

import (
	"fmt"

	"xorm.io/xorm"
)

type Db struct {
	engine *xorm.Engine
}

func (db *Db) ConnectDb()  {
	var err error
	db.engine, err = xorm.NewEngine("mysql", "root:@/m1test?charset=utf8")
	if err != nil {
		fmt.Println("false")
	} else {
		fmt.Println("success")
	}
}
//bai 1 cau 1: tao dababase su dung struct
func (db *Db) CreateTable() {
	db.engine.CreateTables(&User{})
	db.engine.CreateTables(&Point{})
}

func (db *Db) Sync2Table() {
	db.engine.Sync2(&User{})
	db.engine.Sync2(&Point{})
}
