package repository

import (
	"context"

	"github.com/michaelact/firstApi/model/domain"
)

type ActivityRepository interface {
	Save(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity
	Update(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Activity, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Activity
}
