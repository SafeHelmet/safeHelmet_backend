package main

import (
	"log"
	"net/http"
	"os"
	"safecap_backend/models"
	"safecap_backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing argument: toseed")
	}
	toseed := os.Args[1] // Ottieni il valore del parametro toseed

	// Configura la connessione al database PostgreSQL
	dsn := "host=dpg-ctm3uljv2p9s73f9h470-a.frankfurt-postgres.render.com user=safehelmet_db_user password=lBmeOC0lvxjawRiRD5L1pAvRezYH8LPu dbname=safehelmet_db port=5432 TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.Exec("DROP DATABASE safehelmet_db WITH (FORCE);")

	// Migrazione delle strutture
	err = db.AutoMigrate(
		&models.Worksite{},
		&models.Worker{},
		&models.Specialization{},
		&models.WorkerSpecialization{},
		&models.Helmet{},
		&models.HelmetCategory{},
		&models.Reading{},
		&models.WorkerWorksiteAssignment{},
		&models.WorksiteBossAssignment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")

	if toseed == "true" {
		// Popola il database con dati di esempio
		err = utils.SeedDatabase(db)
		if err != nil {
			log.Fatal("Failed to seed database:", err)
		}
		log.Println("Database seeded successfully")
	}

	// Crea un router Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "TODO"}, // Domini consentiti
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Definisci gli endpoint
	r.GET("/worksites/:worksiteId/workers", getWorkers)
	r.GET("/workers/:workerId", getWorkerDetails)
	r.GET("/workers/:workerId/worksites/:worksiteId/readings", getReadings)
	r.GET("/worksites/:worksiteId/readings", getWorksiteReadings)
	r.GET("/worksites/:worksiteId/readings/anomalous", getAnomalousReadings)
	r.POST("/worksites/:worksiteId/workers", assignWorkerToWorksite)

	// test
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test successful"})
	})

	// Avvia il server sulla porta 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Elenco dei Lavoratori in un Cantiere
func getWorkers(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	// Implementa la logica per ottenere i lavoratori di un cantiere
	c.JSON(http.StatusOK, gin.H{"worksiteId": worksiteId})
}

// Dettagli di un Lavoratore
func getWorkerDetails(c *gin.Context) {
	workerId := c.Param("workerId")
	// Implementa la logica per ottenere i dettagli di un lavoratore
	c.JSON(http.StatusOK, gin.H{"workerId": workerId})
}

// Letture di un Lavoratore in un Cantiere
func getReadings(c *gin.Context) {
	workerId := c.Param("workerId")
	worksiteId := c.Param("worksiteId")
	// Implementa la logica per ottenere le letture di un lavoratore in un cantiere
	c.JSON(http.StatusOK, gin.H{"workerId": workerId, "worksiteId": worksiteId})
}

// Letture di un Cantiere
func getWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	// Implementa la logica per ottenere le letture di un cantiere
	c.JSON(http.StatusOK, gin.H{"worksiteId": worksiteId})
}

// Letture Anomale di un Cantiere
func getAnomalousReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	// Implementa la logica per ottenere le letture anomale di un cantiere
	c.JSON(http.StatusOK, gin.H{"worksiteId": worksiteId})
}

// Assegna un Lavoratore a un Cantiere
func assignWorkerToWorksite(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	// Implementa la logica per assegnare un lavoratore a un cantiere
	c.JSON(http.StatusOK, gin.H{"worksiteId": worksiteId})
}
