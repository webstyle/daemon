package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func GetDb() (*sqlx.DB, error) {
	var err error

	if db != nil {
		return db, nil
	}

	db, err = sqlx.Connect("postgres", "user=bot password=qwerty dbname=bot")

	return db, err
}
