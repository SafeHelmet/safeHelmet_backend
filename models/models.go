package models

import (
	"math"
	"time"

	"gorm.io/gorm"
)

func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Truncate(string string, length int) string {
	if len(string) > length {
		return string[:length]
	}
	return string
}

// Worksite model
type Worksite struct {
	ID                   int        `json:"id" gorm:"primaryKey"`
	Name                 string     `json:"name" gorm:"not null;size:100"`
	CreatedAt            time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Latitude             float64    `json:"latitude" gorm:"not null;"`
	Longitude            float64    `json:"longitude" gorm:"not null;"`
	Address              string     `json:"address" gorm:"size:255"`
	City                 string     `json:"city" gorm:"not null;size:100"`
	ZipCode              string     `json:"zip_code" gorm:"not null;size:20"`
	State                string     `json:"state" gorm:"not null;size:100"`
	StartAt              time.Time  `json:"start_date_of_work" gorm:"not null"`
	EndAt                *time.Time `json:"end_date_of_work"`
	TemperatureThreshold float64    `json:"temperature_threshold" gorm:"default:15;"`
	HumidityThreshold    float64    `json:"humidity_threshold" gorm:"default:25;"`
	BrightnessThreshold  float64    `json:"brightness_threshold" gorm:"default:700;"`
	PostureThreshold     float64    `json:"posture_threshold" gorm:"default:0.5;"`
	MaxGThreshold        float64    `json:"max_g_threshold" gorm:"default:4;"`
}

func (w *Worksite) BeforeCreate(tx *gorm.DB) (err error) {
	w.Name = Truncate(w.Name, 100)
	w.Address = Truncate(w.Address, 255)
	w.City = Truncate(w.City, 100)
	w.State = Truncate(w.State, 100)
	w.TemperatureThreshold = RoundFloat(w.TemperatureThreshold, 2)
	w.HumidityThreshold = RoundFloat(w.HumidityThreshold, 2)
	w.BrightnessThreshold = RoundFloat(w.BrightnessThreshold, 2)
	w.PostureThreshold = RoundFloat(w.PostureThreshold, 2)
	w.MaxGThreshold = RoundFloat(w.MaxGThreshold, 2)
	return
}

// Worksites delete hooks
func (w *Worksite) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorksiteBossAssignment{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorksiteWorkerAssignment{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WorkerAttendance{})
	tx.Unscoped().Where("worksite_id = ?", w.ID).Delete(&WeatherData{})
	return
}

// Worker model
type Worker struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null;size:100"`
	Surname    string    `json:"surname" gorm:"not null;size:100"`
	Email      string    `json:"email" gorm:"not null;size:100"`
	Password   string    `json:"password" gorm:"not null;default:'password';size:100"`
	Phone      string    `json:"phone" gorm:"not null;size:20"`
	Active     bool      `json:"active" gorm:"default:true"`
	FiscalCode string    `json:"fiscal_code" gorm:"not null;size:20"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (w *Worker) BeforeCreate(tx *gorm.DB) (err error) {
	w.Name = Truncate(w.Name, 100)
	w.Surname = Truncate(w.Surname, 100)
	w.Phone = Truncate(w.Phone, 20)
	w.FiscalCode = Truncate(w.FiscalCode, 20)
	return
}

// Worker delete hooks
func (w *Worker) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorkerSpecialization{})
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorkerAttendance{})
	tx.Unscoped().Where("worker_id = ?", w.ID).Delete(&WorksiteWorkerAssignment{})
	return
}

// Boss model
type Boss struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null;size:100"`
	Surname    string    `json:"surname" gorm:"not null;size:100"`
	Email      string    `json:"email" gorm:"not null;size:100"`
	Password   string    `json:"password" gorm:"not null;default:'password';size:100"`
	Phone      string    `json:"phone" gorm:"not null;size:20"`
	FiscalCode string    `json:"fiscal_code" gorm:"not null;size:20"`
	Active     bool      `json:"active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (b *Boss) BeforeCreate(tx *gorm.DB) (err error) {
	b.Name = Truncate(b.Name, 100)
	b.Surname = Truncate(b.Surname, 100)
	b.Phone = Truncate(b.Phone, 20)
	b.FiscalCode = Truncate(b.FiscalCode, 20)
	return
}

// Boss delete hooks
func (b *Boss) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("boss_id = ?", b.ID).Delete(&WorksiteBossAssignment{})
	tx.Unscoped().Where("assigned_by = ?", b.ID).Delete(&WorksiteWorkerAssignment{})
	return
}

// Specialization model
type Specialization struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;size:100"`
}

func (s *Specialization) BeforeCreate(tx *gorm.DB) (err error) {
	s.Name = Truncate(s.Name, 100)
	return
}

// Specialization delete hooks
func (s *Specialization) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("specialization_id = ?", s.ID).Delete(&WorkerSpecialization{})
	return
}

// WorkerSpecialization model
type WorkerSpecialization struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	WorkerID         int            `json:"worker_id" gorm:"not null"`
	SpecializationID int            `json:"specialization_id" gorm:"not null"`
	Worker           Worker         `gorm:"foreignKey:WorkerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Helmet model
type Helmet struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	CategoryID int            `json:"category_id" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Category   HelmetCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MACAddress string         `json:"mac_address" gorm:"not null;size:17"`
}

// Helmet delete hooks
func (h *Helmet) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("helmet_id = ?", h.ID).Delete(&WorkerAttendance{})
	return
}

// HelmetCategory model
type HelmetCategory struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;size:100"`
}

func (h *HelmetCategory) BeforeCreate(tx *gorm.DB) (err error) {
	h.Name = Truncate(h.Name, 100)
	return
}

// HelmetCategory delete hooks
func (hc *HelmetCategory) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("category_id = ?", hc.ID).Delete(&Helmet{})
	return
}

// WorkerAttendance model
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

// WorkerAttendance delete hooks
func (w *WorkerAttendance) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Unscoped().Where("attendance_id = ?", w.ID).Delete(&Reading{})
	return
}

// Reading model
type Reading struct {
	ID                    int              `json:"id" gorm:"primaryKey"`
	ReadAt                time.Time        `json:"read_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	AttendanceID          int              `json:"attendance_id" gorm:"not null"`
	Temperature           float64          `json:"temperature" gorm:"not null;"`
	WeatherTemperature    float64          `json:"weather_temperature" gorm:"not null;"`
	WeatherTemperatureMax float64          `json:"weather_temperature_max" gorm:"not null;"`
	WeatherTemperatureMin float64          `json:"weather_temperature_min" gorm:"not null;"`
	AnomalousTemperature  bool             `json:"anomalous_temperature" gorm:"not null"`
	Humidity              float64          `json:"humidity" gorm:"not null;"`
	WeatherHumidity       float64          `json:"weather_humidity" gorm:"not null;"`
	AnomalousHumidity     bool             `json:"anomalous_humidity" gorm:"not null"`
	Brightness            float64          `json:"brightness" gorm:"not null;"`
	WeatherBrightness     float64          `json:"weather_brightness" gorm:"not null;"`
	AnomalousBrightness   bool             `json:"anomalous_brightness" gorm:"not null"`
	Methane               bool             `json:"methane" gorm:"not null"`
	CarbonMonoxide        bool             `json:"carbon_monoxide" gorm:"not null"`
	SmokeDetection        bool             `json:"smoke_detection" gorm:"not null"`
	UsesWeldingProtection bool             `json:"uses_welding_protection" gorm:"not null"`
	UsesGasProtection     bool             `json:"uses_gas_protection" gorm:"not null"`
	Avg_X                 float64          `json:"avg_X" gorm:"not null;"`
	Avg_Y                 float64          `json:"avg_Y" gorm:"not null;"`
	Avg_Z                 float64          `json:"avg_Z" gorm:"not null;"`
	Avg_G                 float64          `json:"avg_G" gorm:"not null;"`
	Std_X                 float64          `json:"std_X" gorm:"not null;"`
	Std_Y                 float64          `json:"std_Y" gorm:"not null;"`
	Std_Z                 float64          `json:"std_Z" gorm:"not null;"`
	Std_G                 float64          `json:"std_G" gorm:"not null;"`
	Max_G                 float64          `json:"max_G" gorm:"not null;"`
	AnomalousMaxG         bool             `json:"anomalous_max_g" gorm:"not null"`
	IncorrectPosture      float64          `json:"incorrect_posture" gorm:"not null;"`
	AnomalousPosture      bool             `json:"anomalous_posture" gorm:"not null"`
	Anomaly               bool             `json:"anomaly" gorm:"not null"`
	Attendance            WorkerAttendance `gorm:"foreignKey:AttendanceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (r *Reading) BeforeCreate(tx *gorm.DB) (err error) {
	r.Temperature = RoundFloat(r.Temperature, 2)
	r.WeatherTemperature = RoundFloat(r.WeatherTemperature, 2)
	r.WeatherTemperatureMax = RoundFloat(r.WeatherTemperatureMax, 2)
	r.WeatherTemperatureMin = RoundFloat(r.WeatherTemperatureMin, 2)
	r.Humidity = RoundFloat(r.Humidity, 2)
	r.WeatherHumidity = RoundFloat(r.WeatherHumidity, 2)
	r.Brightness = RoundFloat(r.Brightness, 2)
	r.WeatherBrightness = RoundFloat(r.WeatherBrightness, 2)
	r.Avg_X = RoundFloat(r.Avg_X, 2)
	r.Avg_Y = RoundFloat(r.Avg_Y, 2)
	r.Avg_Z = RoundFloat(r.Avg_Z, 2)
	r.Avg_G = RoundFloat(r.Avg_G, 2)
	r.Std_X = RoundFloat(r.Std_X, 2)
	r.Std_Y = RoundFloat(r.Std_Y, 2)
	r.Std_Z = RoundFloat(r.Std_Z, 2)
	r.Std_G = RoundFloat(r.Std_G, 2)
	r.Max_G = RoundFloat(r.Max_G, 2)
	r.IncorrectPosture = RoundFloat(r.IncorrectPosture, 2)
	return
}

// WorksiteBossAssignment model
type WorksiteBossAssignment struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	BossID     int        `json:"boss_id" gorm:"not null"`
	WorksiteID int        `json:"worksite_id" gorm:"not null"`
	AssignedAt time.Time  `json:"assigned_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReleasedAt *time.Time `json:"released_at"`
	Boss       Boss       `gorm:"foreignKey:BossID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Worksite   Worksite   `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// WorksiteWorkerAssignment model
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

// WeatherData model
type WeatherData struct {
	WorksiteID int       `json:"worksite_id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at" gorm:"primaryKey;default:CURRENT_TIMESTAMP"`
	Temp       float64   `json:"temp" gorm:"not null;"`
	TempMin    float64   `json:"temp_min" gorm:"not null;"`
	TempMax    float64   `json:"temp_max" gorm:"not null;"`
	Humidity   float64   `json:"humidity" gorm:"not null;"`
	Brightness float64   `json:"brightness" gorm:"not null;"`
	C0         float64   `json:"c0" gorm:"default:0;"`
	PM10       float64   `json:"pm10" gorm:"default:0;"`
	Worksite   Worksite  `gorm:"foreignKey:WorksiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (w *WeatherData) BeforeCreate(tx *gorm.DB) (err error) {
	w.Temp = RoundFloat(w.Temp, 2)
	w.TempMin = RoundFloat(w.TempMin, 2)
	w.TempMax = RoundFloat(w.TempMax, 2)
	w.Humidity = RoundFloat(w.Humidity, 2)
	w.Brightness = RoundFloat(w.Brightness, 2)
	w.C0 = RoundFloat(w.C0, 2)
	w.PM10 = RoundFloat(w.PM10, 2)
	return
}
