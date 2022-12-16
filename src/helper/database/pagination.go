package database

import (
	"net/http"
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

func BuildParams(r *http.Request) map[string]string {
	params := map[string]string{}
	params["limit"] = "10"
	params["currentPage"] = "1"

	if r.URL.Query().Get("example1") != "" {
		params["example1"] = r.URL.Query().Get("example1")
	}
	if r.URL.Query().Get("limit") != "" {
		params["interface"] = r.URL.Query().Get("interface")
	}
	if r.URL.Query().Get("currentPage") != "" {
		params["currentPage"] = r.URL.Query().Get("currentPage")
	}

	return params
}
