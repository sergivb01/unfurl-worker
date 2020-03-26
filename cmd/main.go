package main

import (
	"context"
	"fmt"
)

func main() {
	bench(true)
}

func bench(compress bool) {
	reader, err := getReaderFromURL(context.TODO(), "https://cloudflare.com/blog", compress)
	if err != nil {
		panic(err)
	}

	meta, err := extractFromReader(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%#+v\n", meta)
	// fmt.Printf("%d\n", len(b))
}
