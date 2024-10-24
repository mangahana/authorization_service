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

	GetUserByID(c context.Context, id int) (models.User, error)
	GetCredentialsByPhone(c context.Context, phone string) (models.LoginCredentials, error)
	GetPasswordByID(c context.Context, id int) (string, error)
	GetByToken(c context.Context, token string) (models.UserSession, error)
	GetUserIDbyUsername(c context.Context, username string) (int, error)

	IsUsernameExists(c context.Context, username string) (bool, error)
	IsPhoneExists(c context.Context, phone string) (bool, error)
	IsCodeSent(c context.Context, phone string) (bool, error)
	IsPhoneBlocked(c context.Context, phone string) (bool, error)

	GetSMSByCredentials(c context.Context, dto *dto.ConfirmPhone) error
	CodesCountLastHourByIP(c context.Context, ip string) (int, error)
	AddIPblock(c context.Context, ip string) error
	IsIpBlocked(c context.Context, ip string) (bool, error)

	UpdateUser(c context.Context, userId int, dto *dto.Update) error
	UpdatePassword(c context.Context, userId int, password string) error
	UpdatePhoto(c context.Context, userId int, filename string) error

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
