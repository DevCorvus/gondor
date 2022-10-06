package database

import (
	"github.com/DevCorvus/gondor/database/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migrations.Run(db)

	return db
}

var Conn = initDatabase()
