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
	var reading models.Reading

	// Read the data from the context
	var requestData map[string]interface{}

	/// TODO: non è detto che vada bene magari posso fare il bind diretto di reading?
	// Bind del JSON ricevuto nel body della richiesta alla mappa
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Get worksite ID from helmetID
	var worksite models.Worksite
	if err := db.Joins("JOIN worker_attendances ON worker_attendances.worksite_id = worksites.id").
		Where("worker_attendances.helmet_id = ?", requestData.helmet).
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
	reading.anomaly = false

	// Temperature anomaly check
	if requestData.temp-weather.TempMax > 10 || weather.TempMin-requestData.temp > 10 {
		reading.anomaly = true
	}

	// Humidity anomaly check
	if requestData.humidity > weather.Humidity+15 {
		reading.anomaly = true
	}

	// Brightness anomaly check
	if requestData.brightness > weather.Brightness+100 && !requestData.UsesWeldingProtection {
		reading.anomaly = true
	}

	// Gas anomaly check
	if (requestData.Methane || requestData.CarbonMonoxide || requestData.SmokeDetection) && !requestData.UsesGasProtection {
		reading.anomaly = true
	}

	// Posture anomaly check
	if requestData.IncorrectPosture > 0.5 {
		reading.anomaly = true
	}

	// Crash anomaly check
	if requestData.Max_G > 10 {
		reading.anomaly = true
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
