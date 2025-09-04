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
