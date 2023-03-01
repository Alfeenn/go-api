package exception

import (
	"net/http"

	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/model/web"
	"github.com/go-playground/validator"
)

func ErrHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if ErrNotFound(w, r, err) {
		return
	}
	if ValidationErr(w, r, err) {
		return
	}
	InternalServer(w, r, err)
}

func ValidationErr(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteResponse(w, response)
		return true
	} else {
		return false
	}
}
func ErrNotFound(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	excption, ok := err.(NotFound)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   excption.Err,
		}

		helper.WriteResponse(w, response)
		return true
	} else {
		return false
	}
}

func InternalServer(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL STATUS ERROR",
		Data:   err,
	}

	helper.WriteResponse(w, response)
}
