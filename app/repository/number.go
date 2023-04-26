package repository

import (
	"checkout-backend/app/model"

	"gorm.io/gorm"
)

// NumberRepository godoc
type NumberRepository struct {
	Db *gorm.DB
}

// InitNumber godoc
func InitNumber(db *gorm.DB) *NumberRepository {
	return &NumberRepository{Db: db}
}

// GetNumber godoc
func (service *NumberRepository) GetNumber() (int, error) {
	n := model.Number{}
	err := service.Db.First(&n).Error
	return n.Count, err
}
