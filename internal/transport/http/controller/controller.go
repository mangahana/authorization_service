package controller

import (
	"authorization_service/internal/application"
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func (h *controller) getUser(c echo.Context) (models.UserSession, error) {
	user, ok := c.Get("user").(models.UserSession)
	if !ok {
		return models.UserSession{}, cerror.New(cerror.UNAUTHORIZED, "getUser")
	}

	return user, nil
}
