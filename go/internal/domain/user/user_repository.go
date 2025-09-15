package user

// Repository はユーザ関連の永続化を扱う抽象です。
// 最小構成として、ユーザ保存と、メール+パスワードハッシュでの存在確認を提供します。
type Repository interface{
    // Save はユーザを新規保存する。
    Save(u User) error
    // ExistsByEmailAndPassword は、メール+パスワード（平文入力）に一致するユーザが存在するかを返す。
    // 検証（ハッシュ照合）はDB拡張機能で実施する想定。
    ExistsByEmailAndPassword(email Email, password Password) (bool, error)
}