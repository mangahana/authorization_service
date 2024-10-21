package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) Register(c echo.Context) error {
	var dto dto.Register

	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	token, err := h.useCase.Register(c.Request().Context(), &dto)
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, err.Error()))
	}

	return c.JSON(200, map[string]string{"token": token})
}
