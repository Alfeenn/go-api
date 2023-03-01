package web

type WebResponse struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data"`
	Page_number int         `json:"page_number"`
	Offset      int         `json:"offset"`
	Limit       int
}
