package application

import (
	"authorization_service/internal/core/models"
	"context"
)

func (u *useCase) GetUserByID(c context.Context, id int) (models.User, error) {
	return u.repo.GetUserByID(c, id)
}
