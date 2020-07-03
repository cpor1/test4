package main

import (
	"fmt"
	"sync"
	"test4/database"
)

func printData(jobs chan *database.DataUser, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-jobs:
			fmt.Printf("%v - %v - %v \n", data.Identity, data.User.Id, data.User.Name)
			wg.Done()
		}
	}
}
