package http

import (
	"authorization_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(401, cerror.New(cerror.UNAUTHORIZED, "need an access token"))
		}

		user, err := h.useCase.GetMe(c.Request().Context(), token)
		if err != nil {
			return c.JSON(401, cerror.New(cerror.UNAUTHORIZED, "invalid access token"))
		}

		c.Set("user", user)

		return next(c)
	}
}
