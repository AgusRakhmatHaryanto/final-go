package models

import "time"

type Movies struct {
	ID        int      `gorm:"primaryKey;type:int"`
	Title     string   `gorm:"type:varchar(255);not null"`
	Year      int      `gorm:"type:int"`
	AwardID   int      `gorm:"type:int"`
	GenreID   int      `gorm:"type:int"`
	Awards    *Award `gorm:"foreignKey:AwardID;references:ID"`
	Genres    *Genre `gorm:"foreignKey:GenreID;references:ID"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
