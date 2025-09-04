package controller

import "github.com/gin-gonic/gin"

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (tc *TodoController) Get(c *gin.Context){
	//TODO
}

func (tc *TodoController) Detail(c *gin.Context){
	//TODO
}

func (tc *TodoController) Create(c *gin.Context){
	//TODO
}

func (tc *TodoController) Update(c *gin.Context){
	//TODO
}

func (tc *TodoController) Delete(c *gin.Context){
	//TODO
}