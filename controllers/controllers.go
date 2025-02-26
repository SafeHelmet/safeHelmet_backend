package controllers

import (
	"net/http"

	"your_project_path/models"
	"your_project_path/validators"

	"github.com/gin-gonic/gin"
)

func CreateWorksite(c *gin.Context) {
	var worksite models.Worksite
	if err := c.ShouldBindJSON(&worksite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidateStruct(worksite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&worksite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, worksite)
}
