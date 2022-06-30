//go:build linux || darwin || freebsd

// Package magic detects mimetypes using libmagic.
// This package requires libmagic, install it by the following
// commands below.
//	 - Debian or Ubuntu: apt-get install -y libmagic-dev
//	 - RHEL, CentOS or Fedora: yum install file-devel
//	 - Mac OS X: brew install libmagic
package magic

// #cgo LDFLAGS: -lmagic
// #include <stdlib.h>
// #include <magic.h>
import "C"
import (
	"sync"
	"unsafe"

	"errors"
)

type Magic struct {
	db  C.magic_t
	mux sync.Mutex
}

func NewMagic(flags Flag) (*Magic, error) {
	db := C.magic_open(C.int(0))
	if db == nil {
		return nil, errors.New("failed to open magic")
	}

	magic := &Magic{db: db}
	if ret := C.magic_setflags(db, C.int(flags)); int(ret) != 0 {
		magic.Close()
		return nil, NewError(magic.db)
	}

	if ret := C.magic_load(db, nil); int(ret) != 0 {
		magic.Close()
		return nil, NewError(magic.db)
	}
	return magic, nil
}

func (m *Magic) FromFile(filename string) (string, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	fn := C.CString(filename)
	defer C.free(unsafe.Pointer(fn))

	out := C.magic_file(m.db, fn)
	if out == nil {
		return "", NewError(m.db)
	}

	return C.GoString(out), nil
}

func (m *Magic) FromBuffer(buf []byte) (string, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	out := C.magic_buffer(m.db, unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if out == nil {
		return "", NewError(m.db)
	}

	return C.GoString(out), nil
}

func (m *Magic) FromDescriptor(fd int) (string, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	out := C.magic_descriptor(m.db, C.int(fd))
	if out == nil {
		return "", NewError(m.db)
	}

	return C.GoString(out), nil
}

func (m *Magic) SetFlags(flags Flag) int {
	m.mux.Lock()
	defer m.mux.Unlock()

	return int(C.magic_setflags(m.db, C.int(flags)))
}

func (m *Magic) GetFlags() int {
	m.mux.Lock()
	defer m.mux.Unlock()

	return int(C.magic_getflags(m.db))
}

func (m *Magic) Version() int {
	return int(C.magic_version())
}

func (m *Magic) Close() {
	m.mux.Lock()
	defer m.mux.Unlock()

	C.magic_close(m.db)
	m.db = nil
}

type Error struct {
	db C.magic_t
}

func NewError(db C.magic_t) error {
	return Error{db: db}
}

func (e Error) Error() string {
	return C.GoString(C.magic_error(e.db))
}
