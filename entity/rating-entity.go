package entity

import "gorm.io/gorm"

// GORM entity for Rating
type Rating struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	Certification   string `gorm:"column:certification"`
	Task_Completion string `gorm:"column:task_completion"`
	Help            string `gorm:"column:help"`
	EmployeeID      uint   `gorm:"column:employeeID"`
}
