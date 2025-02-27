package utils

import (
	"log"
	"safecap_backend/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDatabase(db *gorm.DB) error {

	// Crea record di esempio per la tabella Worksite
	worksites := []models.Worksite{
		{Name: "Cantiere buca autostradale", Latitude: 44.609554, Longitude: 10.905925, City: "Modena", ZipCode: "41125", State: "Italy", Address: "Autostrada A1, km 163", StartAt: time.Now()},
		{Name: "Cantiere edificio MO-27 DIEF Unimore", Latitude: 44.629542, Longitude: 10.948303, City: "Modena", ZipCode: "41125", State: "Italy", Address: "Via Vignolese, 905", StartAt: time.Now()},
		{Name: "Cantiere ristrutturazione Duomo di Milano", Latitude: 45.464222, Longitude: 9.1915, City: "Milano", ZipCode: "20122", State: "Italy", Address: "P.za del Duomo", StartAt: time.Now()},
		{Name: "Cantiere Ospedale Casa Sollievo della Sofferenza", Latitude: 41.708638, Longitude: 15.706511, City: "San Giovanni Rotondo", ZipCode: "71013", State: "Italy", Address: "Viale Cappuccini, 1", StartAt: time.Now()},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksites).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Worker
	workers := []models.Worker{
		{Name: "Mario", Surname: "Rossi", Email: "mario.rossi@example.com", Password: "password", Phone: "1234567890", Active: true, FiscalCode: "MRIROS123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Giulia", Surname: "Verdi", Email: "giulia.verdi@example.com", Password: "password", Phone: "1234567891", Active: true, FiscalCode: "GIUVRD123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Gianni", Surname: "Bianchi", Email: "gianni.bianchi@example.com", Password: "password", Phone: "1234567892", Active: true, FiscalCode: "GINBNC123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Chiara", Surname: "Celesti", Email: "chiara.celesti@example.com", Password: "password", Phone: "1234567893", Active: true, FiscalCode: "CHRCLS123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&workers).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Boss
	bosses := []models.Boss{
		{Name: "John", Surname: "Doe", Email: "john.doe@example.com", Password: "password", Phone: "1234567800", Active: true, FiscalCode: "BOSJHN123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Jane", Surname: "Smith", Email: "jane.smith@example.com", Password: "password", Phone: "1234567801", Active: true, FiscalCode: "BOSJNS123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Alice", Surname: "Johnson", Email: "alice.johnson@example.com", Password: "password", Phone: "1234567802", Active: true, FiscalCode: "BOSALJ123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&bosses).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella Specialization
	specializations := []models.Specialization{
		{Name: "Electrician"},
		{Name: "Plumber"},
		{Name: "Carpenter"},
		{Name: "Mechanic"},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&specializations).Error; err != nil {
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
		{CategoryID: helmetCategoriesFromDB[0].ID, MACAddress: "AC:67:B2:05:02:8E", CreatedAt: time.Now()},
		{CategoryID: helmetCategoriesFromDB[1].ID, MACAddress: "AC:15:18:E5:8E:82", CreatedAt: time.Now()},
		{CategoryID: helmetCategoriesFromDB[0].ID, MACAddress: "F0:24:F9:59:02:E2", CreatedAt: time.Now()},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&helmets).Error; err != nil {
		return err
	}

	/* // Recupera gli ID generati automaticamente per i lavoratori e le specializzazioni
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

	// Recupera gli ID generati automaticamente per i caschi
	var helmetsFromDB []models.Helmet
	db.Find(&helmetsFromDB)

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

	// Recupera gli ID generati automaticamente per le attendance
	var workerAttendancesFromDB []models.WorkerAttendance
	db.Find(&workerAttendancesFromDB)

	// Crea record di esempio per la tabella Reading
	readings := []models.Reading{
		{AttendanceID: workerAttendancesFromDB[0].ID, ReadAt: time.Now(), Temperature: 10.5, Humidity: 45.0, Brightness: 20, Methane: false, CarbonMonoxide: true, SmokeDetection: false, UsesWeldingProtection: false, UsesGasProtection: true, Avg_X: 1.0, Avg_Y: 2.0, Avg_Z: 3.0, Avg_G: 4.0, Std_X: 0.1, Std_Y: 0.2, Std_Z: 0.3, Std_G: 0.4, Max_G: 4.5, IncorrectPosture: 0.0, Anomaly: false},
		{AttendanceID: workerAttendancesFromDB[1].ID, ReadAt: time.Now(), Temperature: 11.5, Humidity: 50.0, Brightness: 21, Methane: true, CarbonMonoxide: false, SmokeDetection: true, UsesWeldingProtection: true, UsesGasProtection: false, Avg_X: 1.1, Avg_Y: 2.1, Avg_Z: 3.1, Avg_G: 4.1, Std_X: 0.2, Std_Y: 0.3, Std_Z: 0.4, Std_G: 0.5, Max_G: 4.6, IncorrectPosture: 0.1, Anomaly: true},
		{AttendanceID: workerAttendancesFromDB[2].ID, ReadAt: time.Now(), Temperature: 12.5, Humidity: 55.0, Brightness: 22, Methane: false, CarbonMonoxide: true, SmokeDetection: false, UsesWeldingProtection: false, UsesGasProtection: true, Avg_X: 1.2, Avg_Y: 2.2, Avg_Z: 3.2, Avg_G: 4.2, Std_X: 0.3, Std_Y: 0.4, Std_Z: 0.5, Std_G: 0.6, Max_G: 4.7, IncorrectPosture: 0.2, Anomaly: false},
		{AttendanceID: workerAttendancesFromDB[3].ID, ReadAt: time.Now(), Temperature: 13.5, Humidity: 60.0, Brightness: 23, Methane: true, CarbonMonoxide: false, SmokeDetection: true, UsesWeldingProtection: true, UsesGasProtection: false, Avg_X: 1.3, Avg_Y: 2.3, Avg_Z: 3.3, Avg_G: 4.3, Std_X: 0.4, Std_Y: 0.5, Std_Z: 0.6, Std_G: 0.7, Max_G: 4.8, IncorrectPosture: 0.3, Anomaly: true},
		{AttendanceID: workerAttendancesFromDB[4].ID, ReadAt: time.Now(), Temperature: 14.5, Humidity: 65.0, Brightness: 24, Methane: false, CarbonMonoxide: true, SmokeDetection: false, UsesWeldingProtection: false, UsesGasProtection: true, Avg_X: 1.4, Avg_Y: 2.4, Avg_Z: 3.4, Avg_G: 4.4, Std_X: 0.5, Std_Y: 0.6, Std_Z: 0.7, Std_G: 0.8, Max_G: 4.9, IncorrectPosture: 0.4, Anomaly: false},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&readings).Error; err != nil {
		return err
	}

	// Recupera gli ID generati automaticamente per i boss
	var bossesFromDB []models.Boss
	db.Find(&bossesFromDB)

	var worksitesFromDB []models.Worksite
	db.Find(&worksitesFromDB)

	// Crea record di esempio per la tabella WorksiteBossAssignment
	worksiteBossAssignments := []models.WorksiteBossAssignment{
		{BossID: bossesFromDB[0].ID, WorksiteID: worksitesFromDB[0].ID},
		{BossID: bossesFromDB[1].ID, WorksiteID: worksitesFromDB[1].ID},
		{BossID: bossesFromDB[2].ID, WorksiteID: worksitesFromDB[2].ID},
		{BossID: bossesFromDB[3].ID, WorksiteID: worksitesFromDB[3].ID},
		{BossID: bossesFromDB[4].ID, WorksiteID: worksitesFromDB[4].ID},
		{BossID: bossesFromDB[5].ID, WorksiteID: worksitesFromDB[5].ID},
		{BossID: bossesFromDB[6].ID, WorksiteID: worksitesFromDB[6].ID},
		{BossID: bossesFromDB[7].ID, WorksiteID: worksitesFromDB[7].ID},
		{BossID: bossesFromDB[8].ID, WorksiteID: worksitesFromDB[8].ID},
		{BossID: bossesFromDB[9].ID, WorksiteID: worksitesFromDB[9].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksiteBossAssignments).Error; err != nil {
		return err
	}

	// Crea record di esempio per la tabella WorksiteWorkerAssignment
	worksiteWorkerAssignments := []models.WorksiteWorkerAssignment{
		{WorksiteID: worksitesFromDB[0].ID, WorkerID: workersFromDB[0].ID, AssignedBy: bossesFromDB[0].ID},
		{WorksiteID: worksitesFromDB[1].ID, WorkerID: workersFromDB[1].ID, AssignedBy: bossesFromDB[1].ID},
		{WorksiteID: worksitesFromDB[2].ID, WorkerID: workersFromDB[2].ID, AssignedBy: bossesFromDB[2].ID},
		{WorksiteID: worksitesFromDB[3].ID, WorkerID: workersFromDB[3].ID, AssignedBy: bossesFromDB[3].ID},
		{WorksiteID: worksitesFromDB[4].ID, WorkerID: workersFromDB[4].ID, AssignedBy: bossesFromDB[4].ID},
		{WorksiteID: worksitesFromDB[5].ID, WorkerID: workersFromDB[5].ID, AssignedBy: bossesFromDB[5].ID},
		{WorksiteID: worksitesFromDB[6].ID, WorkerID: workersFromDB[6].ID, AssignedBy: bossesFromDB[6].ID},
		{WorksiteID: worksitesFromDB[7].ID, WorkerID: workersFromDB[7].ID, AssignedBy: bossesFromDB[7].ID},
		{WorksiteID: worksitesFromDB[8].ID, WorkerID: workersFromDB[8].ID, AssignedBy: bossesFromDB[8].ID},
		{WorksiteID: worksitesFromDB[9].ID, WorkerID: workersFromDB[9].ID, AssignedBy: bossesFromDB[9].ID},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&worksiteWorkerAssignments).Error; err != nil {
		return err
	} */

	log.Println("Database seeded successfully")
	return nil
}
