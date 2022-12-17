package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" gorm:"uniqueIndex" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
