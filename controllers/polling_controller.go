package controllers

import (
	"log"
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

	// Ottieni l'attendance del casco
	var attendance models.WorkerAttendance
	if err := db.Where("helmet_id = ?", helmetId).Last(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ritorna le reading di altre attendances o quelle generate da te ma non dal gas
	var anomaliesToNotify []models.Reading
	for _, reading := range readings {
		if reading.AttendanceID != attendance.ID || (!reading.Methane && !reading.CarbonMonoxide && !reading.SmokeDetection) {
			anomaliesToNotify = append(anomaliesToNotify, reading)
		}
	}

	// Log per il debug
	log.Printf("Helmet ID: %s", helmetId)
	log.Printf("Worksite ID: %d", worksite.ID)
	log.Printf("Attendance ID: %d", attendance.ID)
	log.Printf("Number of readings found: %d", len(readings))
	log.Printf("Number of anomalies to notify: %d", len(anomaliesToNotify))

	// Restituisci il risultato
	c.JSON(http.StatusOK, gin.H{"Anomalies": len(anomaliesToNotify)})
}
