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

	UpdateUser(c context.Context, user *models.UserSession, dto *dto.Update) error
	UpdatePassword(c context.Context, user *models.UserSession, dto *dto.ChangePassword) error
	UpdatePhoto(c context.Context, user *models.UserSession, file []byte) (string, error)
}

type AMQP interface {
	SendUserUpdateEvent(data models.UpdateUserEvent) error
}

type useCase struct {
	repo       infrastructure.Repository
	sms        infrastructure.SMS
	s3         infrastructure.S3
	amqp       AMQP
	cdnBaseUrl string
}

func New(
	repository infrastructure.Repository,
	sms infrastructure.SMS,
	s3 infrastructure.S3,
	amqp AMQP,
	cdnBaseUrl string,
) *useCase {
	return &useCase{
		repo:       repository,
		sms:        sms,
		s3:         s3,
		amqp:       amqp,
		cdnBaseUrl: cdnBaseUrl,
	}
}
