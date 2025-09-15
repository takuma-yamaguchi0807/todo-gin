package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// TodoCreateUsecase は作成系のユースケース。
// ここでトランザクション境界（TxManager）を貼り、
// 内部でリポジトリを呼び出す最小構成にする。
type TodoCreateUsecase struct{
    repo todo.TodoRepository
}

func NewTodoCreateUsecase(repo todo.TodoRepository) *TodoCreateUsecase {
    return &TodoCreateUsecase{repo: repo}
}

// Execute は Todo を1件保存する。
// READMEの仕様に合わせ、title必須、description/status/due_dateは任意。
func (uc *TodoCreateUsecase) Execute(ctx context.Context, req dto.TodoCreateRequest) (dto.TodoDetailResponse, error) {
    // ID/ユーザID生成
    uid, err := user.NewId(req.UserID)
    if err != nil { return dto.TodoDetailResponse{}, err }
    id, err := todo.NewId(uuid.NewString())
    if err != nil { return dto.TodoDetailResponse{}, err }

    // 値オブジェクト生成
    title, err := todo.NewTitle(req.Title)
    if err != nil { return dto.TodoDetailResponse{}, err }
    var descStr string
    if req.Description != nil { descStr = *req.Description }
    desc, err := todo.NewDescription(descStr)
    if err != nil { return dto.TodoDetailResponse{}, err }
    st := string(todo.Wait)
    if req.Status != nil { st = *req.Status }
    status, err := todo.NewStatus(st)
    if err != nil { return dto.TodoDetailResponse{}, err }
    var dueStr string
    if req.DueDate != nil { dueStr = *req.DueDate }
    // READMEはISO8601想定。VO側は空文字も許容している実装なのでそのまま委譲。
    _ = time.Now() // keep import if needed later
    due, err := todo.NewDueDate(dueStr)
    if err != nil { return dto.TodoDetailResponse{}, err }

    t := todo.NewTodo(id, uid, title, desc, status, due)
    if err := uc.repo.Save(ctx, &t); err != nil { return dto.TodoDetailResponse{}, err }
    // 作成後の詳細レスポンスを返す（シンプルにドメイン→DTO変換）
    return dto.NewTodoDetailResponseFromDomain(t), nil
}
