package models

import "gorm.io/gorm"

type IALL[T any] interface {
	All(*gorm.DB) ([]T, error)
}

type Model[T any] struct {}

func (Model[T]) All(tx *gorm.DB) ([]T, error) {
	var values []T
	if err := tx.Find(&values).Error; err != nil {
		return nil, err
	}
	return values, nil
}

func (m Model[T]) Create(tx *gorm.DB, value T) error {
	return tx.Create(&value).Error
}

func (m Model[T]) Updates(tx *gorm.DB, value T) error {
	return tx.Updates(&value).Error
}