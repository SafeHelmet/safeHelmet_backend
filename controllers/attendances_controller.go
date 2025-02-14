package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLastAttendanceDetails(c *gin.Context) {
	worker_id := c.Param("worker_id")
	worksite_id := c.Param("worksite_id")
	helmet_id := c.Param("helmet_id")

	var attendance models.WorkerAttendance

	// Trovo l'ultima entry per worker_id, worksite_id, helmet_id
	if err := db.Where("worker_id = ? AND worksite_id = ? AND helmet_id = ?", worker_id, worksite_id, helmet_id).
		Last(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func GetAttendanceDetails(c *gin.Context) {
	attendanceId := c.Param("attendance_id")
	var attendance models.WorkerAttendance

	if err := db.First(&attendance, attendanceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attendance not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func GetAllAttendances(c *gin.Context) {
	var attendances []models.WorkerAttendance

	if err := db.Find(&attendances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendances)
}
