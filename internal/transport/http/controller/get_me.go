package controller

import "github.com/labstack/echo/v4"

func (h *controller) GetMe(c echo.Context) error {
	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, err)
	}

	return c.JSON(200, user)
}
