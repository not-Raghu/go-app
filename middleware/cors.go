package middleware

import (
	"github.com/rs/cors"
)

func CorsConfig() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "example.com"},
		// AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	})
	return c
}
