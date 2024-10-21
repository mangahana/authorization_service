package application

import (
	"authorization_service/internal/core/dto"
	"authorization_service/internal/core/models"
	"authorization_service/internal/infrastructure"
	"context"
)

type UseCase interface {
	Join(c context.Context, ip string, dto *dto.Join) error
	ConfirmPhone(c context.Context, dto *dto.ConfirmPhone) error
	Register(c context.Context, dto *dto.Register) (string, error)
	Login(c context.Context, dto *dto.Login) (string, error)
	IsPhoneExists(c context.Context, phone string) (bool, error)

	GetUserByID(c context.Context, id int) (models.User, error)
	GetMe(c context.Context, token string) (models.UserSession, error)
}

type useCase struct {
	repo infrastructure.Repository
	sms  infrastructure.SMS
	s3   infrastructure.S3
}

func New(
	repository infrastructure.Repository,
	sms infrastructure.SMS,
	s3 infrastructure.S3,
) *useCase {
	return &useCase{
		repo: repository,
		sms:  sms,
		s3:   s3,
	}
}
