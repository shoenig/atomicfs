// Code autogenerated by mockery v2.0.0
//
// Do not manually edit the content of this file.

// Package fstest contains autogenerated mocks.
package fstest

import "indeed/gophers/3rdparty/p/github.com/stretchr/testify/mock"
import "os"
import "time"

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

// Chdir provides a mock function with given fields:
func (mockerySelf *File) Chdir() error {
	ret := mockerySelf.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chmod provides a mock function with given fields: mode
func (mockerySelf *File) Chmod(mode os.FileMode) error {
	ret := mockerySelf.Called(mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(os.FileMode) error); ok {
		r0 = rf(mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chown provides a mock function with given fields: uid, gid
func (mockerySelf *File) Chown(uid int, gid int) error {
	ret := mockerySelf.Called(uid, gid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(uid, gid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (mockerySelf *File) Close() error {
	ret := mockerySelf.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fd provides a mock function with given fields:
func (mockerySelf *File) Fd() uintptr {
	ret := mockerySelf.Called()

	var r0 uintptr
	if rf, ok := ret.Get(0).(func() uintptr); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uintptr)
	}

	return r0
}

// Name provides a mock function with given fields:
func (mockerySelf *File) Name() string {
	ret := mockerySelf.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Read provides a mock function with given fields: b
func (mockerySelf *File) Read(b []byte) (int, error) {
	ret := mockerySelf.Called(b)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAt provides a mock function with given fields: b, off
func (mockerySelf *File) ReadAt(b []byte, off int64) (int, error) {
	ret := mockerySelf.Called(b, off)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte, int64) int); ok {
		r0 = rf(b, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, int64) error); ok {
		r1 = rf(b, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readdir provides a mock function with given fields: n
func (mockerySelf *File) Readdir(n int) ([]os.FileInfo, error) {
	ret := mockerySelf.Called(n)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(int) []os.FileInfo); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readdirnames provides a mock function with given fields: n
func (mockerySelf *File) Readdirnames(n int) ([]string, error) {
	ret := mockerySelf.Called(n)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int) []string); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Seek provides a mock function with given fields: offset, whence
func (mockerySelf *File) Seek(offset int64, whence int) (int64, error) {
	ret := mockerySelf.Called(offset, whence)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int) int64); ok {
		r0 = rf(offset, whence)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(offset, whence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDeadline provides a mock function with given fields: t
func (mockerySelf *File) SetDeadline(t time.Time) error {
	ret := mockerySelf.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReadDeadline provides a mock function with given fields: t
func (mockerySelf *File) SetReadDeadline(t time.Time) error {
	ret := mockerySelf.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWriteDeadline provides a mock function with given fields: t
func (mockerySelf *File) SetWriteDeadline(t time.Time) error {
	ret := mockerySelf.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stat provides a mock function with given fields:
func (mockerySelf *File) Stat() (os.FileInfo, error) {
	ret := mockerySelf.Called()

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func() os.FileInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sync provides a mock function with given fields:
func (mockerySelf *File) Sync() error {
	ret := mockerySelf.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Truncate provides a mock function with given fields: size
func (mockerySelf *File) Truncate(size int64) error {
	ret := mockerySelf.Called(size)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: b
func (mockerySelf *File) Write(b []byte) (int, error) {
	ret := mockerySelf.Called(b)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteAt provides a mock function with given fields: b, off
func (mockerySelf *File) WriteAt(b []byte, off int64) (int, error) {
	ret := mockerySelf.Called(b, off)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte, int64) int); ok {
		r0 = rf(b, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, int64) error); ok {
		r1 = rf(b, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteString provides a mock function with given fields: s
func (mockerySelf *File) WriteString(s string) (int, error) {
	ret := mockerySelf.Called(s)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
