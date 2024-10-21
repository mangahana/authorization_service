package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (u *useCase) Login(c context.Context, dto *dto.Login) (string, error) {
	credentials, err := u.repo.GetCredentialsByPhone(c, dto.Phone)
	if err != nil {
		return "", err
	}

	hashedPassword := credentials.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(dto.Password))
	if err != nil {
		return "", cerror.New(cerror.INVALID_PASSWORD, "invalid password")
	}

	return u.repo.CreateSession(c, credentials.UserID)
}
