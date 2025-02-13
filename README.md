# Blob store

This package provides a list of blob storage interfaces which is compatible with [khatru](https://github.com/fiatjaf/khatru) media server.

## Interface

Here is the khatru compatible interface implemented by blobstore:

```go
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
```

### Storages

Current availabe storage infterfaces are:

* [Disk](./disk/)
* [S3(minio)](./minio/)
* [IPFS](./ipfs/)

## Policy

The khatru supports a a policy interface which helps to reject request to specific blossom routes. The [policy](./policy/) directory contains a simple module that helps you to define basic policy for your blobstore dynamically.

## Roadmap

- [ ] Torrent interface.
- [ ] Khatru compatible blob index interface for different databases.
- [ ] More examples.

## Contribution

Any kind of contribution is very welcomed.

## License

This library is published under [MIT License](./LICENSE).
