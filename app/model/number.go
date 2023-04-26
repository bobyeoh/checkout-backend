package model

import (
	"gorm.io/gorm"
)

// Number godoc
type Number struct {
	gorm.Model
	Count int `gorm:"type:int"`
}
