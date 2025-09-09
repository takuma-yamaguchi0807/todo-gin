package usecase

import (
    "context"

    "github.com/google/uuid"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
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

// Execute はトランザクション内で Todo を1件保存するダミー実装。
// 実際の引数やDTO設計は省略し、最小限に留める。
func (uc *TodoCreateUsecase) Execute(ctx context.Context) error {
    // 最小構成として適当な値オブジェクトを生成
    uid, _ := user.NewId(uuid.NewString())
    id, _ := todo.NewId(uuid.NewString())
    title, _ := todo.NewTitle("サンプルタイトル")
    desc, _ := todo.NewDescription("サンプル説明")
    status, _ := todo.NewStatus(string(todo.Wait))
    due, _ := todo.NewDueDate("")

    t := todo.NewTodo(id, uid, title, desc, status, due)

    // database/sql のシンプルなパス：トランザクションは使わず直接保存（必要なら別途 BeginTx で実装）
    return uc.repo.Save(ctx, &t)
}
