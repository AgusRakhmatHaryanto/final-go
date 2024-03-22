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
	database.Table("genres").AutoMigrate(&models.Genre{})
	database.Table("awards").AutoMigrate(&models.Award{})
	database.Table("movies").AutoMigrate(&models.Movies{})
	database.Table("directors").AutoMigrate(&models.Director{})

	//init repository
	userRepository := repository.NewUserRepositoryImpl(database)
	genreRepository := repository.NewGenreRepositoryImpl(database)
	awayRepository := repository.NewAwardRepositoryImpl(database)
	movieRepository := repository.NewMovieRepositoryImpl(database)
	directorRepository := repository.NewDirectorRepositoryImpl(database)

	//init service
	authenticationService := services.NewAuthenticationServiceImpl(userRepository, validate)
	userService := services.NewUsersServiceImpl(userRepository, validate)
	genreService := services.NewGenreServiceImpl(genreRepository, validate)
	awardService := services.NewAwardServiceImpl(awayRepository, validate)
	movieService := services.NewMovieServiceImpl(movieRepository, validate)
	directorService := services.NewDirectorServiceImpl(directorRepository, validate)


	//init controller
	authenticationController := controllers.NewAuthenticationController(authenticationService)
	userController := controllers.NewUsersController(userService)
	genreController := controllers.NewGenreController(genreService)

	awardController := controllers.NewAwardController(awardService)
	movieController := controllers.NewMovieController(movieService)
	directorController := controllers.NewDirectorController(directorService)


	//init router
	router := routers.NewRouter(
		userRepository,
		directorController,
		movieController,
		awardController,
		genreController,
		authenticationController,
		userController,
	)	
	server := &http.Server{
		Addr:    ":" + loadConfig.ServerPort,
		Handler: router,
		
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
