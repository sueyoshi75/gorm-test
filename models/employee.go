package models

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Employeeid string `json:"employeeid"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Model[Employee] `gorm:"-" json:"-"`
}

func (Employee) GetEmployeesByEmployeeid(tx *gorm.DB, employeeid string) ([]Employee, error) {
	var employees []Employee
	if err := tx.Where("employeeid = ?", employeeid).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}