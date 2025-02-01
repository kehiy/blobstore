package disk

import (
	"bytes"
	"context"
	"io"
	"os"
	"path"
)

func (d Disk) Store(_ context.Context, sha256 string, body []byte) error {
	file, err := os.Create(path.Join(d.Path, sha256))
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, bytes.NewReader(body)); err != nil {
		return err
	}

	return nil
}
