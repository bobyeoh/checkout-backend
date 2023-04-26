package service

import (
	"checkout-backend/app/model"

	"gorm.io/gorm"
)

// NumberService godoc
type NumberService struct {
	Db *gorm.DB
}

// InitNumber godoc
func InitNumber(db *gorm.DB) *NumberService {
	return &NumberService{Db: db}
}

// Update godoc
func (service *NumberService) Update(number int) error {
	n := model.Number{}
	if err := service.Db.First(&n).Error; err != nil {
		return err
	}
	n.Count = number
	return service.Db.Save(&n).Error
}
