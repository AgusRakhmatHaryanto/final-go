package models

import "time"

type Movie struct {
	ID        int    `gorm:"primaryKey;type:int"`
	Title     string `gorm:"type:varchar(255);not null"`
	Year      int    `gorm:"type:int"`
	AwardID   int    `gorm:"type:int"`
	GenreID   int    `gorm:"type:int"`
	UpdatedAt time.Time
	CreatedAt time.Time
	Directors Director `gorm:"foreignKey:MovieID"`
}
