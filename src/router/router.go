package router

import (
	"net/http"

	authRouter "github.com/faisd405/go-restapi-chi/src/app/auth"
	exampleRouter "github.com/faisd405/go-restapi-chi/src/app/example"
	homecontroller "github.com/faisd405/go-restapi-chi/src/app/home/controller"

	customMiddleware "github.com/faisd405/go-restapi-chi/src/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(customMiddleware.Cors())
	// r.Use(customMiddleware.MiddlewareOne)
	// r.Use(customMiddleware.MiddlewareTwo)
	// r.Use(customMiddleware.MiddlewareLogging)

	r.Get("/", homecontroller.Index)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/example", exampleRouter.ExampleRouter())
		r.Mount("/auth", authRouter.AuthRouter())
	})

	return r
}
