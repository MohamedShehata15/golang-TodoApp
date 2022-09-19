package controller

import (
	"net/http"

	service "github.com/MohamedShehata15/todo_app/services"
	"github.com/MohamedShehata15/todo_app/utils"
	"github.com/MohamedShehata15/todo_app/validator"
	"github.com/gin-gonic/gin"
)



func CreateTodo(c *gin.Context) {
	// Validate input
	var input validator.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get Logged User
	userId,_ := utils.ExtractTokenID(c)
	todo, _ := service.CreateTodo(input, userId)

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func UpdateTodo(c *gin.Context) {
	loggedUser,_ := utils.ExtractTokenID(c)
	todo, err := service.UpdateTodo(c, loggedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func DeleteTodo(c *gin.Context) {
	loggedUser,_ := utils.ExtractTokenID(c)
	_, err := service.DeleteTodo(c, loggedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func FindTodos(c *gin.Context) {

	userId,_ := utils.ExtractTokenID(c)

	todos, err  := service.GetTodos(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}

func FindTodo(c *gin.Context) {
	loggedUser,_ := utils.ExtractTokenID(c)

	todo := service.GetTodo(c.Param("id"), loggedUser)

	if todo.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}