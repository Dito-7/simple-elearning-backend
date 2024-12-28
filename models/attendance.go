package models

import "time"

type Attendance struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	ClassID   uint      `gorm:"not null" json:"class_id"`
	Status    string    `gorm:"type:enum('hadir', 'tidak hadir', 'izin')" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AttendanceResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	ClassID uint   `json:"class_id"`
	Status  string `json:"status"`
}
