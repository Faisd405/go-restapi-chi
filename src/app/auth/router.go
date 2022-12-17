package router

import (
	"net/http"

	AuthHandler "github.com/faisd405/go-restapi-chi/src/app/auth/handlers"
	// AuthRepository "github.com/faisd405/go-restapi-chi/src/app/auth/repositories"
	// AuthService "github.com/faisd405/go-restapi-chi/src/app/auth/services"
	// "github.com/faisd405/go-restapi-chi/src/config"

	globalUtils "github.com/faisd405/go-restapi-chi/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func AuthRouter() http.Handler {

	// authRepository := AuthRepository.NewAuthRepository(config.DB)
	// authService := AuthService.NewAuthService(authRepository)

	// handlers := AuthHandler.NewAuthHandler(authService)
	r := chi.NewRouter()
	r.Post("/login", AuthHandler.Login)
	r.Post("/register", AuthHandler.Register)

	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(globalUtils.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/user", AuthHandler.User)
	})

	return r
}
