package application

import "context"

func (u *useCase) IsPhoneExists(c context.Context, phone string) (bool, error) {
	return u.repo.IsPhoneExists(c, phone)
}
