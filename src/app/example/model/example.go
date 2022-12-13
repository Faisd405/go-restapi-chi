package model

type Example struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Example1 string `gorm:"type:varchar(191)" json:"example1" validate:"required"`
	Example2 string `gorm:"type:text" json:"example2" validate:"required"`
}
