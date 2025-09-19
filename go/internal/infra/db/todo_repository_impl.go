package db

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
) // Query details can be extended as needed.
type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepositoryImpl(db *sql.DB) todo.TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, t *todo.Todo) error {
	const q = `
        INSERT INTO todos (
            id,
            user_id,
            title,
            description,
            status,
            due_date
        ) VALUES ($1, $2, $3, NULLIF($4, ''), $5, NULLIF($6, '')::date)
    `
	id := t.ID().UUID()
	uid := t.UserID().UUID()
	ttl := t.Title().String()
	descPtr := t.Description().Ptr()
	st := t.Status().String()
	duePtr := t.DueDate().StringPtr()

	var desc any
	if descPtr != nil {
		desc = *descPtr
	} else {
		desc = nil
	}
	var due any
	if duePtr != nil {
		due = *duePtr
	} else {
		due = nil
	}

	_, err := r.db.ExecContext(ctx, q, id, uid, ttl, desc, st, due)
	return err
}

func (r *TodoRepositoryImpl) FindById(ctx context.Context, id todo.Id) (*todo.Todo, error) {
	const q = `
        SELECT
            id::text,
            user_id::text,
            title,
            description,
            status,
            to_char(due_date, 'YYYY-MM-DD')
        FROM todos
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, q, id.UUID())
	t, err := r.scanTodo(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return t, nil
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
        FROM todos
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
	if err != nil {
		return nil, err
	}
	uid, err := user.NewId(userStr)
	if err != nil {
		return nil, err
	}
	ttl, err := todo.NewTitle(titleStr)
	if err != nil {
		return nil, err
	}
	var descStr string
	if descNS.Valid {
		descStr = descNS.String
	}
	desc, err := todo.NewDescription(descStr)
	if err != nil {
		return nil, err
	}
	st, err := todo.NewStatus(statusStr)
	if err != nil {
		return nil, err
	}
	var dueStr string
	if dueNS.Valid {
		dueStr = dueNS.String
	}
	due, err := todo.NewDueDate(dueStr)
	if err != nil {
		return nil, err
	}

	t := todo.NewTodo(tid, uid, ttl, desc, st, due)
	return &t, nil
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, t *todo.Todo) error {
	const q = `
        UPDATE todos SET
            title = $2,
            description = NULLIF($3, ''),
            status = $4,
            due_date = NULLIF($5, '')::date,
            updated_at = NOW()
        WHERE id = $1
    `
	id := t.ID().UUID()
	ttl := t.Title().String()
	descPtr := t.Description().Ptr()
	st := t.Status().String()
	duePtr := t.DueDate().StringPtr()

	var desc any
	if descPtr != nil {
		desc = *descPtr
	} else {
		desc = nil
	}
	var due any
	if duePtr != nil {
		due = *duePtr
	} else {
		due = nil
	}

	res, err := r.db.ExecContext(ctx, q, id, ttl, desc, st, due)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *TodoRepositoryImpl) DeleteByIds(ctx context.Context, ids []todo.Id) error {
	if len(ids) == 0 {
		return nil
	}
	placeholders := make([]string, 0, len(ids))
	args := make([]any, 0, len(ids))
	for i, id := range ids {
		placeholders = append(placeholders, "$"+strconv.Itoa(i+1))
		args = append(args, id.UUID())
	}
	q := "DELETE FROM todos WHERE id IN (" + join(placeholders, ",") + ")"
	_, err := r.db.ExecContext(ctx, q, args...)
	return err
}

func join(ss []string, sep string) string {
	if len(ss) == 0 {
		return ""
	}
	out := ss[0]
	for i := 1; i < len(ss); i++ {
		out += sep + ss[i]
	}
	return out
}
