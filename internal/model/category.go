package model

import "time"

type Category struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	groupId     int
	CreatedAt   time.Time
}
