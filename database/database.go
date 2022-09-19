package database

import (
	"fmt"
	"os"

	"github.com/MohamedShehata15/todo_app/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


type DatabaseConfig struct {
	User string
	Password string
	Host string
	DBName string
}


func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbConfig := DatabaseConfig{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_Host"),
		DBName: os.Getenv("DB_NAME"),
	}


	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&model.User{}, &model.Todo{})
	DB = db
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	conn.Close()
}