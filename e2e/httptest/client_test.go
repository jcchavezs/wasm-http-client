//go:build e2e

package e2e

import (
	"context"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	httpclient "github.com/jcchavezs/wasm-http-client"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/client.wasm
var clientGuest string

func TestE2E(t *testing.T) {
	call := new(int)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*call++
	}))
	defer srv.Close()

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	if err := httpclient.LoadModuleIntoRuntime(ctx, r); err != nil {
		t.Fatalf("failed to load http-client module: %v", err)
	}

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	cfg := wazero.NewModuleConfig().
		WithStdout(os.Stdout).WithArgs(srv.URL)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.InstantiateWithConfig(ctx, []byte(clientGuest), cfg)
	if err != nil {
		t.Fatalf("failed to instantiate module: %v", err)
	}

	require.Equal(t, 0, *call)

	_, err = mod.ExportedFunction("call_me").Call(context.Background())
	if err != nil {
		t.Fatalf("failed to call the function: %v", err)
	}

	require.Equal(t, 1, *call)
}
