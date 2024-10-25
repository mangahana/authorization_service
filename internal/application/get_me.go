package application

import (
	"authorization_service/internal/core/models"
	"context"
)

func (u *useCase) GetMe(c context.Context, token string) (models.UserSession, error) {
	output, err := u.repo.GetByToken(c, token)
	if err != nil {
		return output, err
	}

	output.Photo = u.cdnBaseUrl + output.Photo

	return output, nil
}
