package minio

import (
	"context"

	"github.com/kehiy/blobstore"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Location        string
	TempPath        string
	UseSSL          bool
	MinioClient     *minio.Client
}

func New(endpoint string,
	accessKeyID string,
	secretAccessKey string,
	useSSL bool,
	bucketName string,
	location string,
) blobstore.Store {
	return &Minio{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
		BucketName:      bucketName,
		Location:        location,
	}
}

func (m *Minio) Init(ctx context.Context) error {
	minioClient, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyID, m.SecretAccessKey, ""),
		Secure: m.UseSSL,
	})
	if err != nil {
		return err
	}

	m.MinioClient = minioClient

	return nil
}

func (m *Minio) Close() error {
	return nil
}
