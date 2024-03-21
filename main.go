package main

import (
	"final-project/config"
	"final-project/controllers"
	"final-project/helper"
	"final-project/models"
	"final-project/repository"
	"final-project/routers"
	"final-project/services"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database Connection
	database := config.InitDB(&loadConfig)
	validate := validator.New()

	database.Table("users").AutoMigrate(&models.Users{})

	//init repository
	userRepository := repository.NewUserRepositoryImpl(database)

	//init service
	authenticationService := services.NewAuthenticationServiceImpl(userRepository, validate)

	//init controller
	authenticationController := controllers.NewAuthenticationController(authenticationService)

	//init router
	router := routers.NewRouter(authenticationController)

	server := &http.Server{
		Addr:    ":" + loadConfig.ServerPort,
		Handler: router,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
