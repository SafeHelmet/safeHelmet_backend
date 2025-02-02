package API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type AirPollutionResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	List []struct {
		Main struct {
			Aqi int `json:"aqi"`
		} `json:"main"`
		Components struct {
			Co   float64 `json:"co"`
			Pm10 float64 `json:"pm10"`
		} `json:"components"`
		Dt int64 `json:"dt"`
	} `json:"list"`
}

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

func weatherAPI(lat, lon string) {
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

}

func airPollutionAPI(lat, lon string) {
	log.Println("PollutionAPI: ")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/air_pollution?lat=%s&lon=%s&appid=%s", lat, lon, os.Getenv("WEATHER_API_KEY"))

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
	var airPollutionData AirPollutionResponse
	if err := json.Unmarshal(body, &airPollutionData); err != nil {
		log.Printf("Errore durante il parsing del JSON: %v\n", err)
		return
	}

	// Filtra i valori di co e pm10
	if len(airPollutionData.List) > 0 {
		co := airPollutionData.List[0].Components.Co
		pm10 := airPollutionData.List[0].Components.Pm10
		log.Printf("Valori filtrati - CO: %f, PM10: %f\n", co, pm10) // Le misure sono in μg/m^3
		// Inserisci i valori nel database o fai qualcos'altro con i dati qui
	} else {
		log.Println("Nessun dato disponibile nella risposta API")
	}
}

// TODO: Fare coroutine che questo è tutto sburato
func StartAPICallScheduler() {
	// Esegui la chiamata API all'avvio
	weatherAPI("45", "9")
	airPollutionAPI("45", "9")

	for {
		// Calcola il tempo fino alla prossima ora esatta
		now := time.Now()
		nextHour := now.Truncate(time.Hour).Add(time.Hour) // Arrotonda all'ora successiva
		durationUntilNextHour := time.Until(nextHour)

		fmt.Println("Attesa fino alla prossima ora:", durationUntilNextHour)

		// Aspetta fino alla prossima ora esatta
		time.Sleep(durationUntilNextHour)

		// Esegui la chiamata API
		weatherAPI("45", "9")
		airPollutionAPI("45", "9")
	}
}

// TODO dare a parki assieme alle letture anche le soglie
