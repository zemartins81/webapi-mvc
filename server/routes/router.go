package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcmartins81/webapi-with-go/controllers"
	"github.com/jcmartins81/webapi-with-go/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		main.POST("login", controllers.Login)
	}

	return router
}
