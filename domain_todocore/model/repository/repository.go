package repository

import (
	"context"

	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/entity"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error)
}
