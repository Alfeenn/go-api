package service

import (
	"context"

	"github.com/Alfeenn/api-go/model/web"
)

type ModelService interface {
	Create(ctx context.Context, request web.RequestService) web.CategoryResponse
	Update(ctx context.Context, request web.UpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	Find(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context, limit int, offset int) []web.CategoryResponse
}
