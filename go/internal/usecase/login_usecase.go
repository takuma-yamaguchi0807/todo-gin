package usecase

import (
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

// UserLoginUsecase はユーザログインのユースケースです。
// 具体的な依存関係や処理は後で実装します。
type UserLoginUsecase struct{
	repo user.Repository
}

// NewUserLoginUsecase は UserLoginUsecase のコンストラクタです。
func NewUserLoginUsecase(repo user.Repository) *UserLoginUsecase {
    return &UserLoginUsecase{repo: repo}
}