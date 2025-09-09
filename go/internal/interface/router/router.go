package router

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
)

// SetupRoutes registers health and todos endpoints and binds handlers.
func SetupRoutes(r *gin.Engine, tc *controller.TodoController) {
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

    // Todos
    todos := r.Group("/todos")
    {
        // ユーザーIDでTodo一覧を取得（パスパラメータ）
        todos.GET("/users/:userId", tc.Get)
        todos.GET(":id", tc.Detail)
        todos.POST("", tc.Create)
        todos.PUT(":id", tc.Update)
        todos.DELETE(":id", tc.Delete)
    }
}
