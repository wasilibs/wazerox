package v2

import (
	"context"
	"runtime"
	"testing"

	wazero "github.com/wasilibs/wazerox"
	"github.com/wasilibs/wazerox/api"
	"github.com/wasilibs/wazerox/internal/engine/wazevo"
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
	c := wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures)
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	wazevo.ConfigureWazevo(c)
	spectest.Run(t, Testcases, context.Background(), c)
}
