package db

import (
	"database/sql"
	"errors"

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
	if u.Email().String() == "" || u.Password().String() == "" {
		return errors.New("empty email or password")
	}
	_, err := r.db.Exec(
		`INSERT INTO users (id, email, password_hash)
			VALUES (gen_random_uuid(), $1, crypt($2, gen_salt('bf')))` ,
		u.Email().String(),
		u.Password().String(),
	)
	return err
}

// FindIdByEmailAndPassword はメール+平文パスワードの一致確認を行い、IDを返す。
// 実際の照合は pgcrypto の crypt を利用。
func (r *UserRepositoryImpl) FindIdByEmailAndPassword(email user.Email, password user.Password) (user.Id, bool, error) {
	var idStr string
	err := r.db.QueryRow(
		`SELECT id FROM users WHERE email = $1 AND password_hash = crypt($2, password_hash) LIMIT 1`,
		email.String(), password.String(),
	).Scan(&idStr)
	if err == sql.ErrNoRows {
		return user.Id{}, false, nil
	}
	if err != nil {
		return user.Id{}, false, err
	}
	uid, uerr := user.NewId(idStr)
	if uerr != nil {
		return user.Id{}, false, uerr
	}
	return uid, true, nil
}