package minio

import (
	"bytes"
	"context"
	"io"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (m *Minio) Load(ctx context.Context, sha256 string) (io.ReadSeeker, error) {
	tempFilePath := path.Join(m.TempPath, uuid.NewString()+sha256)

	err := m.MinioClient.FGetObject(ctx, m.BucketName, sha256, tempFilePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	f, err := os.Open(tempFilePath)
	if err != nil {
		os.Remove(tempFilePath)

		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		f.Close()
		os.Remove(tempFilePath)

		return nil, err
	}

	if err := f.Close(); err != nil {
		os.Remove(tempFilePath)

		return nil, err
	}

	if err := os.Remove(tempFilePath); err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
