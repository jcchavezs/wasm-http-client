//go:build tinygo.wasm

package imports

//go:wasm-module wasm_http_client
//go:export http_do
func httpDo(
	methodPtr uintptr, methodLen uint32,
	urlPtr uintptr, urlLen uint32,
	headersPtr uintptr, headersLen uint32,
	bodyPtr uintptr, bodyLen uint32,
) uint32
