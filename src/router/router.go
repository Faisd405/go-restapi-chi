package router

import (
	"net/http"

	examplerouter "github.com/faisd405/go-restapi-chi/src/app/example"
	homecontroller "github.com/faisd405/go-restapi-chi/src/app/home/controller"

	customMiddleware "github.com/faisd405/go-restapi-chi/src/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(customMiddleware.MiddlewareOne)
	// r.Use(customMiddleware.MiddlewareTwo)
	// r.Use(customMiddleware.MiddlewareLogging)

	r.Get("/", homecontroller.Index)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/example", examplerouter.ExampleRouter())
	})

	return r
}
