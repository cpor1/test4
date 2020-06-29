package database

import "fmt"

type Point struct {
	User_id    string `json:"user_id"`
	Points     int64  `json:"points"`
	Max_points int64  `json:"max_points"`
}

func (db *Db) InsertPoint(p Point) {
	affected, err := db.engine.Insert(&p)
	fmt.Println(affected, err)
}

func (db *Db) ListPoint() error {
	var points []Point
	err := db.engine.Find(&points)
	return err
}
