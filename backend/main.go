package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	apiKey = "tfYOqxI3JsoAwAvljhvUcaY5omWhWK3l" // Replace with your actual API key
)

func main() {
	app := fiber.New()

	// Endpoint to fetch weather data
	app.Get("/weather", func(c *fiber.Ctx) error {
		// Get locationKey from query parameter
		locationKey := c.Query("locationKey")
		if locationKey == "" {
			return c.Status(400).SendString("Location key is required")
		}

		// Fetch weather data from AccuWeather
		url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s?apikey=%s", locationKey, apiKey)
		resp, err := http.Get(url)
		if err != nil {
			return c.Status(500).SendString("Failed to fetch weather data")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.Status(500).SendString("Failed to read weather data")
		}

		return c.SendString(string(body))
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
