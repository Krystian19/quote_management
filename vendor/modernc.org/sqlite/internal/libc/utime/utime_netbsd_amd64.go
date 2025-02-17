// Code generated by 'ccgo utime/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o utime/utime_netbsd_amd64.go -pkgname utime', DO NOT EDIT.

package utime

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
	X_AMD64_INT_TYPES_H_      = 0
	X_FILE_OFFSET_BITS        = 64
	X_LP64                    = 1
	X_SYS_CDEFS_ELF_H_        = 0
	X_SYS_CDEFS_H_            = 0
	X_SYS_COMMON_ANSI_H_      = 0
	X_SYS_COMMON_INT_TYPES_H_ = 0
	X_UTIME_H_                = 0
	X_X86_64_CDEFS_H_         = 0
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

type X__int8_t = int8
type X__uint8_t = uint8
type X__int16_t = int16
type X__uint16_t = uint16
type X__int32_t = int32
type X__uint32_t = uint32
type X__int64_t = int64
type X__uint64_t = uint64

type X__intptr_t = int64
type X__uintptr_t = uint64

type Time_t = X__int64_t

type Utimbuf = struct {
	Factime  Time_t
	Fmodtime Time_t
}

var _ int8
