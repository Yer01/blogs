package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	router := routes()
	logger := log.Default()

	connStr := "postgres://postgres:moura9300@localhost/blogproject?sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Print("starting server on port 8081...")
	if err := http.ListenAndServe("localhost:8081", router); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
