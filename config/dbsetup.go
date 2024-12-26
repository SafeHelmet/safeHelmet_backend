package config

import (
	"log"
	"os"
	"safecap_backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configura la connessione al database PostgreSQL
	return "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" TimeZone=" + os.Getenv("DB_TIMEZONE")
}

func Migrate(db *gorm.DB) {
	// Migrazione delle strutture
	err := db.AutoMigrate(
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
}

func ConnectToDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
