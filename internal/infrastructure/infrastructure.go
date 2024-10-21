package infrastructure

import (
	"authorization_service/internal/core/dto"
	"authorization_service/internal/core/models"
	"context"
)

type Repository interface {
	Create(c context.Context, dto *dto.Register) (int, error)
	CreateConfirmationCode(c context.Context, phone, code, ip string) error
	CreateSession(c context.Context, userId int) (string, error)

	GetCredentialsByPhone(c context.Context, phone string) (models.LoginCredentials, error)

	IsUsernameExists(c context.Context, username string) (bool, error)
	IsPhoneExists(c context.Context, phone string) (bool, error)
	IsCodeSent(c context.Context, phone string) (bool, error)
	IsPhoneBlocked(c context.Context, phone string) (bool, error)

	GetSMSByCredentials(c context.Context, dto *dto.ConfirmPhone) error
	CodesCountLastHourByIP(c context.Context, ip string) (int, error)
	AddIPblock(c context.Context, ip string) error
	IsIpBlocked(c context.Context, ip string) (bool, error)

	DeleteCodesByPhone(c context.Context, phone string) error

	// close database connecion
	Close()
}

type SMS interface {
	Send(c context.Context, number, message string) error
}

type S3 interface {
	Put(c context.Context, object []byte) (string, error)
	Remove(c context.Context, objectName string) error
}
