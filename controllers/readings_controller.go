package controllers

import (
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
		Joins("JOIN readings ON readings.helmet_id = worker_attendances.helmet_id").
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
		Joins("JOIN readings ON readings.helmet_id = worker_attendances.helmet_id").
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

// / TODO: Implementare la funzione per creare la reading dai dati del mobile
func CreateReading(c *gin.Context) {

	const (
		TempAnomalyThreshold       = 10
		HumidityAnomalyThreshold   = 15
		BrightnessAnomalyThreshold = 100
		IncorrectPostureThreshold  = 0.5
		MaxGThreshold              = 10
	)

	var reading models.Reading

	// Bind del JSON ricevuto nel body della richiesta alla mappa
	if err := c.ShouldBindJSON(&reading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Get worksite ID from helmetID
	var worksite models.Worksite
	if err := db.Joins("JOIN worker_attendances ON worker_attendances.worksite_id = worksites.id").
		Where("worker_attendances.helmet_id = ?", reading.HelmetID).
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

	// Check if there are any anomalies
	reading.Anomaly = false

	// Temperature Anomaly check
	if reading.Temperature-weather.TempMax > TempAnomalyThreshold || weather.TempMin-reading.Temperature > TempAnomalyThreshold {
		reading.Anomaly = true
	}

	// Humidity Anomaly check
	if reading.Humidity > weather.Humidity+HumidityAnomalyThreshold {
		reading.Anomaly = true
	}

	// Brightness Anomaly check
	if reading.Brightness > weather.Brightness+BrightnessAnomalyThreshold && !reading.UsesWeldingProtection {
		reading.Anomaly = true
	}

	// Gas Anomaly check
	if (reading.Methane || reading.CarbonMonoxide || reading.SmokeDetection) && !reading.UsesGasProtection {
		reading.Anomaly = true
	}

	// Posture Anomaly check
	if reading.IncorrectPosture > IncorrectPostureThreshold {
		reading.Anomaly = true
	}

	// Crash Anomaly check
	if reading.Max_G > MaxGThreshold {
		reading.Anomaly = true
	}

	if err := db.Create(&reading).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reading)
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
