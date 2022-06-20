package repository

import (
	"context"

	"github.com/michaelact/firstApi/model/domain"
)

type TodoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
