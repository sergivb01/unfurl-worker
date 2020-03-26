package main

import (
	"context"
	"fmt"

	"github.com/sergivb01/unfurl-worker/internal/meta"
	"github.com/sergivb01/unfurl-worker/internal/metaclient"
)

func main() {
	bench(true)
	// time.Sleep(time.Second)
	// bench2(true)
}

func bench(compress bool) {
	reader, err := metaclient.GetReaderFromURL(context.TODO(), "https://sergivos.dev", compress)
	if err != nil {
		panic(err)
	}

	info, _ := meta.ExtractInfoFromReader(reader)
	fmt.Printf("%+v\n", info)
}
