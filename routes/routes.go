package routes

import (
	"github.com/MohamedShehata15/todo_app/routes/api"
	"github.com/gin-gonic/gin"
)

func IndexRoutes(superRoute *gin.RouterGroup) {
	api.UserRoutes(superRoute)
	api.TodoRoutes(superRoute)
}