package v1

import (
	"context"
	"runtime"
	"testing"

	wazero "github.com/wasilibs/wazerox"
	"github.com/wasilibs/wazerox/api"
	"github.com/wasilibs/wazerox/experimental/opt"
	"github.com/wasilibs/wazerox/internal/integration_test/spectest"
	"github.com/wasilibs/wazerox/internal/platform"
)

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestWazevo(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	c := opt.NewRuntimeConfigOptimizingCompiler().WithCoreFeatures(api.CoreFeaturesV1)
	spectest.Run(t, Testcases, context.Background(), c)
}
