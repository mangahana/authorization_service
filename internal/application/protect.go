package application

import (
	"authorization_service/internal/core/cerror"
	"context"
)

func (u *useCase) protect(c context.Context, ip, phone string) error {
	isIpBlocked, err := u.repo.IsIpBlocked(c, ip)
	if err != nil {
		return err
	}
	if isIpBlocked {
		return cerror.New(cerror.BAD_REQUEST, "unknow error")
	}

	isBlocked, err := u.repo.IsPhoneBlocked(c, phone)
	if err != nil {
		return err
	}
	if isBlocked {
		return cerror.New(cerror.PHONE_BLOCKED, "this phone can't be use")
	}

	count, err := u.repo.CodesCountLastHourByIP(c, ip)
	if err != nil {
		return err
	}

	if count >= 5 {
		if err := u.repo.AddIPblock(c, ip); err != nil {
			return err
		}
		return cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST")
	}

	return nil
}
