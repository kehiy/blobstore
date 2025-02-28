package memory

import (
	"context"
	"errors"
)

func (m Memory) Store(_ context.Context, sha256 string, body []byte) error {
	m.Lock()
	defer m.Unlock()

	if m.LimitSize {
		if m.currentSize+len(body) > m.MaxSize {
			return errors.New("max size exceeded")
		}
	}

	m.currentSize += len(body)
	m.Blobs[sha256] = body

	return nil
}
