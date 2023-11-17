//go:build !windows && !plan9

package sysfs

import (
	"syscall"

	"github.com/wasilibs/wazerox/experimental/sys"
)

func unlink(name string) (errno sys.Errno) {
	err := syscall.Unlink(name)
	if errno = sys.UnwrapOSError(err); errno == sys.EPERM {
		errno = sys.EISDIR
	}
	return errno
}
