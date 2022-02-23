package vz

/*
#cgo darwin CFLAGS: -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization
# include "virtualization.h"
*/
import "C"
import "runtime"

type VirtioFileSystemDeviceConfiguration struct {
	pointer
}

type SingleDirectoryShare struct {
	pointer
}

func NewSingleDirectoryShare(directoryPath string, readOnly bool) *SingleDirectoryShare {
	dirPathChar := charWithGoString(directoryPath)
	defer dirPathChar.Free()

	config := &SingleDirectoryShare{
		pointer: pointer{
			ptr: C.newVZSingleDirectoryShare(
				dirPathChar.CString(),
				C.bool(readOnly),
			),
		},
	}
	runtime.SetFinalizer(config, func(self *SingleDirectoryShare) {
		self.Release()
	})
	return config
}

func NewVirtioFileSystemDeviceConfiguration(tag string, share *SingleDirectoryShare) *VirtioFileSystemDeviceConfiguration {
	tagName := charWithGoString(tag)
	defer tagName.Free()

	config := &VirtioFileSystemDeviceConfiguration{
		pointer: pointer{
			ptr: C.newVZVirtioFileSystemDeviceConfiguration(
				tagName.CString(),
				share.Ptr(),
			),
		},
	}
	runtime.SetFinalizer(config, func(self *VirtioFileSystemDeviceConfiguration) {
		self.Release()
	})
	return config
}
