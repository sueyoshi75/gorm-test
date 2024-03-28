package models

type Todo struct {
	ID int `gorm:"primaryKey" json:"id"`
	Body string `json:"body"`
}