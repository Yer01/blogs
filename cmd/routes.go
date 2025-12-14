package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	mux.Get("/", app.home)
	mux.Get("/blogs", app.allView)
	mux.Get("/blogs/{id}", app.singleView)
	mux.Post("/blogs/create", app.blogCreate)
	mux.Put("/blogs/{id}", app.blogUpdate)
	mux.Delete("/blogs/{id}", app.blogDelete)
	return mux
}
