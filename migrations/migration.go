package migrations

import (
	"github.com/leducgiachoang/app-gorm/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB){
	db.AutoMigrate(
		models.File{},
		models.Todo{},
	)
}