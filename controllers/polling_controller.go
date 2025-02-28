package controllers

import (
	"net/http"
	"safecap_backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckRecentAnomaly(c *gin.Context) {
	var readings []models.Reading
	helmetId := c.Param("helmet-id")
	oneAndHalfMinutesAgo := time.Now().Add(-90 * time.Second)

	// Trova il casco specifico
	var helmet models.Helmet
	if err := db.First(&helmet, helmetId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trova il cantiere associato al casco
	var worksite models.Worksite
	if err := db.Joins("JOIN worker_attendances ON worker_attendances.worksite_id = worksites.id").
		Where("worker_attendances.helmet_id = ?", helmetId).
		First(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Controlla se esiste almeno una rilevazione con un timestamp recente nel cantiere specifico
	if err := db.Joins("JOIN worker_attendances ON worker_attendances.ID = readings.attendance_id").
		Where("worker_attendances.worksite_id = ? AND readings.read_at > ? AND readings.anomaly = TRUE", worksite.ID, oneAndHalfMinutesAgo).
		Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Restituisci il risultato
	c.JSON(http.StatusOK, gin.H{"Anomalies": len(readings)})
}
