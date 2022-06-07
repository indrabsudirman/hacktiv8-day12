package main

import (
	"hacktiv8-day12/controllers"
	"hacktiv8-day12/database"
	"hacktiv8-day12/middlewares"
	"hacktiv8-day12/repositories"
	"hacktiv8-day12/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()
	route := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService)

	sosmedRepo := repositories.NewSocialMediaRepository(db)
	sosmedService := services.NewSocialMediaService(sosmedRepo)
	sosmedController := controllers.NewSocialMediaController(sosmedService)

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userController.RegisterUser)
		userRoute.POST("/login", userController.LoginUser)
		userRoute.Use(middlewares.Auth())
		userRoute.PUT("/:userId", userController.UpdateUser)
		userRoute.DELETE("/:userId", userController.DeleteUser)
	}
	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(middlewares.Auth())
		photoRoute.POST("/", photoController.PostPhoto)
		photoRoute.GET("/", photoController.GetAllPhotos)
		photoRoute.PUT("/:photoId", photoController.UpdatePhoto)
		photoRoute.DELETE("/:photoId", photoController.DeletePhoto)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(middlewares.Auth())
		commentRoute.POST("/", commentController.PostComment)
		commentRoute.GET("/", commentController.GetAllComments)
		commentRoute.PUT("/:commentId", commentController.UpdateComment)
		commentRoute.DELETE("/:commentId", commentController.DeleteComment)
	}

	sosmedRoute := route.Group("/socialmedias")
	{
		sosmedRoute.Use(middlewares.Auth())
		sosmedRoute.POST("/", sosmedController.PostSocialMedia)
		sosmedRoute.GET("/", sosmedController.GetAllSocialMedias)
		sosmedRoute.PUT("/:socialMediaId", sosmedController.UpdateSocialMedia)
		sosmedRoute.DELETE("/:socialMediaId", sosmedController.DeleteSocialMedia)
	}

	route.Run(database.APP_PORT)

}
