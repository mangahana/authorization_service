package application

import (
	"authorization_service/internal/core/dto"
	"authorization_service/internal/infrastructure"
	"context"
)

type UseCase interface {
	Join(c context.Context, ip string, dto *dto.Join) error
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
