package disk

import (
	"context"

	"github.com/kehiy/blobstore"
)

type Disk struct {
	Path string
}

func New(path string) blobstore.Store {
	return Disk{
		Path: path,
	}
}

func (d Disk) Init(_ context.Context) error {
	return nil
}

func (d Disk) Close() error {
	return nil
}
