package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{Transport: &http.Transport{
	DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
		dialer := &net.Dialer{
			Timeout: time.Duration(3) * time.Second,
		}
		return dialer.DialContext(ctx, network, addr)
	},
	TLSHandshakeTimeout: 3 * time.Second,
	MaxIdleConns:        150,
	MaxIdleConnsPerHost: 150,
	DisableKeepAlives:   true,
	ForceAttemptHTTP2:   true,
}}

func getReaderFromURL(ctx context.Context, urlAddress string, enableCompression bool) (io.ReadCloser, error) {
	defer BenchmarkFunction(time.Now(), "getReaderFromURL(ctx, "+urlAddress+")")
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}
	fmt.Printf("received %q, sending %q\n", urlAddress, parsedUrl.String())

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

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	reader, err := getReaderFromResponse(res)
	if err != nil {
		return nil, err
	}
	// defer reader.Close()

	fmt.Printf("ContentEncoding: %s, proto=%s, protoMajor=%d, statusCode=%d, compressed=%t\n", res.Header.Get("Content-Encoding"), res.Proto, res.ProtoMajor, res.StatusCode, !res.Uncompressed)

	return reader, nil
}
