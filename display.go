package main

import (
	"fmt"
	"strings"
)

func getWeatherEmoji(description string) string {
	desc := strings.ToLower(description)

	switch {
	case strings.Contains(desc, "clear"):
		return "☀️"
	case strings.Contains(desc, "cloud"):
		return "☁️"
	case strings.Contains(desc, "rain"):
		return "🌧️"
	case strings.Contains(desc, "drizzle"):
		return "🌦️"
	case strings.Contains(desc, "thunder"):
		return "⛈️"
	case strings.Contains(desc, "snow"):
		return "❄️"
	case strings.Contains(desc, "mist"), strings.Contains(desc, "fog"):
		return "🌫️"
	case strings.Contains(desc, "haze"):
		return "😶‍🌫️"
	default:
		return "🌡️"
	}
}

func displayWeather(w *WeatherResponse) {
	emoji := getWeatherEmoji(w.Weather[0].Description)
	description := strings.Title(w.Weather[0].Description)
	windKmh := w.Wind.Speed * 3.6

	border := "─────────────────────────────────"

	fmt.Println()
	fmt.Println(border)
	fmt.Printf("  📍 %s, %s\n", w.Name, w.Sys.Country)
	fmt.Println(border)
	fmt.Printf("  %s  %s\n", emoji, description)
	fmt.Println()
	fmt.Printf("  🌡️  Temperature  : %.1f°C\n", w.Main.Temp)
	fmt.Printf("  🤔  Feels Like   : %.1f°C\n", w.Main.FeelsLike)
	fmt.Printf("  💧  Humidity     : %d%%\n", w.Main.Humidity)
	fmt.Printf("  💨  Wind Speed   : %.1f km/h\n", windKmh)
	fmt.Println(border)
	fmt.Println()
}
