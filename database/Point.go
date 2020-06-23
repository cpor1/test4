package database

import "fmt"

type Point struct {
	User_id    string
	Points     int64
	Max_points int64
}

func (db *Db) InsertPoint(u User) {
	p := Point{u.Id, 10, 12}
	affected, err := db.engine.Insert(&p)
	fmt.Println(affected, err)
}

func (db *Db) ListPoint() {
	var points []Point
	err := db.engine.Find(&points)
	fmt.Println(err, points)
}
