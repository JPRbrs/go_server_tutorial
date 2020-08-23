package main

import (
	"fmt"

	"github.com/SlyMarbo/rss"
)

var example_feed = "https://lapastillaroja.net/feed/"

func GetFeed() *rss.Feed {
	fmt.Println("fetching:" + example_feed)
	feed, err := rss.Fetch(example_feed)
	if err != nil {
		// handle error.
		fmt.Println("error")
	}

	// ... Some time later ...

	// err = feed.Update()
	// if err != nil {
	// 	// handle error.
	// 	Println("error")
	// }

	return feed
}
