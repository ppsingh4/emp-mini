package entity

import "gorm.io/gorm"

// GORM entity for Employee
type Employee struct {
	gorm.Model
    ID           uint  `gorm:"primaryKey"`
    First_Name   string `gorm:"column:first_name"`
    Second_Name  string `gorm:"column:second_name"`
    Email        string `gorm:"column:email"`
    Phone_number string `gorm:"column:phone_number"`
    Department   string `gorm:"column:department"`
}


