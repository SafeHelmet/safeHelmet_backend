package controllers

import (
	"net/http"
	"safecap_backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllWorksites(c *gin.Context) {
	var worksites []models.Worksite
	var count int64

	if err := db.Find(&worksites).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "worksites": worksites})
}

func GetWorksiteDetails(c *gin.Context) {
	worksiteIdStr := c.Param("worksite-id")

	worksiteId, err := strconv.Atoi(worksiteIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worksite ID"})
		return
	}
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
	worksiteId := c.Param("worksite-id")
	var workers []models.Worker

	var worksite_worker_assignments []models.WorksiteWorkerAssignment

	if err := db.Preload("Worker").Where("worksite_id = ?", worksiteId).Find(&worksite_worker_assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, assignment := range worksite_worker_assignments {
		workers = append(workers, assignment.Worker)
	}

	c.JSON(http.StatusOK, workers)
}

func GetWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksite-id")
	var readings []models.Reading

	if err := db.Where("worksite_id = ?", worksiteId).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

func AssignWorkerToWorksite(c *gin.Context) {
	var assignment models.WorksiteWorkerAssignment

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

func CreateWorksite(c *gin.Context) {
	var worksite models.Worksite

	// Binding dei dati JSON della richiesta alla struttura
	if err := c.ShouldBindJSON(&worksite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creazione del cantiere nel database
	if err := db.Create(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksite)
}
