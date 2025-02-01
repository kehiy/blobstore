package minio

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (m *Minio) Load(ctx context.Context, sha256 string) (io.ReadSeeker, error) {
	// todo::: cleanup temp file.
	tempFilePath := path.Join(m.TempPath, uuid.NewString()+sha256)

	err := m.MinioClient.FGetObject(ctx, m.BucketName, sha256, tempFilePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return os.Open(tempFilePath)
}
