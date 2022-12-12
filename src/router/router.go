package router

import (
	"net/http"

	examplecontroller "github.com/faisd405/go-restapi-chi/src/app/example/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", helloWorld)

	r.Route("/api", func(r chi.Router) {
		r.Route("/example", func(r chi.Router) {
			r.Get("/", examplecontroller.Index)
			r.Post("/", examplecontroller.Create)
			r.Get("/{id}", examplecontroller.Show)
			r.Put("/{id}", examplecontroller.Update)
			r.Delete("/{id}", examplecontroller.Delete)
		})
	})

	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}