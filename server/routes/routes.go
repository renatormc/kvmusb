package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/renatormc/kvmusb/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.GET("/test", controllers.Test)
		// books := main.Group("books")
		// {
		// 	books.GET("/", controllers.ShowAllBooks)
		// 	// books.GET("/:id", controllers.ShowBook)
		// 	// books.POST("/", controllers.CreateBook)
		// 	// books.PUT("/", controllers.EditBook)
		// 	// books.DELETE("/:id", controllers.DeleteBook)
		// }
	}

	return router
}
