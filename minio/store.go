package minio

import (
	"context"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (m *Minio) Store(ctx context.Context, sha256 string, body []byte) error {
	tempFileName := path.Join(m.TempPath, uuid.NewString()+sha256)

	if err := os.WriteFile(tempFileName, body, os.ModeTemporary); err != nil {
		return err
	}

	_, err := m.MinioClient.FPutObject(ctx, m.BucketName, sha256, tempFileName, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	if err := os.Remove(tempFileName); err != nil {
		return err
	}

	return nil
}
