package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
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

func GetHelmetId(c *gin.Context) {
	helmetUuid := c.Param("helmet-uuid")
	var helmet models.Helmet

	if err := db.Where("uuid = ?", helmetUuid).First(&helmet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"helmet_id": helmet.ID})
}
