package web

type RequestService struct {
	Name string `validate:"required" json:"name"`
}
