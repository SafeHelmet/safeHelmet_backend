package controllers

import (
	"net/http"
	"safecap_backend/models"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetAllBosses(c *gin.Context) {
	var bosses []models.Boss

	if err := db.Find(&bosses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bosses)
}

func GetBossDetails(c *gin.Context) {
	id := c.Param("boss-id")
	var boss models.Boss

	if err := db.Where("id = ?", id).First(&boss).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, boss)
}

func CreateBoss(c *gin.Context) {
	var boss models.Boss

	if err := c.ShouldBindJSON(&boss); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boss.CreatedAt = time.Now()
	boss.UpdatedAt = time.Now()

	if err := db.Create(&boss).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, boss)
}

func UpdateBoss(c *gin.Context) {
	id := c.Param("boss-id")
	var boss models.Boss

	if err := db.First(&boss, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Boss not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&boss); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boss.UpdatedAt = time.Now()

	if err := db.Save(&boss).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, boss)
}

func DeleteBoss(c *gin.Context) {
	id := c.Param("boss-id")
	var boss models.Boss

	if err := db.Where("id = ?", id).First(&boss).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&boss).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Boss deleted"})
}
