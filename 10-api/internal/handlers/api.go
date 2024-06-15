package handlers

import (
	"10-api/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) { //uppercase functions can be exported. lowercase ones can't
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		//Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
