package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, tmpl *template.Template, data templateData) {
	buf := new(bytes.Buffer)

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
