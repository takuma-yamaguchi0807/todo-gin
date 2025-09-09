package db

import (
    "context"
    "database/sql"
    "errors"

    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

// TodoRepositoryImpl is a minimal database/sql implementation.
// Query details can be extended as needed.
type TodoRepositoryImpl struct{
    db *sql.DB
}

func NewTodoRepositoryImpl(db *sql.DB) todo.TodoRepository {
    return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, t *todo.Todo) error {
    return errors.New("TodoRepositoryImpl.Save not implemented")
}

func (r *TodoRepositoryImpl) FindById(ctx context.Context, id todo.Id) (*todo.Todo, error) {
    return nil, errors.New("TodoRepositoryImpl.FindById not implemented")
}

func (r *TodoRepositoryImpl) FindByUser(ctx context.Context, userId user.Id) ([]*todo.Todo, error) {
    // Postgres is assumed. Table `todo` has the same columns as the aggregate root.
    // Map to domain (value objects) here in the repository layer.
    const q = `
        SELECT
            id::text,
            user_id::text,
            title,
            description,
            status,
            to_char(due_date, 'YYYY-MM-DD')
        FROM todo
        WHERE user_id = $1
    `

    rows, err := r.db.QueryContext(ctx, q, userId.UUID())
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    res := make([]*todo.Todo, 0)
    for rows.Next() {
        t, err := r.scanTodo(rows)
        if err != nil {
            return nil, err
        }
        res = append(res, t)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return res, nil
}

// rowScanner は *sql.Row / *sql.Rows いずれにも共通の Scan をまとめるための最小インタフェースです。
// これにより 1件取得/複数件取得の両方で同じ詰め替え処理を再利用できます。
type rowScanner interface {
    Scan(dest ...any) error
}

// scanTodo は 1 行分のレコードをドメインの Todo 集約へ詰め替えます。
// - NULL になり得るカラム（description, due_date）は sql.NullString で受けてから VO に変換します。
// - UUID/日付/タイトル/ステータスなどのバリデーションは各 VO のコンストラクタに委譲します。
func (r *TodoRepositoryImpl) scanTodo(rs rowScanner) (*todo.Todo, error) {
    var (
        idStr     string
        userStr   string
        titleStr  string
        descNS    sql.NullString
        statusStr string
        dueNS     sql.NullString
    )

    if err := rs.Scan(&idStr, &userStr, &titleStr, &descNS, &statusStr, &dueNS); err != nil {
        return nil, err
    }

    tid, err := todo.NewId(idStr)
    if err != nil { return nil, err }
    uid, err := user.NewId(userStr)
    if err != nil { return nil, err }
    ttl, err := todo.NewTitle(titleStr)
    if err != nil { return nil, err }
    var descStr string
    if descNS.Valid { descStr = descNS.String }
    desc, err := todo.NewDescription(descStr)
    if err != nil { return nil, err }
    st, err := todo.NewStatus(statusStr)
    if err != nil { return nil, err }
    var dueStr string
    if dueNS.Valid { dueStr = dueNS.String }
    due, err := todo.NewDueDate(dueStr)
    if err != nil { return nil, err }

    t := todo.NewTodo(tid, uid, ttl, desc, st, due)
    return &t, nil
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, id todo.Id) error {
    return errors.New("TodoRepositoryImpl.Update not implemented")
}

func (r *TodoRepositoryImpl) DeleteByIds(ctx context.Context, ids []todo.Id) error {
    return errors.New("TodoRepositoryImpl.DeleteByIds not implemented")
}
