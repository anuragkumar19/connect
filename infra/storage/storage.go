package storage

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
)

type Storage struct {
	*minio.Client
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (Storage, error) {
	if err := config.Validate(); err != nil {
		return Storage{}, errors.Join(errors.New("invalid storage config"), err)
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.Secret, ""),
		Secure: config.UseSSL,
		Region: config.Region,
	})
	if err != nil {
		return Storage{}, errors.Join(errors.New("failed to create storage client"), err)
	}

	return Storage{
		Client: client,
	}, nil
}
