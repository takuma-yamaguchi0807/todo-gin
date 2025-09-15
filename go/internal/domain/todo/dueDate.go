package todo

import (
	"strings"
	"time"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

type DueDate struct {
    value *time.Time
}

func NewDueDate(s string) (DueDate,error){
    if strings.TrimSpace(s) == ""{
        return DueDate{}, nil
    }
    parsed, err := time.Parse("2006-01-02", s)
    if err != nil {
        return DueDate{}, common.InvalidErr("todo.due_date", "invalid date format. must be yyyy-mm-dd", err)
    }
    return DueDate{value: &parsed}, nil
}

// StringPtr は日付が設定されていれば "YYYY-MM-DD" の *string を返します。
// 未設定なら nil を返します。
func (d DueDate) StringPtr() *string {
    if d.value == nil {
        return nil
    }
    s := d.value.Format("2006-01-02")
    return &s
}
