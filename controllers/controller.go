package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

// Inizializza il database
func InitDatabase(database *gorm.DB) {
	db = database
}

func GetAllWorksites(c *gin.Context) {
	var worksites []models.Worksite

	if err := db.Find(&worksites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksites)
}

func GetAllWorkers(c *gin.Context) {
	var workers []models.Worker

	if err := db.Find(&workers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workers)
}

func GetAllHelmets(c *gin.Context) {
	var helmets []models.Helmet

	if err := db.Find(&helmets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helmets)
}

func GetAllHelmetCategories(c *gin.Context) {
	var categories []models.HelmetCategory

	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetAllSpecializations(c *gin.Context) {
	var specializations []models.Specialization

	if err := db.Find(&specializations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specializations)
}

func GetAllReadings(c *gin.Context) {
	var readings []models.Reading

	if err := db.Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

func GetAllWorkerWorksiteAssignments(c *gin.Context) {
	var assignments []models.WorkerWorksiteAssignment

	if err := db.Find(&assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func GetAllWorksiteBossAssignments(c *gin.Context) {
	var assignments []models.WorksiteBossAssignment

	if err := db.Find(&assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

// Elenco dei Lavoratori in un Cantiere
func GetWorkers(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var workers []models.Worker

	if err := db.Where("worksite_id = ?", worksiteId).Find(&workers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workers)
}

// Dettagli di un Lavoratore
func GetWorkerDetails(c *gin.Context) {
	workerId := c.Param("workerId")
	var worker models.Worker

	if err := db.First(&worker, workerId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, worker)
}

// Letture di un Lavoratore in un Cantiere
func GetReadings(c *gin.Context) {
	workerId := c.Param("workerId")
	worksiteId := c.Param("worksiteId")
	var readings []models.Reading

	if err := db.Where("worker_id = ? AND worksite_id = ?", workerId, worksiteId).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

// Letture di un Cantiere
func GetWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var readings []models.Reading

	if err := db.Where("worksite_id = ?", worksiteId).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

// Letture Anomale di un Cantiere
func GetAnomalousReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var readings []models.Reading

	if err := db.Where("worksite_id = ? AND anomalous = ?", worksiteId, true).Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

// Assegna un Lavoratore a un Cantiere
func AssignWorkerToWorksite(c *gin.Context) {
	var assignment models.WorkerWorksiteAssignment

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignment)
}
