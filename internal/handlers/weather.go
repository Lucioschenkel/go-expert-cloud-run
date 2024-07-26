package handlers

import (
	"encoding/json"
	"net/http"

	usecase "go-expert-cloud-run/internal/use_case"
	"go-expert-cloud-run/pkg/weather"

	"github.com/go-chi/chi"
)

type CurrentWeatherHandler struct {
	WeatherApiKey string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCurrentWeatherHandler(weatherApiKey string) *CurrentWeatherHandler {
	return &CurrentWeatherHandler{
		WeatherApiKey: weatherApiKey,
	}
}

func (c *CurrentWeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	weatherApiClient := weather.NewWeatherClient(c.WeatherApiKey)

	useCase := usecase.NewCurrentWeatherUseCase(*weatherApiClient)

	cep := chi.URLParam(r, "cep")

	currentWeather, err := useCase.Execute(usecase.Input{
		Cep: cep,
	})

	var statusCode int
	if err == usecase.ErrCannotFindZipCode {
		statusCode = http.StatusNotFound
	} else if err == usecase.ErrInvalidZipCode {
		statusCode = http.StatusBadRequest
	} else {
		statusCode = http.StatusInternalServerError
	}

	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currentWeather)
}
