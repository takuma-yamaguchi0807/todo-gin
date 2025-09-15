package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/config"
	infraauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/infra/auth"
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

    // TokenService は buildDependency から返す設計に変えることも可能だが、
    // 現状は main 内で生成したものをここで再生成して渡すより、buildDependency の戻り値拡張が望ましい。
    // シンプルにするため、ここでは再取得する。
    secret := config.GetenvOrDefault("JWT_SECRET", "changeme-secret")
    issuer := config.GetenvOrDefault("JWT_ISSUER", "todo-gin")
    jwtGen := infraauth.NewHS256Generator(secret, issuer)
    router.SetupRoutes(r, tc, uc, jwtGen)

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
    detailUC := usecase.NewTodoDetailUsecase(todoRepo)
    createUC := usecase.NewTodoCreateUsecase(todoRepo)
    updateUC := usecase.NewTodoUpdateUsecase(todoRepo)
    deleteUC := usecase.NewTodoDeleteUsecase(todoRepo)
    todoController := controller.NewTodoController(getUC, detailUC, createUC, updateUC, deleteUC)

    // dependency user
    userRepo := db.NewUserRepositoryImpl(sdb)
    signupUC := usecase.NewUserSignupUsecase(userRepo)

    // JWT: 環境変数から取得（未設定時はローカル用デフォルト）
    secret := config.GetenvOrDefault("JWT_SECRET", "changeme-secret")
    issuer := config.GetenvOrDefault("JWT_ISSUER", "todo-gin")
    jwtGen := infraauth.NewHS256Generator(secret, issuer)
    loginUC := usecase.NewUserLoginUsecase(userRepo, jwtGen)
    userController := controller.NewUserController(signupUC, loginUC)

    return todoController, userController
}
