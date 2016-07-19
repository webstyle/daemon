package main

import (
	"daemon/db"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn *sqlx.DB
var oldResult int
var newResult int

func init() {
	var err error

	dbConn, err = db.GetDb()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer dbConn.Close()

	log.Printf("Listing subscribers")

	forever := make(chan bool)

	go func() {
		tick := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-tick.C:
				err := checking()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	<-forever
}

func checking() (err error) {
	err = dbConn.Get(&newResult, "SELECT COUNT(*) as count FROM subscribers")
	if err != nil {
		return
	}

	if oldResult <= 0 {
		oldResult = newResult
	}

	if newResult > oldResult {
		log.Print("newResult", newResult)
		log.Print("oldResult", oldResult)
		oldResult = newResult
		log.Printf("Sending a notifaction!!!")
	}

	return
}
