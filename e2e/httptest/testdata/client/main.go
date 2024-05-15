//go:generate tinygo build -o ../client.wasm -scheduler=none --no-debug -target=wasi ./main.go
//go:generate wasm2wat ../client.wasm -o ../client.wat

package main

import (
	"log"
	"os"

	httpclient "github.com/jcchavezs/wasm-http-client/guest/tinygo/client"
)

func main() {}

//export call_me
func callMe() {
	statusCode := httpclient.HTTPDo(httpclient.Request{
		Method: "GET",
		URL:    os.Args[0],
	})

	if statusCode < 200 && statusCode > 299 {
		log.Fatalf("unexpected status code: %d", statusCode)
	}
}
