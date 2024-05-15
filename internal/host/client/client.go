package client

import (
	"context"
	"fmt"
	"strings"

	apiclient "github.com/jcchavezs/wasm-http-client/internal/host/api/client"
	"github.com/tetratelabs/wazero"
	wazeroapi "github.com/tetratelabs/wazero/api"
)

type module struct {
	host apiclient.Host
}

const (
	i32 = wazeroapi.ValueTypeI32
	i64 = wazeroapi.ValueTypeI64
)

func Module(r wazero.Runtime) (wazero.CompiledModule, error) {
	m := &module{host: host{}}

	return r.NewHostModuleBuilder(apiclient.ModuleName).
		NewFunctionBuilder().
		WithGoModuleFunction(wazeroapi.GoModuleFunc(m.Do), []wazeroapi.ValueType{i32, i32, i32, i32, i32, i32, i32, i32}, []wazeroapi.ValueType{i32}).
		WithParameterNames("method", "method_len", "url", "url_len", "headers", "headers_len", "body", "body_len").
		Export(apiclient.FuncHTTPDo).
		Compile(context.Background())
}

func (m *module) Do(ctx context.Context, mod wazeroapi.Module, params []uint64) {
	method := uint32(params[0])
	methodLen := uint32(params[1])
	url := uint32(params[2])
	urlLen := uint32(params[3])
	headers := uint32(params[4])
	headersLen := uint32(params[5])
	body := uint32(params[6])
	bodyLen := uint32(params[7])

	if methodLen == 0 {
		panic("method cannot be empty")
	}

	if urlLen == 0 {
		panic("url cannot be empty")
	}

	statusCode, err := m.host.HTTPDo(
		ctx,
		mustReadString(mod.Memory(), "method", method, methodLen),
		mustReadString(mod.Memory(), "url", url, urlLen),
		toHeaders(mustReadString(mod.Memory(), "headers", headers, headersLen)),
		mustRead(mod.Memory(), "body", body, bodyLen),
	)

	if err != nil {
		fmt.Println(err)
	}

	params[0] = uint64(statusCode)
}

func toHeaders(headers string) map[string][]string {
	if headers == "" {
		return nil
	}

	h := make(map[string][]string)
	for _, line := range strings.Split(headers, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		h[parts[0]] = strings.Split(parts[1], ";")
	}
	return h
}

// mustReadString is a convenience function that casts mustRead
func mustReadString(mem wazeroapi.Memory, fieldName string, offset, byteCount uint32) string {
	if byteCount == 0 {
		return ""
	}
	return string(mustRead(mem, fieldName, offset, byteCount))
}

var emptyBody = make([]byte, 0)

// mustRead is like api.Memory except that it panics if the offset and byteCount are out of range.
func mustRead(mem wazeroapi.Memory, fieldName string, offset, byteCount uint32) []byte {
	if byteCount == 0 {
		return emptyBody
	}
	buf, ok := mem.Read(offset, byteCount)
	if !ok {
		panic(fmt.Errorf("out of memory reading %s", fieldName))
	}
	return buf
}
