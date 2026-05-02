package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type WeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func getWeather(city string) (*WeatherResponse, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY not set")
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("city not found or API error (status: %d)", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &weather, nil
}
