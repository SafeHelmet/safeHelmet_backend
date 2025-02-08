package controllers

import (
	"net/http"
	"safecap_backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllWorkers(c *gin.Context) {
	var workers []models.Worker
	var count int64

	// Leggi i parametri di ordinamento dalla query string
	sortBy := c.DefaultQuery("sortBy", "id") // Campo di default: "id"
	order := c.DefaultQuery("order", "asc")  // Ordine di default: "asc"

	// Verifica che l'ordine sia valido (asc o desc)
	if order != "asc" && order != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Use 'asc' or 'desc'."})
		return
	}

	if err := db.Find(&workers).Order(sortBy + " " + order).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"total": count, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workers)
}

func GetWorkerDetails(c *gin.Context) {
	workerId := c.Param("worker-id")
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

func GetWorksiteOfWorker(c *gin.Context) {
	workerId := c.Param("worker-id")
	var worksites []models.Worksite

	if err := db.Table("worksites").
		Select("worksites.id, worksites.name"). // Corretto il typo
		Joins("JOIN worksite_worker_assignments ON worksite_worker_assignments.worksite_id = worksites.id").
		Where("worksite_worker_assignments.worker_id = ?", workerId).
		Find(&worksites). // Usato Find al posto di Scan
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"worksites": worksites})
}

func UpdateWorker(c *gin.Context) {
	workerId := c.Param("worker-id")
	var worker models.Worker

	if err := db.First(&worker, workerId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	worker.UpdatedAt = time.Now()

	if err := db.Save(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worker)
}

func CreateWorker(c *gin.Context) {
	var worker models.Worker

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, worker)
}

func CreateWorkerAttendance(c *gin.Context) {
	var attendance models.WorkerAttendance

	// Binding del JSON ricevuto
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salvo nel DB
	if err := db.Create(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func UpdateWorkerAttendance(c *gin.Context) {
	var attendance models.WorkerAttendance

	// Binding del JSON ricevuto
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Trovo l'ultima entry per worker_id, worksite_id, helmet_id
	if err := db.Where("worker_id = ? AND worksite_id = ? AND helmet_id = ?", attendance.WorkerID, attendance.WorksiteID, attendance.HelmetID).
		Order("start_at DESC"). // Ordino per start_at in ordine decrescente
		First(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	attendance.EndAt = &now

	// Salvo nel DB
	if err := db.Save(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func DeleteWorker(c *gin.Context) {
	workerId := c.Param("worker-id")
	var worker models.Worker

	if err := db.First(&worker, workerId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Worker deleted"})
}
