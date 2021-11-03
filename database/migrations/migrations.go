package migrations

import (
	"web-api-go/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
