package controller

import (
	"authorization_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) IsPhoneExists(c echo.Context) error {
	phone := c.QueryParam("phone")
	if err := h.validator.Var(phone, "required,min=10,max=10"); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	exists, err := h.useCase.IsPhoneExists(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(400, err)
	}

	if exists {
		return c.JSON(200, true)
	} else {
		return c.JSON(200, false)
	}
}
