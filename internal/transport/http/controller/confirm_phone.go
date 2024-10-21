package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) ConfirmPhone(c echo.Context) error {
	var dto dto.ConfirmPhone
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	if err := h.useCase.ConfirmPhone(c.Request().Context(), &dto); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, "OK")
}
