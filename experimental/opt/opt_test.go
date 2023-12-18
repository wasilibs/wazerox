package opt_test

import (
	"context"
	"runtime"
	"testing"

	wazero "github.com/wasilibs/wazerox"
	"github.com/wasilibs/wazerox/experimental/opt"
	"github.com/wasilibs/wazerox/internal/testing/require"
)

func TestUseOptimizingCompiler(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		return
	}
	c := opt.NewRuntimeConfigOptimizingCompiler()
	r := wazero.NewRuntimeWithConfig(context.Background(), c)
	require.NoError(t, r.Close(context.Background()))
}
