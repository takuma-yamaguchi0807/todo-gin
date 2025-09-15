package usecase

import (
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