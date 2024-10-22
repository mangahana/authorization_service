package s3

import (
	"authorization_service/internal/core/configuration"
	"bytes"
	"context"
	"mime"
	"net/http"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type storage struct {
	client     *minio.Client
	bucketName string
}

func New(cfg *configuration.S3Config) (*storage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return &storage{client: client, bucketName: cfg.BucketName}, nil
}

func (s *storage) Put(c context.Context, object []byte) (string, error) {
	mimeType := http.DetectContentType(object)
	opts := minio.PutObjectOptions{
		ContentType: mimeType,
	}

	exts, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return "", err
	}

	randomKey, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	filename := randomKey.String() + exts[0]

	_, err = s.client.PutObject(c, s.bucketName, filename, bytes.NewReader(object), int64(len(object)), opts)
	return filename, err
}

func (s *storage) Remove(c context.Context, objectName string) error {
	return s.client.RemoveObject(c, s.bucketName, objectName, minio.RemoveObjectOptions{})
}
