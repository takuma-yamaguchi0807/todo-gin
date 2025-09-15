package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/config"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/infra/db"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/router"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

func main() {
    r := gin.Default()

    // Build controller and dependencies (simple wiring function)
    tc, uc := buildDependency()

	// CORS設定（フロントから叩けるように）
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.jsのURL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    router.SetupRoutes(r, tc, uc)

	// サーバー起動
	r.Run(":8080")
}

func buildDependency() (*controller.TodoController, *controller.UserController) {
    // DB接続情報を読み込み、Postgresへ接続
    cfg, _ := config.Load()
    sdb, err := cfg.OpenSQL()
    if err != nil { panic(err) }

    // dependency todo
    todoRepo := db.NewTodoRepositoryImpl(sdb)
    getUC := usecase.NewTodoGetUsecase(todoRepo)
    createUC := usecase.NewTodoCreateUsecase(todoRepo)
    updateUC := usecase.NewTodoUpdateUsecase(todoRepo)
    deleteUC := usecase.NewTodoDeleteUsecase(todoRepo)
    todoController := controller.NewTodoController(getUC, createUC, updateUC, deleteUC)

    // dependency user
    userRepo := db.NewUserRepositoryImpl(sdb)
    signupUC := usecase.NewUserSignupUsecase(userRepo)
    loginUC := usecase.NewUserLoginUsecase(userRepo)
    userController := controller.NewUserController(signupUC, loginUC)

    return todoController, userController
}
