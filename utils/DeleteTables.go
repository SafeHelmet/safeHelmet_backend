package utils

import "gorm.io/gorm"

func DeleteTables(db *gorm.DB) {
	db.Exec("DELETE FROM worksite_boss_assignments")
	db.Exec("DELETE FROM worker_worksite_assignments")
	db.Exec("DELETE FROM worker_attendances")
	db.Exec("DELETE FROM readings")
	db.Exec("DELETE FROM helmets")
	db.Exec("DELETE FROM helmet_categories")
	db.Exec("DELETE FROM worker_specializations")
	db.Exec("DELETE FROM specializations")
	db.Exec("DELETE FROM workers")
	db.Exec("DELETE FROM worksites")
	db.Exec("DELETE FROM bosses")
}
