package todo

import (
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

type Todo struct {
    id          Id
    userId      user.Id
    title       Title
    description Description
    status      Status
    dueDate     DueDate
}

// NewTodo creates a new Todo aggregate with validated value objects.
func NewTodo(id Id, userId user.Id, title Title, description Description, status Status, dueDate DueDate) Todo {
    return Todo{
        id:          id,
        userId:      userId,
        title:       title,
        description: description,
        status:      status,
        dueDate:     dueDate,
    }
}

// ID は集約のIDを返します。
func (t Todo) ID() Id { return t.id }

// UserID は所有者ユーザーIDを返します。
func (t Todo) UserID() user.Id { return t.userId }

// Title はタイトルの値オブジェクトを返します。
func (t Todo) Title() Title { return t.title }

// Description は説明の値オブジェクトを返します。
func (t Todo) Description() Description { return t.description }

// Status はステータスを返します。
func (t Todo) Status() Status { return t.status }

// DueDate は期限の値オブジェクトを返します。
func (t Todo) DueDate() DueDate { return t.dueDate }
