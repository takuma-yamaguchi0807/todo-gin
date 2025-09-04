package todo

import (
	"fmt"
	"strings"
	"time"
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
		return DueDate{}, fmt.Errorf("invalid date format. must be yyyy-mm-dd")
	}
	return DueDate{value: &parsed}, nil
}