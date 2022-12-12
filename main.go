package main

import (
	"fmt"
	"net/http"

	"github.com/faisd405/go-restapi-chi/src/config"
	"github.com/faisd405/go-restapi-chi/src/router"
)

func main() {
	config.ConnectDatabase()

	r := router.Router()
	fmt.Println("Server started on: http://127.0.0.1:3000")
	http.ListenAndServe("127.0.0.1:3000", r)
}
