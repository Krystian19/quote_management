// Code generated by 'ccgo poll/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o poll/poll_darwin_amd64.go -pkgname poll', DO NOT EDIT.

package poll

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	POLLATTRIB                             = 0x0400
	POLLERR                                = 0x0008
	POLLEXTEND                             = 0x0200
	POLLHUP                                = 0x0010
	POLLIN                                 = 0x0001
	POLLNLINK                              = 0x0800
	POLLNVAL                               = 0x0020
	POLLOUT                                = 0x0004
	POLLPRI                                = 0x0002
	POLLRDBAND                             = 0x0080
	POLLRDNORM                             = 0x0040
	POLLSTANDARD                           = 511
	POLLWRBAND                             = 0x0100
	POLLWRITE                              = 0x1000
	POLLWRNORM                             = 4
	X_CDEFS_H_                             = 0
	X_DARWIN_FEATURE_64_BIT_INODE          = 1
	X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
	X_DARWIN_FEATURE_UNIX_CONFORMANCE      = 3
	X_FILE_OFFSET_BITS                     = 64
	X_LP64                                 = 1
	X_Nonnull                              = 0
	X_Null_unspecified                     = 0
	X_Nullable                             = 0
	X_SYS_POLL_H_                          = 0
)

type Ptrdiff_t = int64

type Size_t = uint64

type Wchar_t = int32

type X__int128_t = struct {
	Flo int64
	Fhi int64
}
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
}

type X__builtin_va_list = uintptr
type X__float128 = float64

var X__darwin_check_fd_set_overflow uintptr

type Pollfd = struct {
	Ffd      int32
	Fevents  int16
	Frevents int16
}

type Nfds_t = uint32

var _ int8
