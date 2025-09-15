package usecase

import (
	"context"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

// UserSignupUsecase はユーザ登録のユースケースです。
// 具体的な依存関係や処理は後で実装します。
type UserSignupUsecase struct{
	repo user.Repository
}

// NewUserSignupUsecase は UserSignupUsecase のコンストラクタです。
func NewUserSignupUsecase(repo user.Repository) *UserSignupUsecase {
    return &UserSignupUsecase{repo: repo}
}

// Execute は email と password を受け取り、ユーザを新規作成して保存します。
// 返り値は今は不要のため error のみ返します。
func (uc *UserSignupUsecase) Execute(ctx context.Context, emailStr, passwordStr string) error {
    email, err := user.NewEmail(emailStr)
    if err != nil {
        return err
    }
    password, err := user.NewPassword(passwordStr)
    if err != nil {
        return err
    }
    // ID は最小構成ではDB側で生成する前提のため空のまま渡す（INSERT側で生成）
    u := user.NewUser(user.Id{}, email, password)
    return uc.repo.Save(u)
}