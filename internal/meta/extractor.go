package meta

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/sergivb01/unfurl-worker/internal/utils"
)

var ErrorType = errors.New("should not be non-ptr or nil")

func ExtractInfoFromReader(reader io.ReadCloser) (*PageInfo, error) {
	defer utils.BenchmarkFunction(time.Now(), "ExtractInfoFromReader(io.ReadCloser)")

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating document from reader: %w", err)
	}
	defer reader.Close()

	// info, err := getPageData(doc)
	// if err != nil {
	// 	return nil, fmt.Errorf("error getting info: %w", err)
	// }

	return getPageData(doc), nil
}

func getAttribute(el *goquery.Selection) (string, bool) {
	if name, ok := el.Attr("name"); ok {
		return name, true
	}
	return el.Attr("property")
}

func getPageData(doc *goquery.Document) *PageInfo {
	defer utils.BenchmarkFunction(time.Now(), "getPageData(*goquery.Document)")

	info := &PageInfo{}

	selection := doc.Find("meta[content]")
	for i := range selection.Nodes {
		el := selection.Eq(i)
		name, ok := getAttribute(el)
		if !ok {
			continue
		}

		// no need to check if the attribute exists as we
		// have just query-selected meta tags with content tag
		val, _ := el.Attr("content")
		info.updateField(name, val)
	}
	return info
}

// func getPageData(doc *goquery.Document) *PageInfo {
// 	defer utils.BenchmarkFunction(time.Now(), "getPageData(*goquery.Document)")
//
// 	info := &PageInfo{}
//
// 	for tag, _ := range internalReflectionCache {
// 		var (
// 			content string
// 			ok      bool
// 			sel     *goquery.Selection
// 		)
//
// 		if sel = doc.Find(fmt.Sprintf("meta[property=\"%s\"]", tag)).First(); sel.Size() > 0 {
// 			content, ok = sel.Attr("content")
// 		}
//
// 		if !ok {
// 			if sel = doc.Find(fmt.Sprintf("meta[name=\"%s\"]", tag)).First(); sel.Size() > 0 {
// 				content, ok = sel.Attr("content")
// 			}
// 		}
//
// 		if !ok {
// 			continue
// 		}
//
// 		info.updateField(tag, content)
// 	}
//
// 	return info
// }
