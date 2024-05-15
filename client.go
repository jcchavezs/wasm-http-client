package httpclient

import (
	"context"

	"github.com/jcchavezs/wasm-http-client/internal/host/client"
	"github.com/tetratelabs/wazero"
)

func LoadModuleIntoRuntime(ctx context.Context, r wazero.Runtime) error {
	cm, err := client.Module(r)
	if err != nil {
		return err
	}

	_, err = r.InstantiateModule(ctx, cm, wazero.NewModuleConfig())
	return err
}
