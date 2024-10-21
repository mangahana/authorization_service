package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"context"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func (u *useCase) Register(c context.Context, input *dto.Register) (string, error) {
	confirmDto := dto.ConfirmPhone{
		Phone: input.Phone,
		Code:  input.ConfirmationCode,
	}
	if err := u.repo.GetSMSByCredentials(c, &confirmDto); err != nil {
		return "", err
	}

	regex, err := regexp.Compile("^[a-zA-Z0-9]+(_?[a-zA-Z0-9]+)*$")
	if err != nil {
		return "", err
	}

	if !regex.Match([]byte(input.Username)) {
		return "", cerror.New(cerror.USERNAME_INVALID, "username is not valid")
	}

	usernameExists, err := u.repo.IsUsernameExists(c, input.Username)
	if err != nil {
		return "", err
	}

	if usernameExists {
		return "", cerror.New(cerror.USERNAME_USED, "username is already in use")
	}

	if len(input.Password) < 8 {
		return "", cerror.New(cerror.PASSWORD_TOO_SHORT, "the password is too short")
	}

	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return "", err
	}
	input.Password = string(hashedPasword)

	userId, err := u.repo.Create(c, input)
	if err != nil {
		return "", err
	}

	u.repo.DeleteCodesByPhone(c, input.Phone)

	return u.repo.CreateSession(c, userId)
}