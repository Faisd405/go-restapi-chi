package models

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	Example1 string `json:"example1" gorm:"type:varchar(191)" validate:"required"`
	Example2 string `json:"example2" gorm:"type:text" validate:"required"`
}

func NewExample(example1 string, example2 string) Example {
	return Example{
		Example1: example1,
		Example2: example2,
	}
}
