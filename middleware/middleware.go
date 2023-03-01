package middleware

import (
	"net/http"

	"github.com/Alfeenn/api-go/helper"
	"github.com/Alfeenn/api-go/model/web"
)

type Middleware struct {
	Handler http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}

func (middleware *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		//error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		response := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteResponse(w, response)
	}
}
