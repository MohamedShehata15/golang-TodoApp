package main

import (
	"os"

	"github.com/MohamedShehata15/todo_app/database"
	"github.com/MohamedShehata15/todo_app/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = database.DatabaseConnection()

func main() {
	defer database.CloseDatabaseConnection(db)

	app := gin.New()

	router := app.Group("/api")

	routes.IndexRoutes(router)

	app.Run(os.Getenv("LOCAL_HOST") + ":" + os.Getenv("SERVER_PORT"))
}