package models

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	ID int `gorm:"primaryKey" json:"id"`
	Employeeid string `json:"employeeid"`
	Name string `json:"name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Model[Employee] `gorm:"-" json:"-"`
}

func (Employee) GetEmployeesByEmployeeid(tx *gorm.DB, employeeid string) ([]Employee, error) {
	var employees []Employee
	if err := tx.Where("employeeid = ?", employeeid).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}