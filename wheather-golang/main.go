package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now().Format("15:04:05"))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("WEATHERAPI_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key required", http.StatusInternalServerError)
		return
	}
	city := "Rostov-on-Don"
	url := fmt.Sprintf(
		"http://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=ru",
		apiKey, city,
	)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Weather service error", http.StatusBadGateway)
		return
	}
	var wr WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&wr); err != nil {
		http.Error(w, "Invalid weather data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(
		w,
		"Погода в %s: %.1f°C, %s",
		wr.Location.Name,
		wr.Current.TempC,
		wr.Current.Condition.Text,
	)
}

func main() {
	http.HandleFunc("/date", dateHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/weather", weatherHandler)
	http.ListenAndServe(":8080", nil)
}
