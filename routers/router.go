package routers

import (
	"final-project/controllers"
	"final-project/middleware"
	"final-project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository,
	directorController *controllers.DirectorsController,
	movieController *controllers.MovieController,
	awardController *controllers.AwardController,
	genreController *controllers.GenreController,
	authenticationController *controllers.AuthenticationController, userController *controllers.UsersController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome Home"})
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "failed",
			"message": "Page Not Found",
		})
	})

	router := service.Group("api/v1")

	authenticationRouter := router.Group("auth")
	{
		authenticationRouter.POST("login", authenticationController.Login)
		authenticationRouter.POST("register", authenticationController.Register)
	}

	userRouter := router.Group("users", middleware.AuthMiddleware(userRepository))
	{
		roleAuth := userRouter.Group("", middleware.IsRole("admin"))
		roleAuth.GET("", userController.FindAllUsers)
		roleAuth.GET("/email/:email", userController.GetUserByEmail)
		roleAuth.GET("/username/:username", userController.GetUserByUsername)
		roleAuth.GET("/:id", userController.FindUserById)
		roleAuth.POST("", userController.CreateUser)
		roleAuth.PUT("/:id", userController.UpdateUser)
		roleAuth.DELETE("/:id", userController.DeleteUser)
	}
	genreRouter := router.Group("genres")
	{
		genreRouter.GET("", middleware.AuthMiddleware(userRepository), genreController.FindAllGenres)
		genreRouter.GET("/:id", middleware.AuthMiddleware(userRepository), genreController.FindGenreById)
		genreRouter.POST("", middleware.AuthMiddleware(userRepository), genreController.CreateGenre)
		genreRouter.PUT("/:id", middleware.AuthMiddleware(userRepository), genreController.UpdateGenre)
		genreRouter.DELETE("/:id", middleware.AuthMiddleware(userRepository), genreController.DeleteGenre)
	}

	awardRouter := router.Group("awards")
	{
		awardRouter.GET("", middleware.AuthMiddleware(userRepository), awardController.FindAllAwards)
		awardRouter.GET("/:id", middleware.AuthMiddleware(userRepository), awardController.FindAwardById)
		awardRouter.POST("", middleware.AuthMiddleware(userRepository), awardController.CreateAward)
		awardRouter.PUT("/:id", middleware.AuthMiddleware(userRepository), awardController.UpdateAward)
		awardRouter.DELETE("/:id", middleware.AuthMiddleware(userRepository), awardController.DeleteAward)
	}

	directorRouter := router.Group("directors")
	{
		directorRouter.GET("", middleware.AuthMiddleware(userRepository), directorController.FindAllDirectors)
		directorRouter.GET("/:id", middleware.AuthMiddleware(userRepository), directorController.FindDirectorById)
		directorRouter.POST("", middleware.AuthMiddleware(userRepository), directorController.CreateDirector)
		directorRouter.PUT("/:id", middleware.AuthMiddleware(userRepository), directorController.UpdateDirector)
		directorRouter.DELETE("/:id", middleware.AuthMiddleware(userRepository), directorController.DeleteDirector)
	}

	movieRouter := router.Group("movies")
	{
		movieRouter.GET("", middleware.AuthMiddleware(userRepository), movieController.FindAllMovies)
		movieRouter.GET("/:id", middleware.AuthMiddleware(userRepository), movieController.FindMovieById)
		movieRouter.POST("", middleware.AuthMiddleware(userRepository), movieController.CreateMovie)
		movieRouter.PUT("/:id", middleware.AuthMiddleware(userRepository), movieController.UpdateMovie)
		movieRouter.DELETE("/:id", middleware.AuthMiddleware(userRepository), movieController.DeleteMovie)
	}
	return service
}
