package controller

import (
	"authorization_service/internal/core/cerror"
	"io"

	"github.com/labstack/echo/v4"
)

func (h *controller) UpdatePhoto(c echo.Context) error {
	file, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, err)
	}

	filename, err := h.useCase.UpdatePhoto(c.Request().Context(), &user, file)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.String(200, filename)
}
