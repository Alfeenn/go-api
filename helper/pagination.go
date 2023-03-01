package helper

import (
	"net/http"
	"strconv"

	"github.com/Alfeenn/api-go/model/web"
)

func Pagination(r *http.Request) web.WebResponse {
	query := r.URL.Query()
	pagenum := 1
	limit := 9
	for key, value := range query {
		queryvalue := value[len(value)-1]
		switch key {
		case "page":
			pagenum, _ = strconv.Atoi(queryvalue)

		case "limit":
			limit, _ = strconv.Atoi(queryvalue)

		}

	}
	return web.WebResponse{
		Page_number: pagenum,
		Limit:       limit,
	}
}
