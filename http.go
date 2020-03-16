package mediaprobe

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func getRemoteFile(uri *url.URL) (io.ReadCloser, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
