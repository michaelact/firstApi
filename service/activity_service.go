package service

import(
	"context"

	"github.com/michaelact/firstApi/model/web"
)

type Activity interface {
	Create(ctx context.Context, request web.ActivityCreateRequest) web.ActivityResponse
	Update(ctx context.Context, request web.ActivityUpdateRequest) web.ActivityResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.ActivityResponse
	FindAll(ctx context.Context) []web.ActivityResponse
}
