package routes

import (
	"web-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBooks)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
