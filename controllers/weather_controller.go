package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllWeatherReadings(c *gin.Context) {
	var weathers []models.WorkerAttendance

	if err := db.Find(&weathers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weathers)
}

func GetAllWorksiteWeather(c *gin.Context) {
	var worksiteId = c.Param("worksite-id")

	var weathers []models.WorkerAttendance

	if err := db.Where("worksite = ?", worksiteId).Find(&weathers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weathers)
}

func GetLastWorksiteWeather(c *gin.Context) {
	var worksiteId = c.Param("worksite-id")

	var weather models.WorkerAttendance

	if err := db.Where("worksite = ?", worksiteId).Last(&weather).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weather)
}
