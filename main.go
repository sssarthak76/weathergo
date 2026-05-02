package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: weathergo <city>")
		fmt.Println("Example: weathergo Bengaluru")
		os.Exit(1)
	}

	city := os.Args[1]

	weather, err := getWeather(city)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	displayWeather(weather)
}
