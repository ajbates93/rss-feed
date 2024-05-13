package main

import (
	"ajbates93/rss-feed/pkg/handlers"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	fmt.Println("Initialising the database connection...")
	db, err := InitDB()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database connection initialised successfully!")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	feedHandler := handlers.RSSFeedHandler{DB: db}
	feedHandler.RegisterRoutes(r)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
