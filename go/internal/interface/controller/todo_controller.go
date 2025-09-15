package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

type TodoController struct{
    getUC    *usecase.TodoGetUsecase
    detailUC *usecase.TodoDetailUsecase
    createUC *usecase.TodoCreateUsecase
    updateUC *usecase.TodoUpdateUsecase
    deleteUC *usecase.TodoDeleteUsecase
}

func NewTodoController(getUC *usecase.TodoGetUsecase, detailUC *usecase.TodoDetailUsecase, createUC *usecase.TodoCreateUsecase, updateUC *usecase.TodoUpdateUsecase, deleteUC *usecase.TodoDeleteUsecase) *TodoController {
    return &TodoController{getUC: getUC, detailUC: detailUC, createUC: createUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (tc *TodoController) Get(c *gin.Context){
    // パスパラメータから userId を取得（検証はドメイン層で実施）
    req := dto.TodoGetRequest{UserID: c.Param("userId")}

    // ユースケースへリクエストを渡して実行
    items, err := tc.getUC.Execute(c.Request.Context(), req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    // トップレベルを配列にする
    c.JSON(http.StatusOK, items)
}

func (tc *TodoController) Detail(c *gin.Context){
    req := dto.TodoDetailRequest{ID: c.Param("id")}
    res, err := tc.detailUC.Execute(c.Request.Context(), req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusOK, res)
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
    //requestの内容をdtoにconvert
    var req dto.TodoDeleteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        appErr := common.InvalidErr("body", "invalid request body")
        status, payload := common.JSONFromError(appErr)
        c.JSON(status, payload)
        return
    }
    //execute実行
    err := tc.deleteUC.Execute(c.Request.Context(),req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusOK,nil)
}
