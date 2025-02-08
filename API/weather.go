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

func weatherAPI(db *gorm.DB, lat, lon string) {
	log.Println("WeatherAPI: ")

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

	weather.Brightness = 5000 /// TODO: aggiungere la luminosità

	if err := db.Create(&weather).Error; err != nil {
		log.Printf("Error in weather POST: %v", err)
		return
	}

}

// TODO: Fare coroutine che questo è tutto sburato
func StartAPICallScheduler(db *gorm.DB) {
	for {
		// Calcola il tempo fino alla prossima ora esatta
		now := time.Now()
		nextHour := now.Truncate(time.Hour).Add(time.Hour) // Arrotonda all'ora successiva
		durationUntilNextHour := time.Until(nextHour)

		fmt.Println("Attesa fino alla prossima ora:", durationUntilNextHour)

		// Aspetta fino alla prossima ora esatta
		time.Sleep(durationUntilNextHour)

		// Esegui la chiamata API
		var worksites []models.Worksite

		if err := db.Find(&worksites); err != nil {
			log.Printf("Error fetching worksites: %v", err)
		}

		for _, worksite := range worksites {
			weatherAPI(db, fmt.Sprintf("%f", worksite.Latitude), fmt.Sprintf("%f", worksite.Longitude))
		}
	}
}
