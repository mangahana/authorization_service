package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) Join(c echo.Context) error {
	var dto dto.Join
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	err := h.useCase.Join(c.Request().Context(), c.RealIP(), &dto)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "OK")
}
