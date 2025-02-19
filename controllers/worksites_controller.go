package controllers

import (
	"errors"
	"net/http"
	"safecap_backend/API"    // Import the API package
	"safecap_backend/models" // Import the API package
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllWorksites(c *gin.Context) {
	var worksites []models.Worksite
	var count int64

	// Leggi i parametri di ordinamento dalla query string
	sortBy := c.DefaultQuery("sortBy", "id") // Campo di default: "id"
	order := c.DefaultQuery("order", "asc")  // Ordine di default: "asc"

	// Verifica che l'ordine sia valido (asc o desc)
	if order != "asc" && order != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Use 'asc' or 'desc'."})
		return
	}

	// Ottieni i dati dal database ordinati
	if err := db.Find(&worksites).Order(sortBy + " " + order).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Restituisci i dati ordinati
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
	var count int64

	// Leggi i parametri di ordinamento dalla query string
	sortBy := c.DefaultQuery("sortBy", "id") // Campo di default: "id"
	order := c.DefaultQuery("order", "asc")  // Ordine di default: "asc"

	// Verifica che l'ordine sia valido (asc o desc)
	if order != "asc" && order != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Use 'asc' or 'desc'."})
		return
	}

	if err := db.Preload("Worker").Where("worksite_id = ?", worksiteId).Find(&worksite_worker_assignments).Order(sortBy + " " + order).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, assignment := range worksite_worker_assignments {
		workers = append(workers, assignment.Worker)
	}

	c.JSON(http.StatusOK, gin.H{"total": count, "workers": workers})
}

func GetWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksite-id")
	var readings []models.Reading

	if err := db.Table("readings").
		Select("readings.*").
		Joins("JOIN worker_attendances ON worker_attendances.ID = readings.attendance_id").
		Where("worker_attendances.worksite_id = ?", worksiteId).
		Find(&readings).Error; err != nil {
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

	if err := db.Where("worker_id = ? AND worksite_id = ?", assignment.WorkerID, assignment.WorksiteID).First(&assignment).Error; err != nil {
		// Creazione dell'assegnazione nel database
		if err := db.Create(&assignment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Worker already assigned to worksite"})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func GetWorksiteAttendance(c *gin.Context) {
	worksiteId := c.Param("worksite-id")
	var attendance []models.WorkerAttendance

	if err := db.Where("worksite_id = ?", worksiteId).Find(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
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

	if err := db.Last(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	API.WeatherAPI(db, strconv.FormatFloat(worksite.Latitude, 'f', 6, 64), strconv.FormatFloat(worksite.Longitude, 'f', 6, 64), worksite.ID)

	c.JSON(http.StatusOK, worksite)
}

func UpdateWorksite(c *gin.Context) {
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

	if err := c.ShouldBindJSON(&worksite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksite)
}

func DeleteWorksite(c *gin.Context) {
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

	if err := db.Delete(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Worksite deleted"})
}
