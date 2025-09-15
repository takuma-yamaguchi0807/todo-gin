package user

// Repository はユーザ関連の永続化を扱う抽象です。
// 最小構成として、ユーザ保存と、メール+パスワードでの認証（ID取得）を提供します。
type Repository interface{
	// Save はユーザを新規保存する。
	Save(u User) error
	// FindIdByEmailAndPassword は、メール+平文パスワードに一致するユーザIDを返します。
	// 一致しない場合は (Id{}, false, nil) を返します。
	FindIdByEmailAndPassword(email Email, password Password) (Id, bool, error)
}