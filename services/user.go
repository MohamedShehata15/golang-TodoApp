package service

import (
	"log"

	"github.com/MohamedShehata15/todo_app/database"
	"github.com/MohamedShehata15/todo_app/model"
	"github.com/MohamedShehata15/todo_app/utils"
	"github.com/MohamedShehata15/todo_app/validator"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userData validator.CreateUserInput) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{Name: userData.Name, Email: userData.Email, Password: string(hashedPassword)}
	
	database.DB.Create(&user)

	return user, nil
}

func LoginUser(userData validator.LoginInput) (model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", userData.Email).First(&user).Error; err != nil {
		log.Fatal("Credentials are not correct")
	}

	if err := utils.VerifyPassword(user.Password, userData.Password); err != nil {
		log.Fatal("Credentials are not correct")
	}

	return user, nil
	
}