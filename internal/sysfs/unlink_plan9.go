package sysfs

import (
	"syscall"

	"github.com/wasilibs/wazerox/experimental/sys"
)

func unlink(name string) sys.Errno {
	err := syscall.Remove(name)
	return sys.UnwrapOSError(err)
}
