package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/middleware"
)

// SetupRoutes registers health, auth, and todos endpoints and binds handlers.
func SetupRoutes(r *gin.Engine, tc *controller.TodoController, uc *controller.UserController, ts domainauth.TokenService) {
    // Health check
    r.GET("/healthz", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // Fallback handlers
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    "NOT FOUND",
            "message": "resource not found",
        })
    })

    r.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{
            "code":    "METHOD NOT ALLOWED",
            "message": "http method not allowed",
        })
    })

    // Auth
    auth := r.Group("/auth")
    {
        // ユーザ登録
        auth.POST("/signup", uc.Signup)
        // ログイン
        auth.POST("/login", uc.Login)
    }
    
    todos := r.Group("/todos")
    // 認証必須: すべてのTODO APIでJWT検証を最初に実施（ボディ検証より前）
    todos.Use(middleware.AuthRequired(ts))
    {
        // 一覧（ログインユーザのスコープで返却想定）
        todos.GET("", tc.Get)
        // 詳細
        todos.GET("/:id", tc.Detail)
        // 作成
        todos.POST("", tc.Create)
        // 更新（全項目更新）
        todos.PUT("/:id", tc.Update)
        // 削除（JSON body で id 配列を受け付ける）
        todos.DELETE("", tc.Delete)
    }
}
