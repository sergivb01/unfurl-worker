package old

import (
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/sergivb01/unfurl-worker/cmd"
)

const compressAlgo string = "gzip, br, bzip2, deflate"

func getAttribute(el *goquery.Selection) (string, bool) {
	if name, ok := el.Attr("name"); ok {
		return name, true
	}
	return el.Attr("property")
}

func extractFromReader(reader io.ReadCloser) (*Metatags, error) {
	defer main.BenchmarkFunction(time.Now(), "extractFromReader(io.ReadCloser)")
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating document: %w", err)
	}

	metatags := &Metatags{
		Title: doc.Find("title").Text(),
	}

	selection := doc.Find("meta[content]")
	for i := range selection.Nodes {
		el := selection.Eq(i)
		name, ok := getAttribute(el)
		if !ok {
			continue
		}

		val, _ := el.Attr("content")
		metatags.updateField(name, val)
	}

	return metatags, nil
}
