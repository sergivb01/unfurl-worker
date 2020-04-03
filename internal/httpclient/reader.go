package httpclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dsnet/compress/brotli"
	"github.com/dsnet/compress/bzip2"
	"github.com/dsnet/compress/flate"
	"github.com/klauspost/compress/gzip"
)

// getReaderFromResponse returns the proper io.ReadCloser type according to the
// http response Content Encoding type
func getReaderFromResponse(res *http.Response) (io.ReadCloser, error) {
	var (
		err    error
		reader io.ReadCloser
	)

	switch res.Header.Get("Content-Encoding") {
	case "br":
		// TODO: search for a faster implementation / remove
		reader, err = brotli.NewReader(res.Body, nil)
		break
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		break
	case "bzip2":
		reader, err = bzip2.NewReader(res.Body, nil)
		break
	case "deflate":
		reader, err = flate.NewReader(res.Body, nil)
		break
	default:
		reader = res.Body
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create %q reader: %w", res.Header.Get("Content-Encoding"), err)
	}

	return reader, err
}
