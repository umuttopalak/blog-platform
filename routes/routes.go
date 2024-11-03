package routes

import (
	"blog-platform/controllers"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}

	postRoutes := router.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	{
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:post_id", controllers.GetPost)
		postRoutes.PUT("/:post_id", controllers.UpdatePost)
		postRoutes.DELETE("/:post_id", controllers.DeletePost)

	}

	return router
}
