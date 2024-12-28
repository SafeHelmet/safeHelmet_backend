package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllWorksites(c *gin.Context) {
	var worksites []models.Worksite

	if err := db.Find(&worksites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksites)
}

func GetWorksiteDetails(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var worksite models.Worksite

	if err := db.First(&worksite, worksiteId).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Worksite not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, worksite)
}

func GetWorkersInWorksite(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var workers []models.Worker

	if err := db.Where("worksite_id = ?", worksiteId).Find(&workers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workers)
}

func GetWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var readings []models.Reading

	if err := db.Where("worksite_id = ?", worksiteId).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}
