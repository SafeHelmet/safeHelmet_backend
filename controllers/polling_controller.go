package controllers

import (
	"net/http"
	"safecap_backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckRecentAnomaly(c *gin.Context) {
	var readings []models.Reading
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

	// Controlla se esiste almeno una rilevazione con un timestamp recente
	if err := db.Where("created_at > ?", fiveMinutesAgo).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se ci sono letture recenti, esiste un'anomalia
	exists := len(readings) > 0

	// Restituisci il risultato
	c.JSON(http.StatusOK, gin.H{"anomaly_detected": exists})
}
