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

	// Crea record di esempio per la tabella Boss
	bosses := []models.Boss{
		{Name: "BOSSJohn", Surname: "Doe"},
		{Name: "BOSSJane", Surname: "Smith"},
		{Name: "BOSSAlice", Surname: "Johnson"},
		{Name: "BOSSBob", Surname: "Brown"},
		{Name: "BOSSCharlie", Surname: "Davis"},
		{Name: "BOSSDavid", Surname: "Wilson"},
		{Name: "BOSSEve", Surname: "Taylor"},
		{Name: "BOSSFrank", Surname: "Anderson"},
		{Name: "BOSSGrace", Surname: "Thomas"},
		{Name: "BOSSHank", Surname: "Jackson"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&bosses).Error; err != nil {
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

	// Recupera gli ID generati automaticamente per i lavoratori e le specializzazioni
	var workersFromDB []models.Worker
	var specializationsFromDB []models.Specialization
	db.Find(&workersFromDB)
	db.Find(&specializationsFromDB)

	// Crea record di esempio per la tabella WorkerSpecialization
	workerSpecializations := []models.WorkerSpecialization{
		{WorkerID: workersFromDB[0].ID, SpecializationID: specializationsFromDB[0].ID},
		{WorkerID: workersFromDB[1].ID, SpecializationID: specializationsFromDB[1].ID},
		{WorkerID: workersFromDB[2].ID, SpecializationID: specializationsFromDB[2].ID},
		{WorkerID: workersFromDB[3].ID, SpecializationID: specializationsFromDB[3].ID},
		{WorkerID: workersFromDB[4].ID, SpecializationID: specializationsFromDB[4].ID},
		{WorkerID: workersFromDB[5].ID, SpecializationID: specializationsFromDB[5].ID},
		{WorkerID: workersFromDB[6].ID, SpecializationID: specializationsFromDB[6].ID},
		{WorkerID: workersFromDB[7].ID, SpecializationID: specializationsFromDB[7].ID},
		{WorkerID: workersFromDB[8].ID, SpecializationID: specializationsFromDB[8].ID},
		{WorkerID: workersFromDB[9].ID, SpecializationID: specializationsFromDB[9].ID},
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

	// Recupera gli ID generati automaticamente per le categorie di caschi
	var helmetCategoriesFromDB []models.HelmetCategory
	db.Find(&helmetCategoriesFromDB)

	// Crea record di esempio per la tabella Helmet
	helmets := []models.Helmet{
		{CategoryID: helmetCategoriesFromDB[0].ID},
		{CategoryID: helmetCategoriesFromDB[1].ID},
		{CategoryID: helmetCategoriesFromDB[0].ID},
		{CategoryID: helmetCategoriesFromDB[1].ID},
		{CategoryID: helmetCategoriesFromDB[0].ID},
		{CategoryID: helmetCategoriesFromDB[1].ID},
		{CategoryID: helmetCategoriesFromDB[0].ID},
		{CategoryID: helmetCategoriesFromDB[1].ID},
		{CategoryID: helmetCategoriesFromDB[0].ID},
		{CategoryID: helmetCategoriesFromDB[1].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&helmets).Error; err != nil {
		return err
	}

	// Recupera gli ID generati automaticamente per i caschi
	var helmetsFromDB []models.Helmet
	db.Find(&helmetsFromDB)

	// Crea record di esempio per la tabella Reading
	readings := []models.Reading{
		{HelmetID: helmetsFromDB[0].ID, Sensor1: 10.5, Sensor2: 20.3, Sensor3: 30.1, Anomalous: false},
		{HelmetID: helmetsFromDB[1].ID, Sensor1: 11.5, Sensor2: 21.3, Sensor3: 31.1, Anomalous: true},
		{HelmetID: helmetsFromDB[2].ID, Sensor1: 12.5, Sensor2: 22.3, Sensor3: 32.1, Anomalous: false},
		{HelmetID: helmetsFromDB[3].ID, Sensor1: 13.5, Sensor2: 23.3, Sensor3: 33.1, Anomalous: true},
		{HelmetID: helmetsFromDB[4].ID, Sensor1: 14.5, Sensor2: 24.3, Sensor3: 34.1, Anomalous: false},
		{HelmetID: helmetsFromDB[5].ID, Sensor1: 15.5, Sensor2: 25.3, Sensor3: 35.1, Anomalous: true},
		{HelmetID: helmetsFromDB[6].ID, Sensor1: 16.5, Sensor2: 26.3, Sensor3: 36.1, Anomalous: false},
		{HelmetID: helmetsFromDB[7].ID, Sensor1: 17.5, Sensor2: 27.3, Sensor3: 37.1, Anomalous: true},
		{HelmetID: helmetsFromDB[8].ID, Sensor1: 18.5, Sensor2: 28.3, Sensor3: 38.1, Anomalous: false},
		{HelmetID: helmetsFromDB[9].ID, Sensor1: 19.5, Sensor2: 29.3, Sensor3: 39.1, Anomalous: true},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&readings).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorkerAttendance
	workerAttendances := []models.WorkerAttendance{
		{WorkerID: workersFromDB[0].ID, WorksiteID: worksites[0].ID, HelmetID: helmetsFromDB[0].ID},
		{WorkerID: workersFromDB[1].ID, WorksiteID: worksites[1].ID, HelmetID: helmetsFromDB[1].ID},
		{WorkerID: workersFromDB[2].ID, WorksiteID: worksites[2].ID, HelmetID: helmetsFromDB[2].ID},
		{WorkerID: workersFromDB[3].ID, WorksiteID: worksites[3].ID, HelmetID: helmetsFromDB[3].ID},
		{WorkerID: workersFromDB[4].ID, WorksiteID: worksites[4].ID, HelmetID: helmetsFromDB[4].ID},
		{WorkerID: workersFromDB[5].ID, WorksiteID: worksites[5].ID, HelmetID: helmetsFromDB[5].ID},
		{WorkerID: workersFromDB[6].ID, WorksiteID: worksites[6].ID, HelmetID: helmetsFromDB[6].ID},
		{WorkerID: workersFromDB[7].ID, WorksiteID: worksites[7].ID, HelmetID: helmetsFromDB[7].ID},
		{WorkerID: workersFromDB[8].ID, WorksiteID: worksites[8].ID, HelmetID: helmetsFromDB[8].ID},
		{WorkerID: workersFromDB[9].ID, WorksiteID: worksites[9].ID, HelmetID: helmetsFromDB[9].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&workerAttendances).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorksiteBossAssignment
	worksiteBossAssignments := []models.WorksiteBossAssignment{
		{BossID: workersFromDB[0].ID, WorksiteID: worksites[0].ID},
		{BossID: workersFromDB[1].ID, WorksiteID: worksites[1].ID},
		{BossID: workersFromDB[2].ID, WorksiteID: worksites[2].ID},
		{BossID: workersFromDB[3].ID, WorksiteID: worksites[3].ID},
		{BossID: workersFromDB[4].ID, WorksiteID: worksites[4].ID},
		{BossID: workersFromDB[5].ID, WorksiteID: worksites[5].ID},
		{BossID: workersFromDB[6].ID, WorksiteID: worksites[6].ID},
		{BossID: workersFromDB[7].ID, WorksiteID: worksites[7].ID},
		{BossID: workersFromDB[8].ID, WorksiteID: worksites[8].ID},
		{BossID: workersFromDB[9].ID, WorksiteID: worksites[9].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksiteBossAssignments).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorksiteWorkerAssignment
	worksiteWorkerAssignments := []models.WorksiteWorkerAssignment{
		{WorksiteID: worksites[0].ID, WorkerID: workersFromDB[0].ID, AssignedBy: bosses[0].ID},
		{WorksiteID: worksites[1].ID, WorkerID: workersFromDB[1].ID, AssignedBy: bosses[1].ID},
		{WorksiteID: worksites[2].ID, WorkerID: workersFromDB[2].ID, AssignedBy: bosses[2].ID},
		{WorksiteID: worksites[3].ID, WorkerID: workersFromDB[3].ID, AssignedBy: bosses[3].ID},
		{WorksiteID: worksites[4].ID, WorkerID: workersFromDB[4].ID, AssignedBy: bosses[4].ID},
		{WorksiteID: worksites[5].ID, WorkerID: workersFromDB[5].ID, AssignedBy: bosses[5].ID},
		{WorksiteID: worksites[6].ID, WorkerID: workersFromDB[6].ID, AssignedBy: bosses[6].ID},
		{WorksiteID: worksites[7].ID, WorkerID: workersFromDB[7].ID, AssignedBy: bosses[7].ID},
		{WorksiteID: worksites[8].ID, WorkerID: workersFromDB[8].ID, AssignedBy: bosses[8].ID},
		{WorksiteID: worksites[9].ID, WorkerID: workersFromDB[9].ID, AssignedBy: bosses[9].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksiteWorkerAssignments).Error; err != nil {
		return err
	}

	log.Println("Database seeded successfully")
	return nil
}
