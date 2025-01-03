package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllReadings(c *gin.Context) {
	var readings []models.Reading

	if err := db.Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
