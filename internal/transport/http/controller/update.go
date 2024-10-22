package controller

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"log"

	"github.com/labstack/echo/v4"
)

func (h *controller) Update(c echo.Context) error {
	var dto dto.Update
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "Invalid Data"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, err)
	}

	if err := h.useCase.UpdateUser(c.Request().Context(), &user, &dto); err != nil {
		log.Println(err)
		return c.JSON(400, err)
	}

	return c.String(200, "OK")
}
