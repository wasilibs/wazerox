package v2

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

const enabledFeatures = api.CoreFeaturesV2

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}

func TestWazevo(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	c := opt.NewRuntimeConfigOptimizingCompiler().WithCoreFeatures(enabledFeatures)
	spectest.Run(t, Testcases, context.Background(), c)
}
