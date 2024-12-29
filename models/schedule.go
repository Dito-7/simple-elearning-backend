package models

import "time"

type Schedule struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ClassID      uint      `gorm:"not null" json:"class_id"`
	MentorID     uint      `gorm:"not null" json:"mentor_id"`
	ScheduleDate time.Time `json:"schedule_date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ScheduleResponse struct {
	ID           uint      `json:"id"`
	ClassID      uint      `json:"class_id"`
	MentorID     uint      `json:"mentor_id"`
	ScheduleDate time.Time `json:"schedule_date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}
