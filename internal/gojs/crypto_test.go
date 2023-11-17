package gojs_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/wasilibs/wazerox/experimental"
	"github.com/wasilibs/wazerox/experimental/logging"
	"github.com/wasilibs/wazerox/internal/testing/require"
)

func Test_crypto(t *testing.T) {
	t.Parallel()

	var log bytes.Buffer
	loggingCtx := context.WithValue(testCtx, experimental.FunctionListenerFactoryKey{},
		logging.NewHostLoggingListenerFactory(&log, logging.LogScopeRandom))

	stdout, stderr, err := compileAndRun(loggingCtx, "crypto", defaultConfig)

	require.Zero(t, stderr)
	require.NoError(t, err)
	require.Equal(t, `7a0c9f9f0d
`, stdout)
	require.Equal(t, `==> go.runtime.getRandomData(r_len=32)
<==
==> go.runtime.getRandomData(r_len=8)
<==
==> go.syscall/js.valueCall(crypto.getRandomValues(r_len=5))
<== (n=5)
`, logString(log))
}
