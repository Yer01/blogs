package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	authH := AuthHandler{
		Username: os.Getenv("AUTH_USERNAME"),
		Password: os.Getenv("AUTH_PASSWORD"),
	}

	mux.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	mux.Get("/", app.home)
	mux.Get("/blogs", app.allView)
	mux.Get("/blogs/create", authH.auth(app.blogCreateView))
	mux.Post("/blogs/create", authH.auth(app.blogCreate))
	mux.Get("/blogs/{id}/edit", authH.auth(app.blogUpdateView))
	mux.Post("/blogs/{id}/edit", authH.auth(app.blogUpdate))
	mux.Get("/blogs/{id}", app.singleView)
	mux.Put("/blogs/{id}", authH.auth(app.blogUpdate))
	mux.Delete("/blogs/{id}", authH.auth(app.blogDelete))
	return mux
}
