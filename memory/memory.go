package memory

import (
	"context"
	"sync"

	"github.com/kehiy/blobstore"
)

// Memory is a storage that keeps all blobs in memory.
type Memory struct {
	// Maximum allocated and used size of memory.
	MaxSize int

	// if set to false, you can set MaxSize to 0.
	LimitSize bool

	// size of currently stored blobs.
	currentSize int

	// Stored blobs.
	Blobs map[string][]byte

	*sync.RWMutex
}

func New(ms int, ls bool) blobstore.Store {
	return Memory{
		MaxSize:   ms,
		LimitSize: ls,
		Blobs:     make(map[string][]byte, ms),
	}
}

func (m Memory) Init(_ context.Context) error {
	return nil
}

func (m Memory) Close() error {
	for k := range m.Blobs {
		delete(m.Blobs, k)
	}

	return nil
}
