package models

import "time"

type Award struct {
	ID         int      `gorm:"primaryKey;type:int"`
	Title      string   `gorm:"type:varchar(255);not null"`
	Year       int      `gorm:"type:int"`
	UpdatedAt  time.Time
	CreatedAt  time.Time
}