package database

import (
	"checkout-backend/app/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMYSQL godoc
func InitMYSQL() *gorm.DB {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.Open(config.GetConfig().MYSQL), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err.Error())
	}
	Build(db)
	return db
}

// Build godoc
func Build(DB *gorm.DB) {
	BuildSchema(DB)
}
