package application

import (
	"authorization_service/internal/core/dto"
	"context"
)

func (u *useCase) ConfirmPhone(c context.Context, dto *dto.ConfirmPhone) error {
	return u.repo.GetSMSByCredentials(c, dto)
}
