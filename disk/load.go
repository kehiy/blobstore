package disk

import (
	"context"
	"io"
	"os"
	"path"
)

func (d Disk) Load(_ context.Context, sha256 string) (io.ReadSeeker, error) {
	return os.Open(path.Join(d.Path, sha256))
}
