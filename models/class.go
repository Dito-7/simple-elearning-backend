package models

import "time"

type Class struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
