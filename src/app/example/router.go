package router

import (
	"net/http"

	ExampleHandler "github.com/faisd405/go-restapi-chi/src/app/example/handlers"
	ExampleRepository "github.com/faisd405/go-restapi-chi/src/app/example/repositories"
	ExampleService "github.com/faisd405/go-restapi-chi/src/app/example/services"
	"github.com/faisd405/go-restapi-chi/src/config"
	globalUtils "github.com/faisd405/go-restapi-chi/src/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ExampleRouter() http.Handler {

	exampleRepository := ExampleRepository.NewExampleRepository(config.DB)
	exampleService := ExampleService.NewExampleService(exampleRepository)

	handlers := ExampleHandler.NewExampleHandler(exampleService)
	r := chi.NewRouter()
	r.Get("/", handlers.Index)
	r.Get("/{id}", handlers.Show)

	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(globalUtils.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", handlers.Create)
		r.Patch("/{id}", handlers.Update)
		r.Delete("/{id}", handlers.Delete)
	})

	return r
}
