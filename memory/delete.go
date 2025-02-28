package memory

import (
	"context"
)

func (m Memory) Delete(_ context.Context, sha256 string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.Blobs, sha256)

	return nil
}
