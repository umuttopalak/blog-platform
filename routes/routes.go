package routes

import (
	"blog-platform/controllers"
	_ "blog-platform/docs"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		postRoutes.DELETE("/:post_id", controllers.RemovePost)
	}

	commentRoutes := router.Group("/comments")
	commentRoutes.Use(middleware.AuthMiddleware())
	{
		commentRoutes.GET("/user", controllers.GetCommentsByUser)
		commentRoutes.POST("/:post_id", controllers.CreateComment)
		commentRoutes.GET("/post/:post_id", controllers.GetCommentsByPost)
		commentRoutes.PUT("/:comment_id", controllers.UpdateComment)
		commentRoutes.DELETE("/:comment_id", controllers.RemoveComment)
	}

	reactionRoutes := router.Group("/reactions")
	reactionRoutes.Use(middleware.AuthMiddleware())
	{
		reactionRoutes.POST("/", controllers.AddReaction)
		reactionRoutes.GET("/post/:post_id", controllers.GetReactionsByPost)
		reactionRoutes.GET("/comment/:comment_id", controllers.GetReactionsByComment)
		reactionRoutes.DELETE("/:reaction_id", controllers.RemoveReaction)
	}

	return router
}
