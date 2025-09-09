package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/app/apperror"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
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
    // パスパラメータから userId を取得（検証はドメイン層で実施）
    req := dto.TodoGetRequest{UserID: c.Param("userId")}

    // ユースケースへリクエストを渡して実行
    items, err := tc.getUC.Execute(c.Request.Context(), req)
    if err != nil {
        c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"items": items})
}

func (tc *TodoController) Detail(c *gin.Context){
    // TODO: id パラメータでの個別取得用のUCに差し替える
    c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
}

func (tc *TodoController) Create(c *gin.Context){
    // TODO: bind JSON and call tc.createUC.Execute()
    if err := tc.createUC.Execute(c.Request.Context()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": "created"})
}

func (tc *TodoController) Update(c *gin.Context){
    // TODO: bind JSON and call tc.updateUC.Execute()
    if err := tc.updateUC.Execute(c.Request.Context()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "status": "updated"})
}

func (tc *TodoController) Delete(c *gin.Context){
    // TODO: parse id and call tc.deleteUC.Execute()
    if err := tc.deleteUC.Execute(c.Request.Context()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
