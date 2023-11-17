package wasi_snapshot_preview1

import (
	"os"
	"testing"

	"github.com/wasilibs/wazerox/experimental/sys"
	"github.com/wasilibs/wazerox/internal/sock"
	"github.com/wasilibs/wazerox/internal/testing/require"
	"github.com/wasilibs/wazerox/internal/wasip1"
)

func Test_getExtendedWasiFiletype(t *testing.T) {
	s := testSock{}
	ftype := getExtendedWasiFiletype(s, os.ModeIrregular)
	require.Equal(t, wasip1.FILETYPE_SOCKET_STREAM, ftype)

	c := testConn{}
	ftype = getExtendedWasiFiletype(c, os.ModeIrregular)
	require.Equal(t, wasip1.FILETYPE_SOCKET_STREAM, ftype)
}

type testSock struct {
	sys.UnimplementedFile
}

func (t testSock) Accept() (sock.TCPConn, sys.Errno) {
	panic("no-op")
}

type testConn struct {
	sys.UnimplementedFile
}

func (t testConn) Recvfrom([]byte, int) (n int, errno sys.Errno) {
	panic("no-op")
}

func (t testConn) Shutdown(int) sys.Errno {
	panic("no-op")
}
