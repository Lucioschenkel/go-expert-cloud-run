package main

import (
	"net/http"
	"os"

	"go-expert-cloud-run/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var weatherApiKey string

func init() {
	weatherApiKey = os.Getenv("API_KEY")
}

func main() {
	currentWeatherHandler := handlers.NewCurrentWeatherHandler(weatherApiKey)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/weather/{cep}", currentWeatherHandler.Handle)
	})

	http.ListenAndServe(":8080", r)
}
