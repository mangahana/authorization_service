package infrastructure

import "context"

type Repository interface {
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
