package API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"safecap_backend/models"
	"time"

	"gorm.io/gorm"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func WeatherAPI(db *gorm.DB, lat string, lon string, worksiteID int) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=%s&appid=%s", lat, lon, "metric", os.Getenv("WEATHER_API_KEY"))

	// Esegui la richiesta HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Errore durante la chiamata API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Leggi la risposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Errore durante la lettura della risposta: %v\n", err)
		return
	}

	// Parsing del body
	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		log.Printf("Errore durante il parsing del JSON: %v\n", err)
		return
	}

	// Filtra i valori di temp, temp_min, temp_max e humidity
	temp := weatherData.Main.Temp
	tempMin := weatherData.Main.TempMin
	tempMax := weatherData.Main.TempMax
	humidity := weatherData.Main.Humidity
	log.Printf("Valori filtrati - Temp: %f, Temp Min: %f, Temp Max: %f, Humidity: %d\n", temp, tempMin, tempMax, humidity)
	// Inserisci i valori nel database o fai qualcos'altro con i dati qui

	var weather models.WeatherData

	weather.Temp = temp
	weather.TempMin = tempMin
	weather.TempMax = tempMax
	weather.Humidity = float64(humidity)
	weather.WorksiteID = worksiteID

	weather.Brightness = 500 /// TODO: aggiungere la luminosità

	if err := db.Create(&weather).Error; err != nil {
		log.Printf("Error in weather POST: %v", err)
		return
	}

}

func StartAPICallScheduler(db *gorm.DB) {
	for {
		log.Println("WeatherAPI: ")
		now := time.Now()
		hour := now.Hour()

		// If it's before 8 or after 20, wait until 8
		if hour < 8 || hour >= 20 {
			nextRun := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
			if hour >= 20 { // Se è dopo le 20, aspetta fino a domani alle 8
				nextRun = nextRun.Add(24 * time.Hour)
			}

			waitTime := time.Until(nextRun)
			log.Println("Waiting for new work day:", nextRun)
			time.Sleep(waitTime)
			continue
		}

		var worksites []models.Worksite
		result := db.Find(&worksites)
		if result.Error != nil {
			log.Printf("Error in worksite query: %v", result.Error)
		} else {
			log.Println("Calling weather API at:", now)
			for _, worksite := range worksites {
				WeatherAPI(db, fmt.Sprintf("%f", worksite.Latitude), fmt.Sprintf("%f", worksite.Longitude), worksite.ID)
			}
		}

		// Wait for the next hour
		nextHour := now.Truncate(time.Hour).Add(time.Hour)
		sleepDuration := time.Until(nextHour)

		fmt.Println("Waiting for next hour:", nextHour)
		time.Sleep(sleepDuration)
	}
}
