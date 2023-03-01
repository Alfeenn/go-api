package service

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/api-go/exception"
	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/model"
	"github.com/Alfeenn/api-go/model/web"
	"github.com/Alfeenn/api-go/repository"
	"github.com/go-playground/validator"
)

type ServiceImpl struct {
	CategoryRepository repository.Repository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.Repository, DB *sql.DB, validate *validator.Validate) ModelService {
	return &ServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (service *ServiceImpl) Create(ctx context.Context, request web.RequestService) web.CategoryResponse {
	err := service.validate.Struct(request)
	helper.PanicIfErr(err)
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)

	category := model.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ConvertToModel(category)
}
func (service *ServiceImpl) Update(ctx context.Context, update web.UpdateRequest) web.CategoryResponse {
	err := service.validate.Struct(update)
	helper.PanicIfErr(err)
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.Find(ctx, tx, update.Id)
	if err != nil {
		panic(exception.NewNotFounErr(err.Error()))
	}
	category.Name = update.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ConvertToModel(category)
}
func (service *ServiceImpl) Delete(ctx context.Context, getId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.Find(ctx, tx, getId)
	if err != nil {
		panic(exception.NewNotFounErr(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)

}
func (service *ServiceImpl) Find(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.Find(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFounErr(err.Error()))
	}
	return helper.ConvertToModel(category)

}
func (service *ServiceImpl) FindAll(ctx context.Context, limit int, offset int) []web.CategoryResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfErr(err)

	category := service.CategoryRepository.FindAll(ctx, tx, limit, offset)
	var categryResponses []web.CategoryResponse
	for _, v := range category {
		categryResponses = append(categryResponses, helper.ConvertToModel(v))
	}
	return categryResponses

}
