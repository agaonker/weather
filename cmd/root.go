package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

var (
	lat string
	lon string
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "A weather CLI application",
	Long:  `A command line application to get weather information for any location`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file:", err)
			return
		}

		// Get the environment variable
		apiKey := os.Getenv("OPENWEATHER_API_KEY")

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, apiKey)

		// Make the API request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making API request:", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Parse JSON into our struct
		var weather WeatherResponse
		err = json.Unmarshal(body, &weather)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		// Print the weather information
		fmt.Printf("\nWeather for %s:\n", weather.Name)
		fmt.Printf("-------------------\n")
		fmt.Printf("Current Temperature: %.1f°C\n", weather.Main.Temp)
		fmt.Printf("Feels Like: %.1f°C\n", weather.Main.FeelsLike)
		fmt.Printf("Wind Speed: %.1f m/s\n", weather.Wind.Speed)
		fmt.Printf("-------------------\n")
	},
}

func init() {
	rootCmd.Flags().StringVarP(&lat, "latitude", "a", "37.7799", "Latitude of the location")
	rootCmd.Flags().StringVarP(&lon, "longitude", "o", "-121.9780", "Longitude of the location")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
