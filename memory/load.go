package memory

import (
	"bytes"
	"context"
	"errors"
	"io"
)

func (m Memory) Load(_ context.Context, sha256 string) (io.ReadSeeker, error) {
	m.Lock()
	defer m.Unlock()

	blob, ok := m.Blobs[sha256]
	if !ok {
		return nil, errors.New("not found")
	}

	return bytes.NewReader(blob), nil
}
