package main

import (
	"log"
	"net/http"

	"github.com/drowsyrobotboy/inspire2/internal/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register API routes
	api.RegisterRoutes(r)

	log.Println("Starting server on 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
