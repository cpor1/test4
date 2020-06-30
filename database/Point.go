package database

import (
	"errors"
	"log"
)

type Point struct {
	UserId    string `json:"user_id"`
	Points    int64  `json:"points"`
	MaxPoints int64  `json:"max_points"`
}

func (db *Db) InsertPoint(p Point) error {
	affected, err := db.engine.Insert(&p)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("can not insert")
	}
	return err
}

func (db *Db) ListPoint() ([]*Point, error) {
	points := make([]*Point, 0)
	err := db.engine.Find(&points)
	if err != nil {
		return nil, err
	}
	return points, err
}

func (db *Db) DetailPoint(user_id string) (*Point, error) {
	point := &Point{UserId: user_id}
	has, err := db.engine.Get(point)
	if err != nil {
		log.Println("Fail")
		return nil, err
	}
	if !has {
		return nil, errors.New("Not Found")
	}
	return point, err
}
