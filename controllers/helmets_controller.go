package controllers

import (
	"net/http"
	"safecap_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllHelmets(c *gin.Context) {
	var helmets []models.Helmet

	if err := db.Find(&helmets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helmets)
}
