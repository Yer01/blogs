package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", app.home)
	mux.Get("/blogs", app.allView)
	mux.Get("/blogs/{id}", app.singleView)
	mux.Post("/blogs/create", app.blogCreate)

	return mux
}
