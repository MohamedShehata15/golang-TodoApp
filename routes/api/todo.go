package api

import (
	"github.com/MohamedShehata15/todo_app/controller"
	"github.com/MohamedShehata15/todo_app/middleware"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(superRoute *gin.RouterGroup) {
	todoRoutes := superRoute.Group("/todos")
	{
		todoRoutes.Use(middleware.JwtAuth())
		todoRoutes.POST("/", controller.CreateTodo)
		todoRoutes.GET("/", controller.FindTodos)
		todoRoutes.GET("/:id", controller.FindTodo)
		todoRoutes.PATCH("/:id", controller.UpdateTodo)
		todoRoutes.DELETE("/:id", controller.DeleteTodo)
	}
}