package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherApiClient struct {
	apiKey string
}

func NewWeatherClient(apiKey string) *WeatherApiClient {
	return &WeatherApiClient{
		apiKey: apiKey,
	}
}

type LocationResponse struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	TzID    string  `json:"tz_id"`
}

type CurrentWeatherResponse struct {
	LastUpdated string  `json:"last_updated"`
	TempC       float64 `json:"temp_c"`
	TempF       float64 `json:"temp_f"`
}

type WeatherAPIResponse struct {
	Location LocationResponse       `json:"location"`
	Current  CurrentWeatherResponse `json:"current"`
}

func (c *WeatherApiClient) GetWeatherForCity(cityName string) (*WeatherAPIResponse, error) {
	url := fmt.Sprintf(
		"http://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		c.apiKey,
		url.QueryEscape(cityName),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherQueryResult WeatherAPIResponse
	err = json.Unmarshal(body, &weatherQueryResult)
	if err != nil {
		return nil, err
	}

	return &weatherQueryResult, nil
}
