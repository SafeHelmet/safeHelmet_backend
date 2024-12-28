package controllers

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Inizializza il database
func InitDatabase(database *gorm.DB) {
	db = database
}
