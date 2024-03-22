package models

import "time"

type Genre struct {
	ID        int       `gorm:"primaryKey;type:int"`
	Name      string    `gorm:"type:varchar(255);not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
	Movies    Movies  `gorm:"foreignKey:GenreID"`
}