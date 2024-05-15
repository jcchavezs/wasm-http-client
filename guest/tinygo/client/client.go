package client

import (
	"fmt"
	"strings"

	"github.com/jcchavezs/wasm-http-client/guest/tinygo/client/internal/imports"
)

type Request struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string][]string
}

// HTTPDo executes a HTTP request.
func HTTPDo(req Request) int {
	methodPtr, methodLen := stringToPtr(req.Method)
	urlPtr, urlLen := stringToPtr(req.URL)
	headersPtr, headersLen := stringToPtr(headersToString(req.Headers))
	bodyPtr, bodyLen := bytesToPtr(req.Body)

	return int(imports.HTTPDo(
		methodPtr, methodLen,
		urlPtr, urlLen,
		headersPtr, headersLen,
		bodyPtr, bodyLen,
	))
}

func headersToString(headers map[string][]string) string {
	// TODO(jcchavezs): Is this the best way to serialize the headers?
	hs := ""
	for k, vs := range headers {
		hs += fmt.Sprintf("%s=%s\n", k, strings.Join(vs, ";"))
	}
	return hs
}
