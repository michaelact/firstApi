package service

import(
	"context"

	"github.com/michaelact/firstApi/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}
