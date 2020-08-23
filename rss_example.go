package rss_example

import (
	"fmt"

	"github.com/SlyMarbo/rss"
)

var example_feed = "https://lapastillaroja.net/feed/"

func GetFeed() Feed {
	fmt.Println("fetching :" + example_feed)
	feed, err := rss.Fetch("http://example.com/rss")
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
