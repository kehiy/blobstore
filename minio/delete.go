package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (m *Minio) Delete(ctx context.Context, sha256 string) error {
	return m.MinioClient.RemoveObject(ctx, m.BucketName, sha256, minio.RemoveObjectOptions{
		GovernanceBypass: true,
	})
}
