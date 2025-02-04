package blobstore

import (
	"context"
	"io"
)

// Store is an interface which allows blob create, read, delete.
type Store interface {
	// Init creates storage requirements and starts it.
	Init(ctx context.Context) error

	// Close closes the open connections, contexts and more.
	Close() error

	// Store saves the provided blob.
	Store(ctx context.Context, sha256 string, body []byte) error

	// Load reads the blob with ID of sha256 provided from storage.
	Load(ctx context.Context, sha256 string) (io.ReadSeeker, error)

	// Delete removes the blob with ID of sha256 provided from storage.
	Delete(ctx context.Context, sha256 string) error
}
