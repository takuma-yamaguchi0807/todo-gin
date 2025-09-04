package todo

import (
	"fmt"

	"github.com/google/uuid"
)

type Id struct {
	value uuid.UUID
}

func NewId(v string) (Id, error){
	id, err := uuid.Parse(v)
	if err != nil {
		return Id{}, fmt.Errorf("invalid format: %w",err)
	}
	return Id{value: id}, nil
}