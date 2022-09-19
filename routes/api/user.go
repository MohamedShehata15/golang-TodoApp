package api

import (
	"github.com/MohamedShehata15/todo_app/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(superRoute *gin.RouterGroup) {
	userRoutes := superRoute.Group("/users")
	{
		userRoutes.POST("/register", controller.CreateUser)
		userRoutes.POST("/login", controller.Login)
	}
}