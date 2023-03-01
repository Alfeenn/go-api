package controller

import (
	"net/http"
	"strconv"

	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/model/web"
	"github.com/Alfeenn/api-go/service"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	CategoryService service.ModelService
}

func NewCategoryController(CategoryController service.ModelService) CategoryController {
	return &UserController{
		CategoryService: CategoryController,
	}
}
func (controller *UserController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryRequest := web.RequestService{}
	helper.ReadRequestBody(r, &categoryRequest)

	CategoryResponse := controller.CategoryService.Create(r.Context(), categoryRequest)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteResponse(w, WebResponse)
}
func (controller *UserController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	UpdateRequest := web.UpdateRequest{}

	helper.ReadRequestBody(r, &UpdateRequest)
	req := params.ByName("categoryId")
	id, err := strconv.Atoi(req)
	helper.PanicIfErr(err)
	UpdateRequest.Id = id
	CategoryResponse := controller.CategoryService.Update(r.Context(), UpdateRequest)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteResponse(w, WebResponse)
}
func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := params.ByName("categoryId")
	id, err := strconv.Atoi(req)
	helper.PanicIfErr(err)

	controller.CategoryService.Delete(r.Context(), id)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteResponse(w, WebResponse)
}
func (controller *UserController) Find(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := params.ByName("categoryId")
	id, err := strconv.Atoi(req)
	helper.PanicIfErr(err)

	CategoryResponse := controller.CategoryService.Find(r.Context(), id)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteResponse(w, WebResponse)
}
func (controller *UserController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result := helper.Pagination(r)

	offset := (result.Page_number - 1) * result.Limit
	// Limitresult := result.Limit / 2

	// if result.Page_number > 1 && Limitresult == 0 {
	// 	offset += 1

	// }

	CategoryResponses := controller.CategoryService.FindAll(r.Context(), result.Limit, offset)

	WebResponse := web.WebResponse{
		Code:        200,
		Status:      "OK",
		Data:        CategoryResponses,
		Page_number: result.Page_number,
		Offset:      offset,
	}
	helper.WriteResponse(w, WebResponse)

}
