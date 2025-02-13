# Blob store

This package provides a list of blob storage interfaces compatible with the [khatru](https://github.com/fiatjaf/khatru) media server.

## Interface

Here is the khatru compatible interface implemented by blob store:

```go
// Store is an interface which allows blob to create, read, and delete.
type Store interface {
	// Init creates storage requirements and starts it.
	Init(ctx context.Context) error

	// Close closes the open connections, contexts and more.
	Close() error

	// Store saves the provided blob.
	Store(ctx context.Context, sha256 string, body []byte) error

	// Load reads the blob with the ID of sha256 provided from storage.
	Load(ctx context.Context, sha256 string) (io.ReadSeeker, error)
	
	// Delete removes the blob with the ID of sha256 provided from storage.
	Delete(ctx context.Context, sha256 string) error
}
```

### Storages

Current available storage interfaces are:

* [Disk](./disk/)
* [S3(minio)](./minio/)

## Policy

The khatru supports a policy interface which helps to reject requests to specific blossom routes. The [policy](./policy/) directory contains a simple module that helps you to define basic policy for your blob store dynamically.

## Roadmap

- [ ] IPFS interface.
- [ ] Khatru compatible blob index interface for different databases.
- [ ] More examples.
- [ ] Torrent interface.

## Contribution

Any kind of contribution is very welcomed.

## License

This library is published under [MIT License](./LICENSE).
