package sysfs_test

import (
	"io/fs"
	"testing/fstest"

	wazero "github.com/wasilibs/wazerox"
	"github.com/wasilibs/wazerox/experimental/sysfs"
)

var moduleConfig wazero.ModuleConfig

// This example shows how to adapt a fs.FS to a sys.FS
func ExampleAdaptFS() {
	m := fstest.MapFS{
		"a/b.txt": &fstest.MapFile{Mode: 0o666},
		".":       &fstest.MapFile{Mode: 0o777 | fs.ModeDir},
	}
	root := &sysfs.AdaptFS{FS: m}

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))
}

// This example shows how to configure a sysfs.DirFS
func ExampleDirFS() {
	root := sysfs.DirFS(".")

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))
}

// This example shows how to configure a sysfs.ReadFS
func ExampleReadFS() {
	root := sysfs.DirFS(".")
	readOnly := &sysfs.ReadFS{FS: root}

	moduleConfig = wazero.NewModuleConfig().
		WithFSConfig(wazero.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(readOnly, "/"))
}

// This example shows how to configure for use with raw paths.
func ExampleFSConfig_WithRawPaths() {
	root := sysfs.DirFS("/")
	fsConfig := wazero.NewFSConfig()
	fsConfig = fsConfig.(sysfs.FSConfig).WithSysFSMount(root, "/")
	fsConfig = fsConfig.(sysfs.FSConfig).WithRawPaths()
	moduleConfig = wazero.NewModuleConfig().WithFSConfig(fsConfig)
}
