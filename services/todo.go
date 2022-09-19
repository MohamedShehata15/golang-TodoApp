package service

import (
	"github.com/MohamedShehata15/todo_app/database"
	"github.com/MohamedShehata15/todo_app/model"
	"github.com/MohamedShehata15/todo_app/validator"
	"github.com/gin-gonic/gin"
)

func CreateTodo(todoData validator.CreateTodoInput, userId uint) (model.Todo, error) {
	todo := model.Todo{Title: todoData.Title, Completed: todoData.Completed, UserId: userId}
	database.DB.Create(&todo)

	return todo, nil
}

func UpdateTodo(c *gin.Context, userId uint) (model.Todo, error) {
	var todo model.Todo
	
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userId).First(&todo).Error; err != nil {
		return todo, err
	}

	c.BindJSON(&todo)
	database.DB.Save(&todo)

	return todo, nil
}

func DeleteTodo (c *gin.Context, userId uint) (model.Todo, error) {

	var todo model.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userId).First(&todo).Error; err != nil {
		return todo, err
	}

	database.DB.Delete(&todo)
	return todo, nil
}

func GetTodos (userId uint) ([]model.Todo, error) {
	todos := make([]model.Todo, 0)
	database.DB.Where("user_id = ?", userId).Preload("User").Find(&todos)

	return todos, nil
}

func GetTodo (id string, userId uint) model.Todo {
	var todo model.Todo

	database.DB.Where("id = ? AND user_id=?", id, userId).Preload("User").Find(&todo)

	return todo
}