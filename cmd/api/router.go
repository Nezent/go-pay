package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors" // Import CORS middleware
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Logs requests
	r.Use(middleware.Recoverer) // Recovers from panics

	// CORS Configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache for 5 minutes
	}))

	// Base API path
	r.Route("/api/v1", func(r chi.Router) {
		// r.Get("/users", handler.GetUsers)
		// r.Post("/payments", handler.CreatePayment)
	})

	return r
}
