package models

import "time"

type Worksite struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	StartDateOfWork time.Time `json:"start_date_of_work"`
	EndDateOfWork   time.Time `json:"end_date_of_work"`
}

type Worker struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Surname   string    `json:"surname" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Boss struct {
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
	ID                    int       `json:"id" gorm:"primaryKey"`
	ReadAt                time.Time `json:"read_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	HelmetID              int       `json:"helmet_id" gorm:"not null"`
	Temperature           float64   `json:"temperature"`
	Humidity              float64   `json:"humidity"`
	Brightness            float64   `json:"brightness"`
	Methane               bool      `json:"methane"`
	CarbonMonoxide        bool      `json:"carbon_monoxide"`
	SmokeDetection        bool      `json:"smoke_detection"`
	UsesWeldingProtection bool      `json:"uses_welding_protection"`
	UsesGasProtection     bool      `json:"uses_gas_protection"`
	Avg_X                 float64   `json:"avg_X"`
	Avg_Y                 float64   `json:"avg_Y"`
	Avg_Z                 float64   `json:"avg_Z"`
	Avg_G                 float64   `json:"avg_G"`
	Std_X                 float64   `json:"std_X"`
	Std_Y                 float64   `json:"std_Y"`
	Std_Z                 float64   `json:"std_Z"`
	Std_G                 float64   `json:"std_G"`
	Max_G                 float64   `json:"max_G"`
	IncorrectPosture      float64   `json:"incorrect_posture"`
	Anomaly               bool      `json:"anomaly" gorm:"not null"`
	Helmet                Helmet    `gorm:"foreignKey:HelmetID;references:ID"`
}

type WorkerAttendance struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	WorkerID   int        `json:"worker_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	StartAt    time.Time  `json:"start_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	EndAt      *time.Time `json:"end_at"`
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
	Boss       Boss       `gorm:"foreignKey:BossID;references:ID"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID"`
}

type WorksiteWorkerAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	WorkerID   int        `json:"worker_id" gorm:"not null"`
	AssignedBy int        `json:"assigned_by" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID"`
	Worker     Worker     `gorm:"foreignKey:WorkerID;references:ID"`
	Boss       Boss       `gorm:"foreignKey:AssignedBy;references:ID"`
}

type WeatherData struct {
	WorksiteID int       `json:"worksite_id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at" gorm:"primaryKey;default:CURRENT_TIMESTAMP"`
	Temp       float64   `json:"temp"`
	TempMin    float64   `json:"temp_min"`
	TempMax    float64   `json:"temp_max"`
	Humidity   int       `json:"humidity"`
	Brightness int       `json:"brightness"`
	C0         float64   `json:"c0"`
	PM10       float64   `json:"pm10"`
	Worksite   Worksite  `gorm:"foreignKey:WorksiteID;references:ID"`
}
