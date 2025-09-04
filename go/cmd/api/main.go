package main

import (
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/router"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

func main() {
    r := gin.Default()

    // Build controller and dependencies (simple wiring function)
    tc := buildTodoController()

	// CORS設定（フロントから叩けるように）
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.jsのURL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    router.SetupRoutes(r, tc)

	// サーバー起動
	r.Run(":8080")
}

// buildTodoController composes todo usecases and controller without scattering news.
func buildTodoController() *controller.TodoController {
    getUC := usecase.NewTodoGetUsecase()
    createUC := usecase.NewTodoCreateUsecase()
    updateUC := usecase.NewTodoUpdateUsecase()
    deleteUC := usecase.NewTodoDeleteUsecase()
    return controller.NewTodoController(getUC, createUC, updateUC, deleteUC)
}
