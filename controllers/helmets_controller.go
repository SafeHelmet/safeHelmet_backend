package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllHelmets(c *gin.Context) {
	var helmets []models.Helmet
	var count int64

	// Leggi i parametri di ordinamento dalla query string
	sortBy := c.DefaultQuery("sortBy", "id") // Campo di default: "id"
	order := c.DefaultQuery("order", "asc")  // Ordine di default: "asc"

	// Verifica che l'ordine sia valido (asc o desc)
	if order != "asc" && order != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Use 'asc' or 'desc'."})
		return
	}

	if err := db.Find(&helmets).Order(sortBy + " " + order).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"total": count, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helmets)
}

func GetHelmetDetails(c *gin.Context) {
	helmetId := c.Param("helmet-id")
	var helmet models.Helmet

	if err := db.First(&helmet, helmetId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helmet)
}

// Ritorna l'ID del casco per l'api del polling e verifica quando mi connetto ad un casco che esso effettivamente esista nel DB
func GetHelmetId(c *gin.Context) {
	helmetMacAddress := c.Param("mac-address")
	var helmet models.Helmet

	if err := db.Where("mac_address = ?", helmetMacAddress).First(&helmet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"helmet_id": helmet.ID})
}

func GetHelmetAttendance(c *gin.Context) {
	helmetId := c.Param("helmet-id")
	var attendance []models.WorkerAttendance

	if err := db.Where("helmet_id = ?", helmetId).Find(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func GetHelmetReadings(c *gin.Context) {
	helmetId := c.Param("helmet-id")
	var readings []models.Reading

	if err := db.Table("readings").
		Select("readings.*").
		Joins("JOIN worker_attendances ON worker_attendances.ID = readings.attendance_id").
		Where("worker_attendances.helmet_id = ?", helmetId).
		Find(&readings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, readings)
}

func GetHelmetCategories(c *gin.Context) {
	var categories []models.HelmetCategory

	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetHelmetCategoryDetails(c *gin.Context) {
	categoryId := c.Param("category-id")
	var category models.HelmetCategory

	if err := db.First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateHelmet(c *gin.Context) {
	helmetId := c.Param("helmet-id")
	var helmet models.Helmet

	if err := db.First(&helmet, helmetId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Helmet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&helmet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&helmet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helmet)
}

func CreateHelmet(c *gin.Context) {
	var helmet models.Helmet

	if err := c.ShouldBindJSON(&helmet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&helmet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, helmet)
}

func DeleteHelmet(c *gin.Context) {
	helmetId := c.Param("worker-id")
	var helmet models.Helmet

	if err := db.First(&helmet, helmetId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Helmet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&helmet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Helmet deleted"})
}
