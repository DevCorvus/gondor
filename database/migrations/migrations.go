package migrations

import (
	"github.com/DevCorvus/gondor/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
