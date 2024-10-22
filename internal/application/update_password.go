package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"authorization_service/internal/core/models"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (u *useCase) UpdatePassword(c context.Context, user *models.UserSession, dto *dto.ChangePassword) error {
	password, err := u.repo.GetPasswordByID(c, user.ID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dto.OldPassword))
	if err != nil {
		return cerror.New(cerror.INVALID_PASSWORD, "invalid password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), 10)
	if err != nil {
		return err
	}

	return u.repo.UpdatePassword(c, user.ID, string(hashedPassword))
}
