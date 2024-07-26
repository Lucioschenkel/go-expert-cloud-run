package usecase

import (
	"errors"

	"go-expert-cloud-run/pkg/cep"
	"go-expert-cloud-run/pkg/weather"
)

type CurrentWeatherUseCase struct {
	weatherApiClient weather.WeatherApiClient
}

type CurrentWeather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type Input struct {
	Cep string
}

var ErrInvalidZipCode = errors.New("invalid zip code")
var ErrCannotFindZipCode = errors.New("can not find zip code")
var ErrUnknowError = errors.New("unknown error")

func NewCurrentWeatherUseCase(weatherApiClient weather.WeatherApiClient) *CurrentWeatherUseCase {
	return &CurrentWeatherUseCase{
		weatherApiClient: weatherApiClient,
	}
}

func (c *CurrentWeatherUseCase) Execute(input Input) (*CurrentWeather, error) {
	isCepValid, err := cep.Validate(input.Cep)
	if err != nil || !isCepValid {
		return nil, ErrInvalidZipCode
	}

	location, err := cep.GetLocationFromCep(input.Cep)
	if err != nil {
		return nil, ErrCannotFindZipCode
	}

	currentWeather, err := c.weatherApiClient.GetWeatherForCity(location.Localidade)
	if err != nil {
		return nil, ErrUnknowError
	}

	return &CurrentWeather{
		TempC: currentWeather.Current.TempC,
		TempF: weather.CelsiusToFarenheit(currentWeather.Current.TempC),
		TempK: weather.CelsiusToKelvin(currentWeather.Current.TempC),
	}, nil
}
