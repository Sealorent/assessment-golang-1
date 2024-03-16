package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type WeatherData struct {
	Water       int    `json:"water"`        // Value water dalam satuan meter
	StatusWater string `json:"status_water"` // Value water dalam satuan meter
	Wind        int    `json:"wind"`         // Value wind dalam satuan meter per detik
	StatusWind  string `json:"status_wind"`  // Status cuaca (siaga atau bahaya)
}

func main() {
	//  running the css file
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	http.HandleFunc("/", templateHtml)
	http.HandleFunc("/refresh", refreshPage)
	http.ListenAndServe(":8080", nil)

}

func templateHtml(w http.ResponseWriter, r *http.Request) {
	// call the refreshPage function
	refreshPage(w, r)

	// HTML template
	jsonTemplate := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="refresh" content="5">
			<title>Weather</title>
			<link rel="stylesheet" type="text/css" href="/style/styles.css">
		</head>
		<body>
			<h1>Weather</h1>
			<p>Water: {{.Water}} m</p>
			<p>Status: {{.StatusWater}}</p>
			<p>Wind: {{.Wind}} m/s</p>
			<p>Status: {{.StatusWind}}</p>
		</body>
		</html>`

	// Create and parse the JSON template
	tmpl, err := template.New("weather").Parse(jsonTemplate)
	if err != nil {
		log.Fatal("Error parsing JSON template:", err)
	}

	// Read JSON file and parse data
	data, err := os.ReadFile("weather.json")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error reading JSON file:", err)
		return
	}

	var weatherData WeatherData
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing JSON data:", err)
		return
	}

	// Execute template and write HTML response
	err = tmpl.Execute(w, weatherData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}

}

func refreshPage(w http.ResponseWriter, r *http.Request) {

	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// Determine weather status
	var statusWind string
	switch {
	case wind < 6:
		statusWind = "Aman"
	case wind == 6:
		statusWind = "Diantara Aman dan Siaga"
	case wind >= 7 && wind <= 15:
		statusWind = "Siaga"
	case wind > 15:
		statusWind = "Bahaya"
	default:
		statusWind = "Tidak diketahui"
	}

	// Determine water status
	var statusWater string
	switch {
	case water <= 5:
		statusWater = "Aman"
	case wind == 5:
		statusWind = "Diantara Aman dan Siaga"
	case water >= 6 && water <= 8:
		statusWater = "Siaga"
	case water > 8:
		statusWater = "Bahaya"
	default:
		statusWater = "Tidak diketahui"
	}

	// Create WeatherData struct with random values and status
	weather := WeatherData{
		Water:       water,
		StatusWater: statusWater,
		Wind:        wind,
		StatusWind:  statusWind,
	}

	// Marshal WeatherData struct to JSON
	jsonData, err := json.Marshal(weather)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
	}

	// Write JSON data to a file
	err = os.WriteFile("weather.json", jsonData, 0644)
	if err != nil {
		log.Println("Error writing JSON to file:", err)
	}

	// log the updated weather data

	log.Println("Updated weather.json with water:", water, "m and wind:", wind, "m/s (StatusWater:", statusWater, " StatusWind:", statusWind, ")")

}
