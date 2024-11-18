package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

type Storage struct {
	*minio.Client
}

func New(ctx context.Context, config *Config) (Storage, error) {
	if err := config.Validate(); err != nil {
		return Storage{}, fmt.Errorf("invalid storage config: %w", err)
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.Secret, ""),
		Secure: config.UseSSL,
		Region: config.Region,
	})
	if err != nil {
		return Storage{}, fmt.Errorf("failed to create storage client: %w", err)
	}

	if _, err := client.ListBuckets(ctx); err != nil {
		return Storage{}, fmt.Errorf("failed to list buckets from storage: %w", err)
	}

	log.Info().Msg("successfully authenticated with storage server")

	return Storage{
		Client: client,
	}, nil
}
