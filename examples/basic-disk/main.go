package main

import (
	"fmt"
	"net/http"

	"github.com/fiatjaf/eventstore/badger"
	"github.com/fiatjaf/khatru/blossom"
	"github.com/kehiy/blobstore/disk"
	"github.com/fiatjaf/khatru"
)

func main() {
	relay := khatru.NewRelay()

	db := &badger.BadgerBackend{Path: "/tmp/khatru-badgern-tmp"}
	if err := db.Init(); err != nil {
		panic(err)
	}

	
	bl := blossom.New(relay, "http://localhost:3334")
	
	d := disk.New("/blossom")
	bl.LoadBlob = append(bl.LoadBlob, d.Load)
	bl.DeleteBlob = append(bl.DeleteBlob, d.Delete)
	bl.StoreBlob = append(bl.StoreBlob, d.Store)

	bl.Store = blossom.EventStoreBlobIndexWrapper{Store: db, ServiceURL: bl.ServiceURL}

	fmt.Println("running on :3334")
	http.ListenAndServe(":3334", relay)
}
