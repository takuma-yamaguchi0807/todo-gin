package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

type TodoController struct{
    getUC    *usecase.TodoGetUsecase
    createUC *usecase.TodoCreateUsecase
    updateUC *usecase.TodoUpdateUsecase
    deleteUC *usecase.TodoDeleteUsecase
}

func NewTodoController(getUC *usecase.TodoGetUsecase, createUC *usecase.TodoCreateUsecase, updateUC *usecase.TodoUpdateUsecase, deleteUC *usecase.TodoDeleteUsecase) *TodoController {
    return &TodoController{getUC: getUC, createUC: createUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (tc *TodoController) Get(c *gin.Context){
    // TODO: map query params and call tc.getUC.Execute()
    tc.getUC.Execute()
    c.JSON(http.StatusOK, gin.H{"items": []any{}})
}

func (tc *TodoController) Detail(c *gin.Context){
    // TODO: parse id and call tc.getUC.Execute()
    tc.getUC.Execute()
    c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
}

func (tc *TodoController) Create(c *gin.Context){
    // TODO: bind JSON and call tc.createUC.Execute()
    tc.createUC.Execute()
    c.JSON(http.StatusCreated, gin.H{"id": "created"})
}

func (tc *TodoController) Update(c *gin.Context){
    // TODO: bind JSON and call tc.updateUC.Execute()
    tc.updateUC.Execute()
    c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "status": "updated"})
}

func (tc *TodoController) Delete(c *gin.Context){
    // TODO: parse id and call tc.deleteUC.Execute()
    tc.deleteUC.Execute()
    c.Status(http.StatusNoContent)
}
