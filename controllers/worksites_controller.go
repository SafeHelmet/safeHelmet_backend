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

func AssignWorkerToWorksite(c *gin.Context) {
	var assignment models.WorkerWorksiteAssignment

	// Binding dei dati JSON della richiesta alla struttura
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creazione dell'assegnazione nel database
	if err := db.Create(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignment)
}
