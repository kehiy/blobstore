package blobstore

import (
	"context"
	"io"
)

type Store interface {
	Init(ctx context.Context) error
	Close() error
	Store(ctx context.Context, sha256 string, body []byte) error
	Load(ctx context.Context, sha256 string) (io.ReadSeeker, error)
	Delete(ctx context.Context, sha256 string) error
}
