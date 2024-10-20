package application

import (
	"authorization_service/internal/infrastructure"
)

type UseCase interface {
}

type useCase struct {
	repository infrastructure.Repository
	sms        infrastructure.SMS
	s3         infrastructure.S3
}

func New(
	repository infrastructure.Repository,
	sms infrastructure.SMS,
	s3 infrastructure.S3,
) *useCase {
	return &useCase{
		repository: repository,
		sms:        sms,
		s3:         s3,
	}
}
