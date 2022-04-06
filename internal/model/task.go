package model

import "time"

type Task struct {
	Id           int `gorm:"primaryKey"`
	Title        string
	Description  string
	Duration     int // (in minutes)
	CategoryId   int
	ExternalLink string
	CreatedAt    time.Time
}
