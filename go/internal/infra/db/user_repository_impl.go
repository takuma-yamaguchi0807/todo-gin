package db

import (
	"database/sql"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) user.Repository {
	return &UserRepositoryImpl{db: db}
}

// Save はユーザを新規保存する（ハッシュ化はDB拡張に委譲）。
func (r *UserRepositoryImpl) Save(u user.User) error {
	// TODO: implement
	return nil
}

// ExistsByEmailAndPassword はメール+平文パスワードの一致確認を行う。
// 実際の照合はDB拡張に委譲する想定。
func (r *UserRepositoryImpl) ExistsByEmailAndPassword(email user.Email, password user.Password) (bool, error) {
	// TODO: implement
	return false, nil
}