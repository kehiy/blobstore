package policy

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/nbd-wtf/go-nostr"
)

type Policy struct {
	// MaxSize is the maximum size of a blob to be stored in bytes.
	MaxSize int

	// AllowedExts is the list of acceptable file extensions. Empty array means accept all.
	AllowedExts []string

	// Only allow the owner to check their blob lists.
	ListOwnerOnly bool

	// BannedBlobs is the list of sha256 hash of medias that no body can write or read.
	BannedBlobs []string

	// BannedPubkeys is the list of pubkeys that are not allowed to take any actions.
	BannedPubkeys []string

	// WriteAllowedPubkeys is the list of whitelisted pubkeys for writing methods. Would be checked iff len > 0.
	WriteAllowedPubkeys []string

	// ReadAllowedPubkeys is the list of whitelisted pubkeys for reading methods. Would be checked iff len > 0.
	ReadAllowedPubkeys []string

	// ReadAllowedBlobs is the list of whitelisted blob ids to be read. Would be checked iff len > 0.
	ReadAllowedBlobs []string
}

func (p *Policy) RejectUpload(_ context.Context, auth *nostr.Event, size int, ext string) (bool, string, int) {
	if size > p.MaxSize {
		return false,
			fmt.Sprintf("max accepted media size is: %d", p.MaxSize),
			http.StatusRequestEntityTooLarge
	}

	if !slices.Contains(p.AllowedExts, ext) {
		return false,
			fmt.Sprintf("%s extension is not supported, supported extensions: %v", ext, p.AllowedExts),
			http.StatusPreconditionFailed
	}

	if slices.Contains(p.BannedPubkeys, auth.PubKey) {
		return false,
			fmt.Sprintf("the %s pubkey is not allowed to upload blobs", auth.PubKey),
			http.StatusForbidden
	}

	if len(p.WriteAllowedPubkeys) > 0 {
		if !slices.Contains(p.WriteAllowedPubkeys, auth.PubKey) {
			return false,
				fmt.Sprintf("the %s pubkey is not allowed to upload blobs", auth.PubKey),
				http.StatusForbidden
		}
	}

	return true, "", http.StatusOK
}

func (p *Policy) RejectGet(_ context.Context, auth *nostr.Event, sha256 string) (bool, string, int) {
	if slices.Contains(p.BannedBlobs, sha256) {
		return false,
			fmt.Sprintf("the %s blob is not allowed to be read", sha256),
			http.StatusForbidden
	}

	if slices.Contains(p.BannedPubkeys, auth.PubKey) {
		return false,
			fmt.Sprintf("the %s pubkey is not allowed to get blobs", auth.PubKey),
			http.StatusForbidden
	}

	if len(p.ReadAllowedBlobs) > 0 {
		if !slices.Contains(p.ReadAllowedBlobs, sha256) {
			return false,
				fmt.Sprintf("the %s blob is not allowed to be read", auth.PubKey),
				http.StatusForbidden
		}
	}

	if len(p.ReadAllowedPubkeys) > 0 {
		if !slices.Contains(p.ReadAllowedPubkeys, auth.PubKey) {
			return false,
				fmt.Sprintf("the %s pubkey is not allowed to read blobs", auth.PubKey),
				http.StatusForbidden
		}
	}

	return true, "", http.StatusOK
}

func (p *Policy) RejectList(_ context.Context, auth *nostr.Event, pubkey string) (bool, string, int) {
	if p.ListOwnerOnly {
		if auth.PubKey != pubkey {
			return false,
				fmt.Sprintf("the %s pubkey is not allowed to get blobs list", auth.PubKey),
				http.StatusForbidden
		}
	}

	if len(p.ReadAllowedPubkeys) > 0 {
		if !slices.Contains(p.ReadAllowedPubkeys, auth.PubKey) {
			return false,
				fmt.Sprintf("the %s pubkey is not allowed to read blobs", auth.PubKey),
				http.StatusForbidden
		}
	}

	return true, "", http.StatusOK
}

func (p *Policy) RejectDelete(_ context.Context, auth *nostr.Event, sha256 string) (bool, string, int) {
	if slices.Contains(p.BannedPubkeys, auth.PubKey) {
		return false,
			fmt.Sprintf("the %s pubkey is not allowed to upload blobs", auth.PubKey),
			http.StatusForbidden
	}

	if len(p.WriteAllowedPubkeys) > 0 {
		if !slices.Contains(p.WriteAllowedPubkeys, auth.PubKey) {
			return false,
				fmt.Sprintf("the %s pubkey is not allowed to upload blobs", auth.PubKey),
				http.StatusForbidden
		}
	}

	return true, "", http.StatusOK
}
