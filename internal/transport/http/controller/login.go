package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) Login(c echo.Context) error {
	var dto dto.Login

	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	token, err := h.useCase.Login(c.Request().Context(), &dto)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, map[string]string{"token": token})
}
