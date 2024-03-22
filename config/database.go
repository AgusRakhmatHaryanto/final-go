package config

import (
	"final-project/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)



	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}