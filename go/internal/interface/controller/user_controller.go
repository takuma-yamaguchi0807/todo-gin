package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

type UserController struct {
	signupUC *usecase.UserSignupUsecase
	loginUC  *usecase.UserLoginUsecase
}

func NewUserController(signupUC *usecase.UserSignupUsecase, loginUC *usecase.UserLoginUsecase) *UserController {
	return &UserController{signupUC: signupUC, loginUC: loginUC}
}

func (uc *UserController) Signup(c *gin.Context) {
	//TODO: implement
}

func (uc *UserController) Login(c *gin.Context) {
	//TODO: implement
}