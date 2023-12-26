package sysfs

import (
	wazero "github.com/wasilibs/wazerox"
	experimentalsys "github.com/wasilibs/wazerox/experimental/sys"
)

// FSConfig extends wazero.FSConfig, allowing access to the experimental
// sys.FS until it is moved to the "sys" package.
type FSConfig interface {
	// WithSysFSMount assigns a sys.FS file system for any paths beginning at
	// `guestPath`.
	//
	// This is an alternative to WithFSMount, allowing more features.
	WithSysFSMount(fs experimentalsys.FS, guestPath string) wazero.FSConfig

	// WithRawPaths indicates the FS should accept raw paths from the
	// guest. Currently, only WASIx is known to use absolute or uncleaned paths
	// in FS calls. The experimentalsys.FS will need to be able to serve such
	// paths. Default fs.FS implementations will not be able to.
	WithRawPaths() wazero.FSConfig
}
