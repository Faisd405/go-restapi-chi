package router

import (
	"net/http"

	ExampleHandler "github.com/faisd405/go-restapi-chi/src/app/example/handlers"
	ExampleRepository "github.com/faisd405/go-restapi-chi/src/app/example/repositories"
	ExampleService "github.com/faisd405/go-restapi-chi/src/app/example/services"
	"github.com/faisd405/go-restapi-chi/src/config"

	"github.com/go-chi/chi/v5"
)

func ExampleRouter() http.Handler {

	exampleRepository := ExampleRepository.NewExampleRepository(config.DB)
	exampleService := ExampleService.NewExampleService(exampleRepository)

	handlers := ExampleHandler.NewExampleHandler(exampleService)
	r := chi.NewRouter()
	r.Get("/", handlers.Index)
	r.Get("/{id}", handlers.Show)
	r.Post("/", handlers.Create)
	r.Patch("/{id:[0-9]+}", handlers.Update)
	r.Delete("/{id:[0-9]+}", handlers.Delete)

	return r
}
