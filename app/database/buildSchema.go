package database

import (
	"checkout-backend/app/model"

	"gorm.io/gorm"
)

// BuildSchema godoc
func BuildSchema(DB *gorm.DB) {
	if !DB.Migrator().HasTable(&model.Number{}) {
		DB.AutoMigrate(&model.Number{})
		initNumber := model.Number{Count: 0}
		DB.Create(&initNumber)
	}
}
