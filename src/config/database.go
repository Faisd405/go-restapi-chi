package config

import (
	auth "github.com/faisd405/go-restapi-chi/src/app/auth/models"
	example "github.com/faisd405/go-restapi-chi/src/app/example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/golang?parseTime=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&example.Example{})
	database.AutoMigrate(&auth.User{})

	DB = database
}
