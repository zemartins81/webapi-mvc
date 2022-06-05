package migrations

import (
	"github.com/jcmartins81/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{})
	if err != nil {
		return
	}
}
