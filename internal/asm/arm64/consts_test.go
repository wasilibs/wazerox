package arm64

import (
	"testing"

	"github.com/wasilibs/wazerox/internal/asm"
	"github.com/wasilibs/wazerox/internal/testing/require"
)

// TestInstructionName ensures that all the instruction's name is defined.
func TestInstructionName(t *testing.T) {
	for inst := asm.Instruction(0); inst < instructionEnd; inst++ {
		require.NotEqual(t, "", InstructionName(inst))
	}
}
