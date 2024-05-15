# wasm-http-client

**EXPERIMENTAL**

This is an HTTP client to be used in wasm binaries that target wazero as runtime until the WASI spec allows the features to do it.

In order to use it we need two things:

1. To consume a function in the guest

```go
package main

import httpclient "github.com/jcchavezs/wasm-http-client/guest/tinygo/client"

//export do_something
func do_something()
  statusCode := httpclient.HTTPDo(httpclient.Request{
    Method: "GET",
    URL:    "http//my-url",
  })
```

2. To load the module in the host

```go

import httpclient "github.com/jcchavezs/wasm-http-client"

// ...

// Create a new WebAssembly Runtime.
 r := wazero.NewRuntime(ctx)
 defer r.Close(ctx) // This closes everything this Runtime created.

 if err := httpclient.LoadModuleIntoRuntime(ctx, r); err != nil {
  t.Fatalf("failed to load http-client module: %v", err)
 }
```

**Important:** The current API only returns the status code, being 0 if there was an error that is also printed to STDOUT. The reason for this is because copying the entire response into the stack sounds overkill and I do not believe it should be taken as a general approach. Instead I would like to learn about use cases for this and then make a better design on the API.

Some ideas to approach the response are:

1. Allow status code and response headers.
2. Add a config so runtimes can decide what fields to copy into the stack (feasible if there is only one call per guest execution).
