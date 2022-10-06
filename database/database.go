package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var Conn = initDatabase()
