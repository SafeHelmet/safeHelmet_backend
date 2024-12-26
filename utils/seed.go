package utils

import (
	"log"
	"safecap_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDatabase(db *gorm.DB) error {

	// Crea record di esempio per la tabella Worksite
	worksites := []models.Worksite{
		{Name: "Worksite 1"},
		{Name: "Worksite 2"},
		{Name: "Worksite 3"},
		{Name: "Worksite 4"},
		{Name: "Worksite 5"},
		{Name: "Worksite 6"},
		{Name: "Worksite 7"},
		{Name: "Worksite 8"},
		{Name: "Worksite 9"},
		{Name: "Worksite 10"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksites).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Worker
	workers := []models.Worker{
		{Name: "John", Surname: "Doe"},
		{Name: "Jane", Surname: "Smith"},
		{Name: "Alice", Surname: "Johnson"},
		{Name: "Bob", Surname: "Brown"},
		{Name: "Charlie", Surname: "Davis"},
		{Name: "David", Surname: "Wilson"},
		{Name: "Eve", Surname: "Taylor"},
		{Name: "Frank", Surname: "Anderson"},
		{Name: "Grace", Surname: "Thomas"},
		{Name: "Hank", Surname: "Jackson"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&workers).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Specialization
	specializations := []models.Specialization{
		{Name: "Electrician"},
		{Name: "Plumber"},
		{Name: "Carpenter"},
		{Name: "Mason"},
		{Name: "Painter"},
		{Name: "Welder"},
		{Name: "Mechanic"},
		{Name: "Roofer"},
		{Name: "Plasterer"},
		{Name: "Glazier"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&specializations).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorkerSpecialization
	workerSpecializations := []models.WorkerSpecialization{
		{WorkerID: 1, SpecializationID: 1},
		{WorkerID: 2, SpecializationID: 2},
		{WorkerID: 3, SpecializationID: 3},
		{WorkerID: 4, SpecializationID: 4},
		{WorkerID: 5, SpecializationID: 5},
		{WorkerID: 6, SpecializationID: 6},
		{WorkerID: 7, SpecializationID: 7},
		{WorkerID: 8, SpecializationID: 8},
		{WorkerID: 9, SpecializationID: 9},
		{WorkerID: 10, SpecializationID: 10},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&workerSpecializations).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella HelmetCategory
	helmetCategories := []models.HelmetCategory{
		{Name: "Standard"},
		{Name: "Advanced"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&helmetCategories).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Helmet
	helmets := []models.Helmet{
		{CategoryID: 1},
		{CategoryID: 2},
		{CategoryID: 1},
		{CategoryID: 2},
		{CategoryID: 1},
		{CategoryID: 2},
		{CategoryID: 1},
		{CategoryID: 2},
		{CategoryID: 1},
		{CategoryID: 2},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&helmets).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Reading
	readings := []models.Reading{
		{HelmetID: 1, Sensor1: 10.5, Sensor2: 20.3, Sensor3: 30.1, Anomalous: false},
		{HelmetID: 2, Sensor1: 11.5, Sensor2: 21.3, Sensor3: 31.1, Anomalous: true},
		{HelmetID: 3, Sensor1: 12.5, Sensor2: 22.3, Sensor3: 32.1, Anomalous: false},
		{HelmetID: 4, Sensor1: 13.5, Sensor2: 23.3, Sensor3: 33.1, Anomalous: true},
		{HelmetID: 5, Sensor1: 14.5, Sensor2: 24.3, Sensor3: 34.1, Anomalous: false},
		{HelmetID: 6, Sensor1: 15.5, Sensor2: 25.3, Sensor3: 35.1, Anomalous: true},
		{HelmetID: 7, Sensor1: 16.5, Sensor2: 26.3, Sensor3: 36.1, Anomalous: false},
		{HelmetID: 8, Sensor1: 17.5, Sensor2: 27.3, Sensor3: 37.1, Anomalous: true},
		{HelmetID: 9, Sensor1: 18.5, Sensor2: 28.3, Sensor3: 38.1, Anomalous: false},
		{HelmetID: 10, Sensor1: 19.5, Sensor2: 29.3, Sensor3: 39.1, Anomalous: true},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&readings).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorkerWorksiteAssignment
	workerWorksiteAssignments := []models.WorkerWorksiteAssignment{
		{WorkerID: 1, WorksiteID: 1, HelmetID: 1},
		{WorkerID: 2, WorksiteID: 2, HelmetID: 2},
		{WorkerID: 3, WorksiteID: 3, HelmetID: 3},
		{WorkerID: 4, WorksiteID: 4, HelmetID: 4},
		{WorkerID: 5, WorksiteID: 5, HelmetID: 5},
		{WorkerID: 6, WorksiteID: 6, HelmetID: 6},
		{WorkerID: 7, WorksiteID: 7, HelmetID: 7},
		{WorkerID: 8, WorksiteID: 8, HelmetID: 8},
		{WorkerID: 9, WorksiteID: 9, HelmetID: 9},
		{WorkerID: 10, WorksiteID: 10, HelmetID: 10},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&workerWorksiteAssignments).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorksiteBossAssignment
	worksiteBossAssignments := []models.WorksiteBossAssignment{
		{BossID: 1, WorksiteID: 1},
		{BossID: 2, WorksiteID: 2},
		{BossID: 3, WorksiteID: 3},
		{BossID: 4, WorksiteID: 4},
		{BossID: 5, WorksiteID: 5},
		{BossID: 6, WorksiteID: 6},
		{BossID: 7, WorksiteID: 7},
		{BossID: 8, WorksiteID: 8},
		{BossID: 9, WorksiteID: 9},
		{BossID: 10, WorksiteID: 10},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksiteBossAssignments).Error; err != nil {
		return err
	}

	log.Println("Database seeded successfully")
	return nil
}
