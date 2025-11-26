package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func blogs(w http.ResponseWriter, r *http.Request) {

}

func blogView(w http.ResponseWriter, r *http.Request) {

}
