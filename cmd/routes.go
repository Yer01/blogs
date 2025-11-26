package main

import (
	"github.com/go-chi/chi/v5"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", home)
	mux.Get("/blogs", blogs)
	mux.Get("/blogs/{id}", blogView)

	return mux
}
