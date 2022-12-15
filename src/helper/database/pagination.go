package database

import (
	"strconv"

	"gorm.io/gorm"
)

func Paginate(params map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit := 10
		currentPage := 1
		if params["limit"] != "" {
			limit, _ = strconv.Atoi(params["limit"])
		}
		if params["currentPage"] != "" {
			currentPage, _ = strconv.Atoi(params["currentPage"])
		}
		offset := (currentPage - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}
