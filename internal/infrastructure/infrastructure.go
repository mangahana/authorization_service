package infrastructure

import "context"

type Repository interface {
	CreateConfirmationCode(c context.Context, phone, code, ip string) error

	IsPhoneExists(c context.Context, phone string) (bool, error)
	IsCodeSent(c context.Context, phone string) (bool, error)
	IsPhoneBlocked(c context.Context, phone string) (bool, error)

	CodesCountLastHourByIP(c context.Context, ip string) (int, error)
	AddIPblock(c context.Context, ip string) error
	IsIpBlocked(c context.Context, ip string) (bool, error)

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
