package application

import (
	"authorization_service/internal/core/cerror"
	"authorization_service/internal/core/dto"
	"context"
	"crypto/rand"
	"log"
	"math/big"
)

func (u *useCase) Join(c context.Context, ip string, dto *dto.Join) error {
	exists, err := u.repo.IsPhoneExists(c, dto.Phone)
	if err != nil {
		return err
	}
	if exists {
		return cerror.New(cerror.PHONE_USED, "this phone is already in use")
	}

	if err := u.protect(c, ip, dto.Phone); err != nil {
		return err
	}

	isSent, err := u.repo.IsCodeSent(c, dto.Phone)
	if err != nil {
		log.Println(err)
		return err
	}
	if isSent {
		return cerror.New(cerror.CODE_SENT, "the code has been sent")
	}

	newCode, err := code(6)
	if err != nil {
		return cerror.New(cerror.SMS_CANNOT_SENT, "code cannot be sent")
	}

	err = u.sms.Send(c, "7"+dto.Phone, "mangahana\nРастау коды: "+newCode)
	if err != nil {
		return cerror.New(cerror.SMS_CANNOT_SENT, "code cannot be sent")
	}

	return u.repo.CreateConfirmationCode(c, dto.Phone, newCode, ip)
}

func code(length int) (string, error) {
	var output string

	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return "", err
		}
		output += nBig.String()
	}

	return output, nil
}
