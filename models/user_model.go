package models

import (
	"time"

)


type User struct {
	ID        int    `gorm:"primaryKey;type:int"`
	Username  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Role      string `gorm:"type:role_type"`
	Password  string `gorm:"type:varchar(255);not null" `
	UpdatedAt time.Time
	CreatedAt time.Time
}
