package httpclient

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

const compressAlgo string = "gzip, br, bzip2, deflate"

// TODO: move to fasthttp, I think I'm memory leaking as fuck
func getClient() *http.Client {
	return &http.Client{Transport: &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
			dialer := &net.Dialer{
				Timeout: time.Millisecond * 1500, // 1.5s
			}
			return dialer.DialContext(ctx, network, addr)
		},
		TLSHandshakeTimeout: 3 * time.Second,
		MaxIdleConns:        -1,
		MaxIdleConnsPerHost: -1,
		DisableKeepAlives:   true,
		ForceAttemptHTTP2:   true,
	}}
}

// GetReaderFromURL returns the corresponding io.ReadCloser from a website
// according to the response type for compression
func GetReaderFromURL(ctx context.Context, urlAddress string, enableCompression bool) (io.ReadCloser, error) {
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", parsedUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// we will cache responses in the backend. If a request is being made is because the cache
	// got invalidated or expired, therefore we *must* get an uncached response
	req.Header.Set("Cache-Control", "no-cache")
	if enableCompression {
		req.Header.Set("Accept-Encoding", compressAlgo)
	}

	res, err := getClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	// reader will be closed in ExtractInfoFromReader
	reader, err := getReaderFromResponse(res)
	if err != nil {
		return nil, err
	}

	return reader, nil
}
