package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) ChangePassword(c echo.Context) error {
	var dto dto.ChangePassword
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "Invalid data"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, err)
	}

	if err := h.useCase.UpdatePassword(c.Request().Context(), &user, &dto); err != nil {
		return c.JSON(400, err.Error())
	}

	return nil
}
