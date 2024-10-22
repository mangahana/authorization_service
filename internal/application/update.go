package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"authorization_service/internal/core/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (u *useCase) UpdateUser(c context.Context, user *models.UserSession, dto *dto.Update) error {
	if err := checkUsername(dto.Username); err != nil {
		return err
	}

	userId, err := u.repo.GetUserIDbyUsername(c, dto.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		if user.ID != userId {
			return cerror.New(cerror.USERNAME_USED, "this username already in use")
		}
	}

	return u.repo.UpdateUser(c, user.ID, dto)
}
