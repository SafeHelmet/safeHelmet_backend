package controllers

import (
	"math"
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllReadings(c *gin.Context) {
	var readings []models.Reading
	var count int64

	// Leggi i parametri di ordinamento dalla query string
	sortBy := c.DefaultQuery("sortBy", "id") // Campo di default: "id"
	order := c.DefaultQuery("order", "asc")  // Ordine di default: "asc"

	// Verifica che l'ordine sia valido (asc o desc)
	if order != "asc" && order != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Use 'asc' or 'desc'."})
		return
	}

	if err := db.Find(&readings).Order(sortBy + " " + order).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"total": count, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

func GetReadingDetails(c *gin.Context) {
	readingId := c.Param("reading-id")
	var reading models.Reading

	if err := db.First(&reading, readingId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reading)
}

// Forse si puo fare scritta meglio in ORM non so
func GetReadingWorker(c *gin.Context) {
	readingId := c.Param("reading-id")
	var worker models.Worker

	if err := db.Table("workers").
		Select("workers.*").
		Joins("JOIN worker_attendances ON worker_attendances.worker_id = workers.id").
		Joins("JOIN readings ON readings.attendance_id = worker_attendances.ID").
		Where("readings.id = ?", readingId).
		Where("worker_attendances.end_at IS NULL"). // Considera solo chi è attualmente presente
		Scan(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worker)
}

// Forse si puo fare scritta meglio in ORM non so
func GetReadingWorksite(c *gin.Context) {
	readingId := c.Param("reading-id")
	var worksite models.Worksite

	if err := db.Table("worksites").
		Select("worksites.*").
		Joins("JOIN worker_attendances ON worker_attendances.worksite_id = worksites.id").
		Joins("JOIN readings ON readings.attendance_id = worker_attendances.ID").
		Where("readings.id = ?", readingId).
		Where("worker_attendances.end_at IS NULL"). // Considera solo chi è attualmente presente
		Scan(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksite)
}

func UpdateReading(c *gin.Context) {
	readingId := c.Param("reading-id")
	var reading models.Reading

	if err := db.First(&reading, readingId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&reading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&reading).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reading)
}

func CreateReading(c *gin.Context) {

	var reading models.Reading

	// Bind del JSON ricevuto nel body della richiesta alla mappa
	if err := c.ShouldBindJSON(&reading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Get worksite ID from helmetID
	var worksite models.Worksite
	if err := db.Table("worksites").
		Joins("JOIN worker_attendances ON worker_attendances.worksite_id = worksites.id").
		Where("worker_attendances.id = ?", reading.AttendanceID).
		First(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Worksite not found"})
		return
	}

	// Find last weather data for the worksite
	var weather models.WeatherData
	if err := db.Where("worksite_id = ?", worksite.ID).
		Order("created_at DESC").
		First(&weather).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Weather data not found"})
		return
	}

	reading.WeatherTemperatureMax = weather.TempMax
	reading.WeatherTemperatureMin = weather.TempMin
	reading.WeatherTemperature = weather.Temp
	reading.WeatherHumidity = weather.Humidity
	reading.WeatherBrightness = weather.Brightness

	// Initialize booleand values
	reading.Anomaly = false
	reading.AnomalousTemperature = false
	reading.AnomalousHumidity = false
	reading.AnomalousBrightness = false
	reading.AnomalousMaxG = false
	reading.AnomalousPosture = false

	// Temperature Anomaly check
	if reading.Temperature-weather.TempMax > worksite.TemperatureThreshold || weather.TempMin-reading.Temperature > worksite.TemperatureThreshold {
		reading.Anomaly = true
		reading.AnomalousTemperature = true
	}

	// Humidity Anomaly check
	if math.Abs(reading.Humidity-weather.Humidity) > worksite.HumidityThreshold {
		reading.Anomaly = true
		reading.AnomalousHumidity = true
	}

	// Brightness Anomaly check
	if reading.Brightness > weather.Brightness+worksite.BrightnessThreshold && !reading.UsesWeldingProtection {
		reading.Anomaly = true
		reading.AnomalousBrightness = true
	}

	// Gas Anomaly check
	if (reading.Methane || reading.CarbonMonoxide || reading.SmokeDetection) && !reading.UsesGasProtection {
		reading.Anomaly = true
	}

	// Crash Anomaly check
	if reading.Max_G > worksite.MaxGThreshold {
		reading.Anomaly = true
		reading.AnomalousMaxG = true
	}

	// Posture Anomaly check
	if reading.IncorrectPosture > worksite.PostureThreshold {
		reading.Anomaly = true
		reading.AnomalousPosture = true
	}

	if err := db.Create(&reading).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"reading": reading, "anomaly": reading.Anomaly})
}

func DeleteReading(c *gin.Context) {
	readingId := c.Param("reading-id")
	var reading models.Reading

	if err := db.First(&reading, readingId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Delete(&reading).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reading deleted"})
}
