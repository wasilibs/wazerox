package arm64

import (
	"testing"

	"github.com/wasilibs/wazerox/internal/engine/wazevo/backend/regalloc"
	"github.com/wasilibs/wazerox/internal/testing/require"
)

func Test_regs(t *testing.T) {
	require.Equal(t, x29, fpVReg.RealReg())
	require.Equal(t, regalloc.RegTypeInt, fpVReg.RegType())
	require.Equal(t, sp, spVReg.RealReg())
	require.Equal(t, regalloc.RegTypeInt, spVReg.RegType())
	require.Equal(t, xzr, xzrVReg.RealReg())
	require.Equal(t, regalloc.RegTypeInt, xzrVReg.RegType())
}
