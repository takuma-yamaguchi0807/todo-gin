package user

// User はユーザの集約ルート。
// 最小構成として ID/Email/PasswordHash のみ保持する。
type User struct {
    id       Id
    email    Email
    password Password
}

// NewUser はユーザを生成する。
func NewUser(id Id, email Email, password Password) User {
    return User{id: id, email: email, password: password}
}

// ID はユーザIDを返す。
func (u User) ID() Id { return u.id }

// Email はメールアドレスを返す。
func (u User) Email() Email { return u.email }

// Password は平文パスワード（入力値）を返す。
// ハッシュ化はDB側の拡張で実施する想定。
func (u User) Password() Password { return u.password }