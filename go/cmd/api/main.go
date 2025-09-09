package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/app/config"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/infra/db"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/router"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

func main() {
    r := gin.Default()

    // Build controller and dependencies (simple wiring function)
    tc := buildDependency()

	// CORS設定（フロントから叩けるように）
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.jsのURL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    router.SetupRoutes(r, tc)

	// サーバー起動
	r.Run(":8080")
}

func buildDependency() *controller.TodoController {
    // DB接続情報を読み込み、Postgresへ接続
    cfg, _ := config.Load()
    sdb, err := cfg.OpenSQL()
    if err != nil { panic(err) }

    // TxManager と Repository を初期化
    
    repo := db.NewTodoRepositoryImpl(sdb)

    getUC := usecase.NewTodoGetUsecase(repo)
    createUC := usecase.NewTodoCreateUsecase(repo)
    updateUC := usecase.NewTodoUpdateUsecase(repo)
    deleteUC := usecase.NewTodoDeleteUsecase(repo)
    return controller.NewTodoController(getUC, createUC, updateUC, deleteUC)
}
