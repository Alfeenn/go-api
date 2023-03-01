package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Alfeenn/api-go/app"
	"github.com/Alfeenn/api-go/controller"
	"github.com/Alfeenn/api-go/exception"
	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/middleware"
	"github.com/Alfeenn/api-go/repository"
	"github.com/Alfeenn/api-go/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	router := httprouter.New()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	CategoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryControler := controller.NewCategoryController(CategoryService)

	router.GET("/api/categories", categoryControler.FindAll)
	router.GET("/api/categories/:categoryId", categoryControler.Find)
	router.POST("/api/categories/", categoryControler.Create)
	router.PUT("/api/categories/:categoryId", categoryControler.Update)
	router.DELETE("/api/categories/:categoryId", categoryControler.Delete)
	router.PanicHandler = exception.ErrHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
