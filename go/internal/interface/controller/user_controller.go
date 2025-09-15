package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
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
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, payload := common.JSON(common.Invalid, "body", "invalid request body")
		c.JSON(status, payload)
		return
	}
	if err := uc.signupUC.Execute(c.Request.Context(), req.Email, req.Password); err != nil {
		status, payload := common.JSONFromError(err)
		c.JSON(status, payload)
		return
	}
	c.Status(http.StatusCreated)
}

func (uc *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, payload := common.JSON(common.Invalid, "body", "invalid request body")
		c.JSON(status, payload)
		return
	}
	res, err := uc.loginUC.Execute(c.Request.Context(), req)
	if err != nil {
		status, payload := common.JSONFromError(err)
		c.JSON(status, payload)
		return
	}
	c.JSON(http.StatusOK, res)
}