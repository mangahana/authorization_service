package dto

type Join struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
}
