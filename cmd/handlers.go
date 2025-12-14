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
		return
	}
	render(w, *tmpl)
	fmt.Fprintf(w, "homepage")
}

func (app *application) allView(w http.ResponseWriter, r *http.Request) {
	res, err := app.blogs.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, err.Error(), http.StatusConflict)
		return
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
		http.Error(w, "Missing fields", http.StatusConflict)
	}
	res, err := app.blogs.Insert(name, content)
	if err != nil {
		http.Error(w, "Duplicate blog name", http.StatusConflict)
		return
	}
	fmt.Printf("Created successfully under id: %d\n", res)
	//http.Redirect(w, r, fmt.Sprintf("/blogs/%d", res), http.StatusSeeOther)
}

func (app *application) blogUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
	}
	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "Missing fields", http.StatusConflict)
	}
	res, err := app.blogs.Update(id, content)

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	fmt.Printf("Blog with id %d updated succesfully!\n", res)
}

func (app *application) blogDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
	}

	if err = app.blogs.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	fmt.Printf("Blog with id %d deleted succesfully!\n", id)
}
