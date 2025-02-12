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
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

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

	/// TODO: controllare se funziona
	// Controlla se esiste almeno una rilevazione con un timestamp recente nel cantiere specifico
	if err := db.Joins("JOIN worker_attendances ON worker_attendances.ID = readings.attendance_id").
		Where("worker_attendances.worksite_id = ? AND readings.read_at > ? AND readings.anomaly = TRUE", worksite.ID, fiveMinutesAgo).
		Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/// TODO: magari ritornare true/false
	// Se ci sono letture recenti anomale
	exists := len(readings) > 0

	// Restituisci il risultato
	c.JSON(http.StatusOK, gin.H{"anomaly_detected": exists})
}
