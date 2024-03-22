package models

import "time"

type Director struct {
	ID        int    `gorm:"primaryKey;type:int"`
	Name      string `gorm:"type:varchar(255);not null"`
	MovieID   int    `gorm:"type:int"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
