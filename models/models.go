package models

import "time"

type Worksite struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Worker struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null"`
	Surname string `json:"surname" gorm:"not null"`
}

type Specialization struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type WorkerSpecialization struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	WorkerID         int            `json:"worker_id" gorm:"not null"`
	SpecializationID int            `json:"specialization_id" gorm:"not null"`
	Worker           Worker         `gorm:"foreignKey:WorkerID;references:ID"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID;references:ID"`
}

type Helmet struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	CategoryID int            `json:"category_id" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Category   HelmetCategory `gorm:"foreignKey:CategoryID;references:ID"`
}

type HelmetCategory struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type Reading struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	ReadAt    time.Time `json:"read_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	HelmetID  int       `json:"helmet_id" gorm:"not null"`
	Sensor1   float64   `json:"sensor1"`
	Sensor2   float64   `json:"sensor2"`
	Sensor3   float64   `json:"sensor3"`
	Anomalous bool      `json:"anomalous"`
	Helmet    Helmet    `gorm:"foreignKey:HelmetID;references:ID"`
}

type WorkerWorksiteAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	WorkerID   int        `json:"worker_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	HelmetID   int        `json:"helmet_id"`
	Worker     Worker     `gorm:"foreignKey:WorkerID;references:ID"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID"`
	Helmet     Helmet     `gorm:"foreignKey:HelmetID;references:ID"`
}

type WorksiteBossAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	BossID     int        `json:"boss_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	Boss       Worker     `gorm:"foreignKey:BossID;references:ID"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID"`
}
