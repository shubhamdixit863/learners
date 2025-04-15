package models

import "time"

type Products struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Quantity  uint
	Color     string
}
