package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/base.tmpl", "templates/home.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	blogs, err := app.blogs.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Title:      "Welcome",
		Year:       time.Now().Year(),
		Blogs:      blogs,
		TotalPosts: len(blogs),
	}

	render(w, t, data)

	fmt.Fprintf(w, "homepage")
}

func (app *application) allView(w http.ResponseWriter, r *http.Request) {
	res, err := app.blogs.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := template.ParseFiles("templates/base.tmpl", "templates/blogs.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Title:      "All blogs",
		Year:       time.Now().Year(),
		Blogs:      res,
		TotalPosts: len(res),
	}

	render(w, t, data)

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

	t, err := template.ParseFiles("templates/base.tmpl", "templates/blog.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := templateData{
		Title:      fmt.Sprintf("Blog %d", id),
		Year:       time.Now().Year(),
		Blog:       &res,
		TotalPosts: 1,
	}
	render(w, t, data)
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

func (app *application) blogCreateView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.tmpl", "templates/create.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := templateData{
		Title: "Create blog",
		Year:  time.Now().Year(),
	}

	render(w, t, data)
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
