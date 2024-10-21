package dto

type Join struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
}

type ConfirmPhone struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
	Code  string `json:"code" validate:"required,min=6,max=6"`
}
