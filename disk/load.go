package disk

import (
	"context"
	"errors"
	"io"
	"os"
	"path"
)

func (d Disk) Load(_ context.Context, sha256 string) (io.ReadSeeker, error) {
	f, err := os.Open(path.Join(d.Path, sha256))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	return f, nil
}
