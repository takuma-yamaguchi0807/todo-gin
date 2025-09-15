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
    // JWTのクレームから userId を取得
    claims, ok := common.ClaimsFromContext(c)
    if !ok {
        status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing claims"))
        c.JSON(status, payload)
        return
    }
    // リクエストDTOを経由してユースケースへ
    req := dto.TodoGetRequest{UserID: claims.UserIDString()}

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
    claims, ok := common.ClaimsFromContext(c)
    if !ok {
        status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing claims"))
        c.JSON(status, payload)
        return
    }
    req := dto.TodoDetailRequest{ID: c.Param("id"), UserID: claims.UserIDString()}
    res, err := tc.detailUC.Execute(c.Request.Context(), req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusOK, res)
}

func (tc *TodoController) Create(c *gin.Context){
    claims, ok := common.ClaimsFromContext(c)
    if !ok {
        status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing claims"))
        c.JSON(status, payload)
        return
    }
    var req dto.TodoCreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        appErr := common.InvalidErr("body", "invalid request body")
        status, payload := common.JSONFromError(appErr)
        c.JSON(status, payload)
        return
    }
    req.UserID = claims.UserIDString()
    res, err := tc.createUC.Execute(c.Request.Context(), req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusCreated, res)
}

func (tc *TodoController) Update(c *gin.Context){
    claims, ok := common.ClaimsFromContext(c)
    if !ok {
        status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing claims"))
        c.JSON(status, payload)
        return
    }
    var req dto.TodoUpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        appErr := common.InvalidErr("body", "invalid request body")
        status, payload := common.JSONFromError(appErr)
        c.JSON(status, payload)
        return
    }
    req.ID = c.Param("id")
    req.UserID = claims.UserIDString()
    if err := tc.updateUC.Execute(c.Request.Context(), req); err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusOK, gin.H{"id": req.ID, "status": "updated"})
}

func (tc *TodoController) Delete(c *gin.Context){
    // JWT のクレームから userId を取得
    claims, ok := common.ClaimsFromContext(c)
    if !ok {
        status, payload := common.JSONFromError(common.New(common.Unauthorized, "authorization", "missing claims"))
        c.JSON(status, payload)
        return
    }

    //requestの内容をdtoにconvert
    var req dto.TodoDeleteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        appErr := common.InvalidErr("body", "invalid request body")
        status, payload := common.JSONFromError(appErr)
        c.JSON(status, payload)
        return
    }
    // 所有者情報を付与
    req.UserID = claims.UserIDString()

    //execute実行
    err := tc.deleteUC.Execute(c.Request.Context(),req)
    if err != nil {
        status, payload := common.JSONFromError(err)
        c.JSON(status, payload)
        return
    }
    c.JSON(http.StatusOK,nil)
}
