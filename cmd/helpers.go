package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, tmpl template.Template) {
	buf := new(bytes.Buffer)

	err := tmpl.ExecuteTemplate(buf, "base", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	buf.WriteTo(w)
}
