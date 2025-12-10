package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Yer01/internal/model"
	_ "github.com/lib/pq"
)

type application struct {
	blogs *model.BlogModel
}

func main() {

	logger := log.Default()

	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		blogs: &model.BlogModel{DB: db},
	}

	router := app.routes()

	logger.Print("starting server on port 8081...")
	if err := http.ListenAndServe("localhost:8081", router); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
