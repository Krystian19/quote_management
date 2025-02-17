// Code generated by 'ccgo utime/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o utime/utime_freebsd_amd64.go -pkgname utime', DO NOT EDIT.

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
	X_FILE_OFFSET_BITS   = 64
	X_LP64               = 1
	X_MACHINE__LIMITS_H_ = 0
	X_MACHINE__TYPES_H_  = 0
	X_Nonnull            = 0
	X_Null_unspecified   = 0
	X_Nullable           = 0
	X_SYS_CDEFS_H_       = 0
	X_SYS__TYPES_H_      = 0
	X_TIME_T_DECLARED    = 0
	X_UTIME_H_           = 0
	Unix                 = 1
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

type X__clock_t = X__int32_t
type X__critical_t = X__int64_t
type X__double_t = float64
type X__float_t = float32
type X__intfptr_t = X__int64_t
type X__intptr_t = X__int64_t
type X__intmax_t = X__int64_t
type X__int_fast8_t = X__int32_t
type X__int_fast16_t = X__int32_t
type X__int_fast32_t = X__int32_t
type X__int_fast64_t = X__int64_t
type X__int_least8_t = X__int8_t
type X__int_least16_t = X__int16_t
type X__int_least32_t = X__int32_t
type X__int_least64_t = X__int64_t
type X__ptrdiff_t = X__int64_t
type X__register_t = X__int64_t
type X__segsz_t = X__int64_t
type X__size_t = X__uint64_t
type X__ssize_t = X__int64_t
type X__time_t = X__int64_t
type X__uintfptr_t = X__uint64_t
type X__uintptr_t = X__uint64_t
type X__uintmax_t = X__uint64_t
type X__uint_fast8_t = X__uint32_t
type X__uint_fast16_t = X__uint32_t
type X__uint_fast32_t = X__uint32_t
type X__uint_fast64_t = X__uint64_t
type X__uint_least8_t = X__uint8_t
type X__uint_least16_t = X__uint16_t
type X__uint_least32_t = X__uint32_t
type X__uint_least64_t = X__uint64_t
type X__u_register_t = X__uint64_t
type X__vm_offset_t = X__uint64_t
type X__vm_paddr_t = X__uint64_t
type X__vm_size_t = X__uint64_t
type X___wchar_t = int32

type X__blksize_t = X__int32_t
type X__blkcnt_t = X__int64_t
type X__clockid_t = X__int32_t
type X__fflags_t = X__uint32_t
type X__fsblkcnt_t = X__uint64_t
type X__fsfilcnt_t = X__uint64_t
type X__gid_t = X__uint32_t
type X__id_t = X__int64_t
type X__ino_t = X__uint64_t
type X__key_t = int64
type X__lwpid_t = X__int32_t
type X__mode_t = X__uint16_t
type X__accmode_t = int32
type X__nl_item = int32
type X__nlink_t = X__uint64_t
type X__off_t = X__int64_t
type X__off64_t = X__int64_t
type X__pid_t = X__int32_t
type X__rlim_t = X__int64_t

type X__sa_family_t = X__uint8_t
type X__socklen_t = X__uint32_t
type X__suseconds_t = int64
type X__timer_t = uintptr
type X__mqd_t = uintptr
type X__uid_t = X__uint32_t
type X__useconds_t = uint32
type X__cpuwhich_t = int32
type X__cpulevel_t = int32
type X__cpusetid_t = int32
type X__daddr_t = X__int64_t

type X__ct_rune_t = int32
type X__rune_t = X__ct_rune_t
type X__wint_t = X__ct_rune_t

type X__char16_t = X__uint_least16_t
type X__char32_t = X__uint_least32_t

type X__max_align_t = struct {
	F__max_align1 int64
	F__max_align2 float64
}

type X__dev_t = X__uint64_t

type X__fixpt_t = X__uint32_t

type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
}

type X__rman_res_t = X__uintmax_t

type X__va_list = X__builtin_va_list
type X__gnuc_va_list = X__va_list

type Time_t = X__time_t

type Utimbuf = struct {
	Factime  Time_t
	Fmodtime Time_t
}

var _ int8
