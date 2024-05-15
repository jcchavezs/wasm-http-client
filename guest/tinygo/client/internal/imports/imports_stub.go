//go:build !tinygo.wasm

package imports

// httpDo is stubbed for compilation outside TinyGo.
func httpDo(uintptr, uint32, uintptr, uint32, uintptr, uint32, uintptr, uint32) uint32 {
	return 0
}
