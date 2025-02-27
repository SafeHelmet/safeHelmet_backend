package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Worksite struct {
	ID                   int       `json:"id" gorm:"primaryKey"`
	Name                 string    `json:"name" gorm:"not null;size:100" validate:"max=100"`
	CreatedAt            time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Latitude             float64   `json:"latitude" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Longitude            float64   `json:"longitude" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Address              string    `json:"address" gorm:"size:255" validate:"max=255"`
	City                 string    `json:"city" gorm:"not null;size:100" validate:"max=100"`
	ZipCode              string    `json:"zip_code" gorm:"not null;size:20" validate:"max=20"`
	State                string    `json:"state" gorm:"not null;size:100" validate:"max=100"`
	StartAt              time.Time `json:"start_date_of_work" gorm:"not null"`
	EndAt                time.Time `json:"end_date_of_work"`
	TemperatureThreshold float64   `json:"temperature_threshold" gorm:"default:10;type:decimal(5,2)" validate:"decimal=5,2"`
	HumidityThreshold    float64   `json:"humidity_threshold" gorm:"default:10;type:decimal(5,2)" validate:"decimal=5,2"`
	BrightnessThreshold  float64   `json:"brightness_threshold" gorm:"default:700;type:decimal(10,2)" validate:"decimal=10,2"`
	PostureThreshold     float64   `json:"posture_threshold" gorm:"default:0.5;type:decimal(5,2)" validate:"decimal=5,2"`
	MaxGThreshold        float64   `json:"max_g_threshold" gorm:"default:4;type:decimal(5,2)" validate:"decimal=5,2"`
}

type Worker struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null;size:100" validate:"max=100"`
	Surname    string    `json:"surname" gorm:"not null;size:100" validate:"max=100"`
	Email      string    `json:"email" gorm:"not null;size:100" validate:"max=100"`
	Password   string    `json:"password" gorm:"not null;default:'password';size:100" validate:"max=100"`
	Phone      string    `json:"phone" gorm:"not null;size:20" validate:"max=20"`
	Active     bool      `json:"active" gorm:"default:true"`
	FiscalCode string    `json:"fiscal_code" gorm:"not null;size:20" validate:"max=20"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Boss struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null;size:100" validate:"max=100"`
	Surname    string    `json:"surname" gorm:"not null;size:100" validate:"max=100"`
	Email      string    `json:"email" gorm:"not null;size:100" validate:"max=100"`
	Password   string    `json:"password" gorm:"not null;default:'password';size:100" validate:"max=100"`
	Phone      string    `json:"phone" gorm:"not null;size:20" validate:"max=20"`
	FiscalCode string    `json:"fiscal_code" gorm:"not null;size:20" validate:"max=20"`
	Active     bool      `json:"active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Specialization struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;size:100" validate:"max=100"`
}

type WorkerSpecialization struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	WorkerID         int            `json:"worker_id" gorm:"not null"`
	SpecializationID int            `json:"specialization_id" gorm:"not null"`
	Worker           Worker         `gorm:"foreignKey:WorkerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Helmet struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	CategoryID int            `json:"category_id" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Category   HelmetCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MACAddress string         `json:"mac_address" gorm:"not null;size:20" validate:"max=20"`
}

type HelmetCategory struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;size:100" validate:"max=100"`
}

type WorkerAttendance struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	WorkerID   int        `json:"worker_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	StartAt    time.Time  `json:"start_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	EndAt      *time.Time `json:"end_at"`
	HelmetID   int        `json:"helmet_id"`
	Worker     Worker     `gorm:"foreignKey:WorkerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Helmet     Helmet     `gorm:"foreignKey:HelmetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Reading struct {
	ID                    int              `json:"id" gorm:"primaryKey"`
	ReadAt                time.Time        `json:"read_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	AttendanceID          int              `json:"attendance_id" gorm:"not null"`
	Temperature           float64          `json:"temperature" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	WeatherTemperature    float64          `json:"weather_temperature" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	WeatherTemperatureMax float64          `json:"weather_temperature_max" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	WeatherTemperatureMin float64          `json:"weather_temperature_min" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	AnomalousTemperature  bool             `json:"anomalous_temperature" gorm:"not null"`
	Humidity              float64          `json:"humidity" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	WeatherHumidity       float64          `json:"weather_humidity" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	AnomalousHumidity     bool             `json:"anomalous_humidity" gorm:"not null"`
	Brightness            float64          `json:"brightness" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	WeatherBrightness     float64          `json:"weather_brightness" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	AnomalousBrightness   bool             `json:"anomalous_brightness" gorm:"not null"`
	Methane               bool             `json:"methane" gorm:"not null"`
	CarbonMonoxide        bool             `json:"carbon_monoxide" gorm:"not null"`
	SmokeDetection        bool             `json:"smoke_detection" gorm:"not null"`
	UsesWeldingProtection bool             `json:"uses_welding_protection" gorm:"not null"`
	UsesGasProtection     bool             `json:"uses_gas_protection" gorm:"not null"`
	Avg_X                 float64          `json:"avg_X" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Avg_Y                 float64          `json:"avg_Y" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Avg_Z                 float64          `json:"avg_Z" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Avg_G                 float64          `json:"avg_G" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Std_X                 float64          `json:"std_X" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Std_Y                 float64          `json:"std_Y" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Std_Z                 float64          `json:"std_Z" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Std_G                 float64          `json:"std_G" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	Max_G                 float64          `json:"max_G" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	AnomalousMaxG         bool             `json:"anomalous_max_g" gorm:"not null"`
	IncorrectPosture      float64          `json:"incorrect_posture" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	AnomalousPosture      bool             `json:"anomalous_posture" gorm:"not null"`
	Anomaly               bool             `json:"anomaly" gorm:"not null"`
	Attendance            WorkerAttendance `gorm:"foreignKey:AttendanceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type WorksiteBossAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	BossID     int        `json:"boss_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	Boss       Boss       `gorm:"foreignKey:BossID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type WorksiteWorkerAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	WorkerID   int        `json:"worker_id" gorm:"not null"`
	AssignedBy int        `json:"assigned_by" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Worker     Worker     `gorm:"foreignKey:WorkerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Boss       Boss       `gorm:"foreignKey:AssignedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type WeatherData struct {
	WorksiteID int       `json:"worksite_id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at" gorm:"primaryKey;default:CURRENT_TIMESTAMP"`
	Temp       float64   `json:"temp" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	TempMin    float64   `json:"temp_min" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	TempMax    float64   `json:"temp_max" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	Humidity   float64   `json:"humidity" gorm:"not null;type:decimal(5,2)" validate:"decimal=5,2"`
	Brightness float64   `json:"brightness" gorm:"not null;type:decimal(10,2)" validate:"decimal=10,2"`
	C0         float64   `json:"c0" gorm:"type:decimal(10,2)" validate:"decimal=10,2"`
	PM10       float64   `json:"pm10" gorm:"type:decimal(10,2)" validate:"decimal=10,2"`
	Worksite   Worksite  `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Worksites delete hooks
func (w *Worksite) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorksiteBossAssignment{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorksiteWorkerAssignment{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorkerAttendance{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WeatherData{})
	return
}

func (w *WorkerAttendance) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("attendance_id = ?", w.ID).Delete(&Reading{})
	return
}

// Worker delete hooks
func (w *Worker) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorkerSpecialization{})
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorkerAttendance{})
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorksiteWorkerAssignment{})
	return
}

// Boss delete hooks
func (b *Boss) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("boss_id = ?", b.ID).Delete(&WorksiteBossAssignment{})
	tx.Unscoped().Where("assigned_by = ?", b.ID).Delete(&WorksiteWorkerAssignment{})
	return
}

// Helmet delete hooks
func (h *Helmet) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("helmet_id = ?", h.ID).Delete(&WorkerAttendance{})
	return
}

// Specialization delete hooks
func (s *Specialization) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("specialization_id = ?", s.ID).Delete(&WorkerSpecialization{})
	return
}

// HelmetCategory delete hooks
func (hc *HelmetCategory) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("category_id = ?", hc.ID).Delete(&Helmet{})
	return
}

func (w *Worksite) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(w)
}

func (w *WorkerAttendance) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(w)
}

// Worker delete hooks
func (w *Worker) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(w)
}

// Boss delete hooks
func (b *Boss) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(b)
}

// Helmet delete hooks
func (h *Helmet) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(h)
}

// Specialization delete hooks
func (s *Specialization) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(s)
}

// HelmetCategory delete hooks
func (hc *HelmetCategory) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(hc)
}

func (r *Reading) BeforeCreate(tx *gorm.DB) (err error) {
	return validate.Struct(r)
}
