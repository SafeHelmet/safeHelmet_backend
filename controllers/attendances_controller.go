package controllers

import (
	"errors"
	"net/http"
	"safecap_backend/models"
	"strconv"
	"time"

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

func CheckAttendanceExistance(c *gin.Context) {
	worker_id := c.Param("worker_id")
	worksite_id := c.Param("worksite_id")
	helmet_id := c.Param("helmet_id")

	var attendance models.WorkerAttendance

	// Trovo l'ultima entry per worker_id, worksite_id, helmet_id
	if err := db.Where("worker_id = ? AND worksite_id = ? AND helmet_id = ?", worker_id, worksite_id, helmet_id).
		Order("start_at DESC").First(&attendance).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Controllo se la entry trovata è di oggi
	if attendance.ID != 0 {
		startYear, startMonth, startDay := attendance.StartAt.Date()
		nowYear, nowMonth, nowDay := time.Now().Date()

		if startYear == nowYear && startMonth == nowMonth && startDay == nowDay {
			c.JSON(http.StatusOK, gin.H{"attendance": attendance})
			return
		}
	}

	// Se non esiste una entry di oggi, ne creo una nuova
	var newAttendance models.WorkerAttendance
	workerID, err := strconv.Atoi(worker_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worker_id"})
		return
	}
	worksiteID, err := strconv.Atoi(worksite_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid worksite_id"})
		return
	}
	helmetID, err := strconv.Atoi(helmet_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid helmet_id"})
		return
	}

	newAttendance.WorkerID = workerID
	newAttendance.WorksiteID = worksiteID
	newAttendance.HelmetID = helmetID

	if err := db.Create(&newAttendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"attendance": newAttendance})
}
