package models

import "time"

// Users table
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Role      string    `gorm:"type:enum('admin', 'mentor', 'siswa')" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
