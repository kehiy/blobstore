package disk

import (
	"context"
	"os"
	"path"
)

func (d Disk) Delete(ctx context.Context, sha256 string) error {
	return os.Remove(path.Join(d.Path, sha256))
}
