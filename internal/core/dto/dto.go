package dto

type Join struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
}

type ConfirmPhone struct {
	Phone string `json:"phone" validate:"required,min=10,max=10"`
	Code  string `json:"code" validate:"required,min=6,max=6"`
}

type Register struct {
	Phone            string `json:"phone" validate:"required,min=10,max=10"`
	ConfirmationCode string `json:"confirmation_code" validate:"required,min=6,max=6"`
	Username         string `json:"username" validate:"required"`
	Password         string `json:"password" validate:"required,min=8"`
}

type Login struct {
	Phone    string `json:"phone" validate:"required,min=10,max=10"`
	Password string `json:"password" validate:"required"`
}

type Update struct {
	Username    string `json:"username" validate:"required,min=3"`
	Description string `json:"description"`
}
