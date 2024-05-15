package client

import "context"

// Host is the interface that the WebAssembly module expects to be
// implemented by the host.
type Host interface {
	HTTPDo(ctx context.Context, method, url string, headers map[string][]string, body []byte) (int, error)
}
