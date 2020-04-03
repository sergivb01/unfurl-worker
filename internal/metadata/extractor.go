package metadata

import (
	"errors"
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"

	"github.com/sergivb01/unfurl-worker/internal/utils"
)

var ErrorType = errors.New("should not be non-ptr or nil")

// ExtractInfoFromReader returns the page information from a response reader
func ExtractInfoFromReader(reader io.ReadCloser) (*PageInfo, error) {
	defer utils.Track(utils.BenchFunc("ExtractInfoFromReader(io.ReadCloser)"))
	// TODO: handle reader close error
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating document from reader: %w", err)
	}

	// TODO: maybe some parallelism??
	info := &PageInfo{}
	selection := doc.Find("meta[content]")
	for i := range selection.Nodes {
		el := selection.Eq(i)
		name, ok := getAttribute(el)
		if !ok {
			continue
		}

		// no need to check if the attribute exists as we
		// have just query-selected metadata tags with content tag
		val, _ := el.Attr("content")
		info.updateField(name, val)
	}

	return info, nil
}

func getAttribute(el *goquery.Selection) (string, bool) {
	if name, ok := el.Attr("name"); ok {
		return name, true
	}
	return el.Attr("property")
}
