package migrations

import (
	"github.com/jcmartins81/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
