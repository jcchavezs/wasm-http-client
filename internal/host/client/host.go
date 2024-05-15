package client

import (
	"bytes"
	"context"
	"io"
	"net/http"

	apiclient "github.com/jcchavezs/wasm-http-client/internal/host/api/client"
)

type host struct{}

var _ apiclient.Host = host{}

func (host) HTTPDo(ctx context.Context, method, url string, headers map[string][]string, body []byte) (int, error) {
	var rBody io.Reader
	if len(body) > 0 {
		rBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, rBody)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)
	if len(headers) > 0 {
		req.Header = http.Header(headers)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
