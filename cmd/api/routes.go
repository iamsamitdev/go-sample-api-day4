package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	// Create a router mux
	mux := chi.NewRouter()

	// Add middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// register routes
	mux.Get("/", app.Home)
	mux.Get("/about", app.About)
	mux.Get("/demomovies", app.AllDemoMovies)

	// Authenticated routes
	mux.Post("/authenticate", app.authenticate)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Get("/movies", app.AllMovies)
	mux.Get("/movies/{id}", app.GetMovie)

	mux.Route("/admin", func(mux chi.Router) {

		// Protected routes
		mux.Use(app.authRequired)

		mux.Get("/movies/{id}", app.MovieForEdit)
	})

	return mux
}
