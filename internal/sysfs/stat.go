package sysfs

import (
	"io/fs"

	experimentalsys "github.com/wasilibs/wazerox/experimental/sys"
	"github.com/wasilibs/wazerox/sys"
)

func defaultStatFile(f fs.File) (sys.Stat_t, experimentalsys.Errno) {
	if info, err := f.Stat(); err != nil {
		return sys.Stat_t{}, experimentalsys.UnwrapOSError(err)
	} else {
		return sys.NewStat_t(info), 0
	}
}
