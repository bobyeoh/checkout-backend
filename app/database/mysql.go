package database

import (
	"checkout-backend/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMYSQL godoc
func InitMYSQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetConfig().MYSQL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.Logger.LogMode(logger.Info)
	Build(db)
	return db
}

// Build godoc
func Build(DB *gorm.DB) {
	BuildSchema(DB)
}
