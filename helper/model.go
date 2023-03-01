package helper

import (
	"github.com/Alfeenn/api-go/model"
	"github.com/Alfeenn/api-go/model/web"
)

func ConvertToModel(category model.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
