// Code generated by 'ccgo stdlib/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o stdlib/stdlib_linux_loong64.go -pkgname stdlib', DO NOT EDIT.

package stdlib

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
	BIG_ENDIAN                   = 4321
	BYTE_ORDER                   = 1234
	EXIT_FAILURE                 = 1
	EXIT_SUCCESS                 = 0
	FD_SETSIZE                   = 1024
	LITTLE_ENDIAN                = 1234
	PDP_ENDIAN                   = 3412
	RAND_MAX                     = 2147483647
	WCONTINUED                   = 8
	WEXITED                      = 4
	WNOHANG                      = 1
	WNOWAIT                      = 0x01000000
	WSTOPPED                     = 2
	WUNTRACED                    = 2
	X_ABILP64                    = 3
	X_ALLOCA_H                   = 1
	X_ATFILE_SOURCE              = 1
	X_BITS_ATOMIC_WIDE_COUNTER_H = 0
	X_BITS_BYTESWAP_H            = 1
	X_BITS_ENDIANNESS_H          = 1
	X_BITS_ENDIAN_H              = 1
	X_BITS_FLOATN_COMMON_H       = 0
	X_BITS_FLOATN_H              = 0
	X_BITS_PTHREADTYPES_ARCH_H   = 1
	X_BITS_PTHREADTYPES_COMMON_H = 1
	X_BITS_STDINT_INTN_H         = 1
	X_BITS_TIME64_H              = 1
	X_BITS_TYPESIZES_H           = 1
	X_BITS_TYPES_H               = 1
	X_BITS_UINTN_IDENTITY_H      = 1
	X_BSD_SIZE_T_                = 0
	X_BSD_SIZE_T_DEFINED_        = 0
	X_DEFAULT_SOURCE             = 1
	X_ENDIAN_H                   = 1
	X_FEATURES_H                 = 1
	X_FILE_OFFSET_BITS           = 64
	X_GCC_SIZE_T                 = 0
	X_GCC_WCHAR_T                = 0
	X_LOONGARCH_ARCH             = "loongarch64"
	X_LOONGARCH_ARCH_LOONGARCH64 = 1
	X_LOONGARCH_FPSET            = 32
	X_LOONGARCH_SIM              = 3
	X_LOONGARCH_SPFPSET          = 32
	X_LOONGARCH_SZINT            = 32
	X_LOONGARCH_SZLONG           = 64
	X_LOONGARCH_SZPTR            = 64
	X_LOONGARCH_TUNE             = "la464"
	X_LOONGARCH_TUNE_LA464       = 1
	X_LP64                       = 1
	X_POSIX_C_SOURCE             = 200809
	X_POSIX_SOURCE               = 1
	X_SIZET_                     = 0
	X_SIZE_T                     = 0
	X_SIZE_T_                    = 0
	X_SIZE_T_DECLARED            = 0
	X_SIZE_T_DEFINED             = 0
	X_SIZE_T_DEFINED_            = 0
	X_STDC_PREDEF_H              = 1
	X_STDLIB_H                   = 1
	X_STRUCT_TIMESPEC            = 1
	X_SYS_CDEFS_H                = 1
	X_SYS_SELECT_H               = 1
	X_SYS_SIZE_T_H               = 0
	X_SYS_TYPES_H                = 1
	X_THREAD_MUTEX_INTERNAL_H    = 1
	X_THREAD_SHARED_TYPES_H      = 1
	X_T_SIZE                     = 0
	X_T_SIZE_                    = 0
	X_T_WCHAR                    = 0
	X_T_WCHAR_                   = 0
	X_WCHAR_T                    = 0
	X_WCHAR_T_                   = 0
	X_WCHAR_T_DECLARED           = 0
	X_WCHAR_T_DEFINED            = 0
	X_WCHAR_T_DEFINED_           = 0
	X_WCHAR_T_H                  = 0
	Linux                        = 1
	Unix                         = 1
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

type Div_t = struct {
	Fquot int32
	Frem  int32
}

type Ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type Lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type X__u_char = uint8
type X__u_short = uint16
type X__u_int = uint32
type X__u_long = uint64

type X__int8_t = int8
type X__uint8_t = uint8
type X__int16_t = int16
type X__uint16_t = uint16
type X__int32_t = int32
type X__uint32_t = uint32
type X__int64_t = int64
type X__uint64_t = uint64

type X__int_least8_t = int8
type X__uint_least8_t = uint8
type X__int_least16_t = int16
type X__uint_least16_t = uint16
type X__int_least32_t = int32
type X__uint_least32_t = uint32
type X__int_least64_t = int64
type X__uint_least64_t = uint64

type X__quad_t = int64
type X__u_quad_t = uint64

type X__intmax_t = int64
type X__uintmax_t = uint64

type X__dev_t = uint64
type X__uid_t = uint32
type X__gid_t = uint32
type X__ino_t = uint64
type X__ino64_t = uint64
type X__mode_t = uint32
type X__nlink_t = uint32
type X__off_t = int64
type X__off64_t = int64
type X__pid_t = int32
type X__fsid_t = struct{ F__val [2]int32 }
type X__clock_t = int64
type X__rlim_t = uint64
type X__rlim64_t = uint64
type X__id_t = uint32
type X__time_t = int64
type X__useconds_t = uint32
type X__suseconds_t = int64
type X__suseconds64_t = int64

type X__daddr_t = int32
type X__key_t = int32

type X__clockid_t = int32

type X__timer_t = uintptr

type X__blksize_t = int32

type X__blkcnt_t = int64
type X__blkcnt64_t = int64

type X__fsblkcnt_t = uint64
type X__fsblkcnt64_t = uint64

type X__fsfilcnt_t = uint64
type X__fsfilcnt64_t = uint64

type X__fsword_t = int64

type X__ssize_t = int64

type X__syscall_slong_t = int64

type X__syscall_ulong_t = uint64

type X__loff_t = int64
type X__caddr_t = uintptr

type X__intptr_t = int64

type X__socklen_t = uint32

type X__sig_atomic_t = int32

type U_char = uint8
type U_short = uint16
type U_int = uint32
type U_long = uint64
type Quad_t = int64
type U_quad_t = uint64
type Fsid_t = X__fsid_t
type Loff_t = int64

type Ino_t = uint64

type Dev_t = uint64

type Gid_t = uint32

type Mode_t = uint32

type Nlink_t = uint32

type Uid_t = uint32

type Off_t = int64

type Pid_t = int32

type Id_t = uint32

type Ssize_t = int64

type Daddr_t = int32
type Caddr_t = uintptr

type Key_t = int32

type Clock_t = int64

type Clockid_t = int32

type Time_t = int64

type Timer_t = uintptr

type Ulong = uint64
type Ushort = uint16
type Uint = uint32

type Int8_t = int8
type Int16_t = int16
type Int32_t = int32
type Int64_t = int64

type U_int8_t = uint8
type U_int16_t = uint16
type U_int32_t = uint32
type U_int64_t = uint64

type Register_t = int32

type X__sigset_t = struct{ F__val [16]uint64 }

type Sigset_t = X__sigset_t

type Timeval = struct {
	Ftv_sec  int64
	Ftv_usec int64
}

type Timespec = struct {
	Ftv_sec  int64
	Ftv_nsec int64
}

type Suseconds_t = int64

type X__fd_mask = int64

type Fd_set = struct{ F__fds_bits [16]int64 }

type Fd_mask = int64

type Blksize_t = int32

type Blkcnt_t = int64
type Fsblkcnt_t = uint64
type Fsfilcnt_t = uint64

type X__atomic_wide_counter = struct{ F__value64 uint64 }

type X__pthread_internal_list = struct {
	F__prev uintptr
	F__next uintptr
}

type X__pthread_list_t = X__pthread_internal_list

type X__pthread_internal_slist = struct{ F__next uintptr }

type X__pthread_slist_t = X__pthread_internal_slist

type X__pthread_mutex_s = struct {
	F__lock   int32
	F__count  uint32
	F__owner  int32
	F__nusers uint32
	F__kind   int32
	F__spins  int32
	F__list   X__pthread_list_t
}

type X__pthread_rwlock_arch_t = struct {
	F__readers       uint32
	F__writers       uint32
	F__wrphase_futex uint32
	F__writers_futex uint32
	F__pad3          uint32
	F__pad4          uint32
	F__flags         uint8
	F__shared        uint8
	F__pad1          uint8
	F__pad2          uint8
	F__cur_writer    int32
}

type X__pthread_cond_s = struct {
	F__wseq         X__atomic_wide_counter
	F__g1_start     X__atomic_wide_counter
	F__g_refs       [2]uint32
	F__g_size       [2]uint32
	F__g1_orig_size uint32
	F__wrefs        uint32
	F__g_signals    [2]uint32
}

type X__tss_t = uint32
type X__thrd_t = uint64

type X__once_flag = struct{ F__data int32 }

type Pthread_t = uint64

type Pthread_mutexattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]int8
}

type Pthread_condattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]int8
}

type Pthread_key_t = uint32

type Pthread_once_t = int32

type Pthread_attr_t1 = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [56]int8
}

type Pthread_attr_t = Pthread_attr_t1

type Pthread_mutex_t = struct{ F__data X__pthread_mutex_s }

type Pthread_cond_t = struct{ F__data X__pthread_cond_s }

type Pthread_rwlock_t = struct {
	F__ccgo_pad1 [0]uint64
	F__data      X__pthread_rwlock_arch_t
	F__ccgo_pad2 [24]byte
}

type Pthread_rwlockattr_t = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [8]int8
}

type Pthread_spinlock_t = int32

type Pthread_barrier_t = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [32]int8
}

type Pthread_barrierattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]int8
}

type Random_data = struct {
	Ffptr        uintptr
	Frptr        uintptr
	Fstate       uintptr
	Frand_type   int32
	Frand_deg    int32
	Frand_sep    int32
	F__ccgo_pad1 [4]byte
	Fend_ptr     uintptr
}

type Drand48_data = struct {
	F__x     [3]uint16
	F__old_x [3]uint16
	F__c     uint16
	F__init  uint16
	F__a     uint64
}

type X__compar_fn_t = uintptr

var _ int8
