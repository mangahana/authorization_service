package controller

import (
	"authorization_service/internal/core/cerror"
	"errors"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func (h *controller) GetByID(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Var(userId, "required,gt=0"); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	user, err := h.useCase.GetUserByID(c.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(404, cerror.New(cerror.NOT_FOUND, "user not found"))
		}
		log.Println(err)
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, "INTERNAL_SERVER_ERROR"))
	}

	return c.JSON(200, user)
}
