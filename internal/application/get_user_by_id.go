package application

import (
	"authorization_service/internal/core/models"
	"context"
)

func (u *useCase) GetUserByID(c context.Context, id int) (models.User, error) {
	output, err := u.repo.GetUserByID(c, id)
	if err != nil {
		return output, err
	}

	output.Photo = u.cdnBaseUrl + output.Photo

	return output, nil
}
