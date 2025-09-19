package main

import (
	"encoding/json"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/config"
	infraauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/infra/auth"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/infra/db"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/controller"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/router"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/usecase"
)

func main() {
	r := gin.New()
	// リクエストID付与（ヘッダ優先、無ければ生成）
	r.Use(func(c *gin.Context) {
		rid := c.GetHeader("X-Request-Id")
		if rid == "" {
			rid = uuid.NewString()
		}
		c.Set("request_id", rid)
		c.Writer.Header().Set("X-Request-Id", rid)
		c.Next()
	})

	// 構造化アクセスログ（レベル/サービス/リクエストIDを含む1行JSON）
	serviceName := config.GetenvOrDefault("SERVICE_NAME", "api")
	r.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		level := "info"
		if p.StatusCode >= 500 {
			level = "error"
		} else if p.StatusCode >= 400 {
			level = "warn"
		}
		var requestID string
		if p.Keys != nil {
			if v, ok := p.Keys["request_id"]; ok {
				if s, sok := v.(string); sok {
					requestID = s
				}
			}
		}
		b, _ := json.Marshal(map[string]any{
			"time":       p.TimeStamp.Format(time.RFC3339Nano),
			"level":      level,
			"service":    serviceName,
			"request_id": requestID,
			"status":     p.StatusCode,
			"method":     p.Method,
			"path":       p.Path,
			"ip":         p.ClientIP,
			"latency_ms": p.Latency.Milliseconds(),
			"ua":         p.Request.UserAgent(),
		})
		return string(b) + "\n"
	}))
	r.Use(gin.Recovery())

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
	if err != nil {
		panic(err)
	}

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
