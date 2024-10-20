package controller

import (
	"authorization_service/internal/application"

	"github.com/go-playground/validator/v10"
)

type controller struct {
	validator *validator.Validate
	useCase   application.UseCase
}

func New(useCase application.UseCase) *controller {
	return &controller{
		validator: validator.New(),
		useCase:   useCase,
	}
}
