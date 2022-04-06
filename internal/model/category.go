package model

import "time"

type Category struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	CreatedAt   time.Time
}
