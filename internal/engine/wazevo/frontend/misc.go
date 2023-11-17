package frontend

import (
	"github.com/wasilibs/wazerox/internal/engine/wazevo/ssa"
	"github.com/wasilibs/wazerox/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
