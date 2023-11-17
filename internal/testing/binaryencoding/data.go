package binaryencoding

import (
	"github.com/wasilibs/wazerox/internal/leb128"
	"github.com/wasilibs/wazerox/internal/wasm"
)

func encodeDataSegment(d *wasm.DataSegment) (ret []byte) {
	// Currently multiple memories are not supported.
	if d.Passive {
		ret = append(ret, leb128.EncodeInt32(1)...)
	} else {
		ret = append(ret, leb128.EncodeInt32(0)...) // active segment
		ret = append(ret, encodeConstantExpression(d.OffsetExpression)...)
	}
	ret = append(ret, leb128.EncodeUint32(uint32(len(d.Init)))...)
	ret = append(ret, d.Init...)
	return
}
