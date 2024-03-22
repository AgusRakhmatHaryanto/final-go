package models

import "time"

type Users struct {
	ID        int    `gorm:"primaryKey;type:int"`
	Username  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null" `
	UpdatedAt time.Time
	CreatedAt time.Time
}
