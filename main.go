package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/faisd405/go-restapi-chi/src/config"
	"github.com/faisd405/go-restapi-chi/src/router"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	argAppName = kingpin.Arg("name", "Application name").Default("Golang Restful").String()
	argPort    = kingpin.Arg("port", "Web server port").Default("3000").Int()
)

func main() {
	// Take the arguments from the command line
	kingpin.Parse()
	appName := *argAppName
	port := fmt.Sprintf(":%d", *argPort)

	confServerPort := os.Getenv("SERVER_PORT")
	if confServerPort == "" {
		confServerPort = "3000"
	}

	// Connect to the database
	config.ConnectDatabase()

	// Create a new router
	r := router.Router()
	fmt.Printf("Server %s started on: http://127.0.0.1:%s \n", appName, port)

	// Start the server
	http.ListenAndServe("127.0.0.1:3000", r)
}
