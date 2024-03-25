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

		genreAuthRole := genreRouter.Group("", middleware.IsRole("admin"))
		genreAuthRole.POST("", genreController.CreateGenre)
		genreAuthRole.PUT("/:id",  genreController.UpdateGenre)
		genreAuthRole.DELETE("/:id", genreController.DeleteGenre)
	}

	awardRouter := router.Group("awards")
	{
		awardRouter.GET("", middleware.AuthMiddleware(userRepository), awardController.FindAllAwards)
		awardRouter.GET("/:id", middleware.AuthMiddleware(userRepository), awardController.FindAwardById)

		awardAuthRole := awardRouter.Group("",middleware.AuthMiddleware(userRepository), middleware.IsRole("admin"))
		awardAuthRole.POST("", awardController.CreateAward)
		awardAuthRole.PUT("/:id", awardController.UpdateAward)
		awardAuthRole.DELETE("/:id", awardController.DeleteAward)
	}

	directorRouter := router.Group("directors")
	{
		directorRouter.GET("", middleware.AuthMiddleware(userRepository), directorController.FindAllDirectors)
		directorRouter.GET("/:id", middleware.AuthMiddleware(userRepository), directorController.FindDirectorById)

		directorAuthRole := directorRouter.Group("",middleware.AuthMiddleware(userRepository), middleware.IsRole("admin"))
		directorAuthRole.POST("",  directorController.CreateDirector)
		directorAuthRole.PUT("/:id", directorController.UpdateDirector)
		directorAuthRole.DELETE("/:id", directorController.DeleteDirector)
	}

	movieRouter := router.Group("movies")
	{
		movieRouter.GET("", middleware.AuthMiddleware(userRepository), movieController.FindAllMovies)
		movieRouter.GET("/:id", middleware.AuthMiddleware(userRepository), movieController.FindMovieById)

		movieAuthRole := movieRouter.Group("",middleware.AuthMiddleware(userRepository), middleware.IsRole("admin"))
		movieAuthRole.POST("", movieController.CreateMovie)
		movieAuthRole.PUT("/:id", movieController.UpdateMovie)
		movieAuthRole.DELETE("/:id", movieController.DeleteMovie)
	}
	return service
}
