package config

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://safehelmet.github.io/safeHelmet_ng/"}, // Domini consentiti
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(authMiddleware())

	return r
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ottieni il token dall'header Authorization
		token := c.GetHeader("Authorization")

		if token == "" {
			// Risposta se il token è mancante
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token required",
			})
			c.Abort() // Blocca ulteriori elaborazioni
			return
		}

		// Logica per verificare la validità del token (esempio semplice)
		if token != os.Getenv("AUTH_TOKEN") { // Sostituisci con la tua logica di verifica
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization token",
			})
			c.Abort() // Blocca ulteriori elaborazioni
			return
		}

		// Se il token è valido, consenti la richiesta
		c.Next()
	}
}
