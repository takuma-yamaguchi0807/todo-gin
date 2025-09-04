package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
)

func SetupRoutes(r *gin.Engine, tc *controller.TodoController){
		// ヘルスチェック
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	//エンドポイントがない場合
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404,gin.H{
			"code":"NOT FOUND",
			"message":"resource not found",
		})
	})
	//httpメソッドがない場合
	r.NoMethod(func(c *gin.Context) {
		c.JSON(405,gin.H{
			"code":"METHOD NOT ALLOWED",
			"message":"http method not allowed",
		})
	})

	todos := r.Group("/todos")
	{
		todos.GET("",tc.Get)
		todos.GET(":id",tc.Detail)
		todos.POST("",tc.Create)
		todos.DELETE("",tc.Delete)
		todos.PATCH("",tc.Update)
	}
}