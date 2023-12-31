//go:build !windows && !plan9

package sysfs

import (
	"io/fs"
	"syscall"

	experimentalsys "github.com/wasilibs/wazerox/experimental/sys"
	"github.com/wasilibs/wazerox/sys"
)

func inoFromFileInfo(_ string, info fs.FileInfo) (sys.Inode, experimentalsys.Errno) {
	switch v := info.Sys().(type) {
	case *sys.Stat_t:
		return v.Ino, 0
	case *syscall.Stat_t:
		return v.Ino, 0
	default:
		return 0, 0
	}
}
