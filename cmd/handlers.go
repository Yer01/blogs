package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}
	render(w, *tmpl)
	fmt.Fprintf(w, "homepage")
}

func (app *application) allView(w http.ResponseWriter, r *http.Request) {
	res, err := app.blogs.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("encode response err: %v", err)
	}
}

func (app *application) singleView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
	}

	res, err := app.blogs.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("encode response err: %v", err)
	}
}

func (app *application) blogCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	content := r.FormValue("content")
	if name == "" || content == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
	}
	res, err := app.blogs.Insert(name, content)
	if err != nil {
		http.Error(w, "Duplicate blog name", http.StatusConflict)
	}
	fmt.Printf("Created successfully under id: %d\n", res)
	//http.Redirect(w, r, fmt.Sprintf("/blogs/%d", res), http.StatusSeeOther)
}
