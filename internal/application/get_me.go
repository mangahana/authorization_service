package application

import (
	"authorization_service/internal/core/models"
	"context"
)

func (u *useCase) GetMe(c context.Context, token string) (models.UserSession, error) {
	return u.repo.GetByToken(c, token)
}
