// Code autogenerated by mockery v2.0.0
//
// Do not manually edit the content of this file.

// Package fstest contains autogenerated mocks.
package fstest

import "github.com/shoenig/atomicfs/fs"
import "indeed/gophers/3rdparty/p/github.com/stretchr/testify/mock"
import "os"

// FileSystem is an autogenerated mock type for the FileSystem type
type FileSystem struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (mockerySelf *FileSystem) Create(name string) (fs.File, error) {
	ret := mockerySelf.Called(name)

	var r0 fs.File
	if rf, ok := ret.Get(0).(func(string) fs.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mkdir provides a mock function with given fields: name, perm
func (mockerySelf *FileSystem) Mkdir(name string, perm os.FileMode) error {
	ret := mockerySelf.Called(name, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, os.FileMode) error); ok {
		r0 = rf(name, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MkdirAll provides a mock function with given fields: path, perm
func (mockerySelf *FileSystem) MkdirAll(path string, perm os.FileMode) error {
	ret := mockerySelf.Called(path, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, os.FileMode) error); ok {
		r0 = rf(path, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Open provides a mock function with given fields: name
func (mockerySelf *FileSystem) Open(name string) (fs.File, error) {
	ret := mockerySelf.Called(name)

	var r0 fs.File
	if rf, ok := ret.Get(0).(func(string) fs.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenFile provides a mock function with given fields: name, flag, perm
func (mockerySelf *FileSystem) OpenFile(name string, flag int, perm os.FileMode) (fs.File, error) {
	ret := mockerySelf.Called(name, flag, perm)

	var r0 fs.File
	if rf, ok := ret.Get(0).(func(string, int, os.FileMode) fs.File); ok {
		r0 = rf(name, flag, perm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, os.FileMode) error); ok {
		r1 = rf(name, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: name
func (mockerySelf *FileSystem) Remove(name string) error {
	ret := mockerySelf.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: path
func (mockerySelf *FileSystem) RemoveAll(path string) error {
	ret := mockerySelf.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rename provides a mock function with given fields: old, new
func (mockerySelf *FileSystem) Rename(old string, new string) error {
	ret := mockerySelf.Called(old, new)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(old, new)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stat provides a mock function with given fields: name
func (mockerySelf *FileSystem) Stat(name string) (os.FileInfo, error) {
	ret := mockerySelf.Called(name)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(string) os.FileInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}