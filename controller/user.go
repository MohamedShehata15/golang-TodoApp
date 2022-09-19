package controller

import (
	"net/http"

	service "github.com/MohamedShehata15/todo_app/services"
	"github.com/MohamedShehata15/todo_app/utils"
	"github.com/MohamedShehata15/todo_app/validator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(c *gin.Context) {
	// Validate Input
	var input validator.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.Password = string(hashedPassword)

	user, err := service.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(user.Id)

	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"token": token,
	})
}

func Login(c *gin.Context) {
	var input validator.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, _ := service.LoginUser(input)
	token, _ := utils.GenerateToken(user.Id)

	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"token": token,
	})
}