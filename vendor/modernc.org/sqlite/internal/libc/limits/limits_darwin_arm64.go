// Code generated by 'ccgo limits/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o limits/limits_darwin_arm64.go -pkgname limits', DO NOT EDIT.

package limits

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
	ARG_MAX                                = 1048576
	BC_BASE_MAX                            = 99
	BC_DIM_MAX                             = 2048
	BC_SCALE_MAX                           = 99
	BC_STRING_MAX                          = 1000
	CHARCLASS_NAME_MAX                     = 14
	CHAR_BIT                               = 8
	CHAR_MAX                               = 127
	CHAR_MIN                               = -128
	CHILD_MAX                              = 266
	CLK_TCK                                = 100
	COLL_WEIGHTS_MAX                       = 2
	EQUIV_CLASS_MAX                        = 2
	EXPR_NEST_MAX                          = 32
	GID_MAX                                = 2147483647
	INT_MAX                                = 2147483647
	INT_MIN                                = -2147483648
	IOV_MAX                                = 1024
	LINE_MAX                               = 2048
	LINK_MAX                               = 32767
	LLONG_MAX                              = 9223372036854775807
	LLONG_MIN                              = -9223372036854775808
	LONG_BIT                               = 64
	LONG_LONG_MAX                          = 9223372036854775807
	LONG_LONG_MIN                          = -9223372036854775808
	LONG_MAX                               = 9223372036854775807
	LONG_MIN                               = -9223372036854775808
	MAX_CANON                              = 1024
	MAX_INPUT                              = 1024
	MB_LEN_MAX                             = 6
	NAME_MAX                               = 255
	NGROUPS_MAX                            = 16
	NL_ARGMAX                              = 9
	NL_LANGMAX                             = 14
	NL_MSGMAX                              = 32767
	NL_NMAX                                = 1
	NL_SETMAX                              = 255
	NL_TEXTMAX                             = 2048
	NZERO                                  = 20
	OFF_MAX                                = 9223372036854775807
	OFF_MIN                                = -9223372036854775808
	OPEN_MAX                               = 10240
	PASS_MAX                               = 128
	PATH_MAX                               = 1024
	PIPE_BUF                               = 512
	PTHREAD_DESTRUCTOR_ITERATIONS          = 4
	PTHREAD_KEYS_MAX                       = 512
	PTHREAD_STACK_MIN                      = 16384
	QUAD_MAX                               = 9223372036854775807
	QUAD_MIN                               = -9223372036854775808
	RE_DUP_MAX                             = 255
	SCHAR_MAX                              = 127
	SCHAR_MIN                              = -128
	SHRT_MAX                               = 32767
	SHRT_MIN                               = -32768
	SIZE_T_MAX                             = 18446744073709551615
	SSIZE_MAX                              = 9223372036854775807
	UCHAR_MAX                              = 255
	UID_MAX                                = 2147483647
	UINT_MAX                               = 4294967295
	ULLONG_MAX                             = 18446744073709551615
	ULONG_LONG_MAX                         = 18446744073709551615
	ULONG_MAX                              = 18446744073709551615
	UQUAD_MAX                              = 18446744073709551615
	USHRT_MAX                              = 65535
	WORD_BIT                               = 32
	X_ARM_LIMITS_H_                        = 0
	X_ARM__LIMITS_H_                       = 0
	X_BSD_MACHINE_LIMITS_H_                = 0
	X_CDEFS_H_                             = 0
	X_DARWIN_FEATURE_64_BIT_INODE          = 1
	X_DARWIN_FEATURE_ONLY_64_BIT_INODE     = 1
	X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
	X_DARWIN_FEATURE_ONLY_VERS_1050        = 1
	X_DARWIN_FEATURE_UNIX_CONFORMANCE      = 3
	X_FILE_OFFSET_BITS                     = 64
	X_GCC_LIMITS_H_                        = 0
	X_LIMITS_H_                            = 0
	X_LIMITS_H___                          = 0
	X_LP64                                 = 1
	X_Nonnull                              = 0
	X_Null_unspecified                     = 0
	X_Nullable                             = 0
	X_POSIX2_BC_BASE_MAX                   = 99
	X_POSIX2_BC_DIM_MAX                    = 2048
	X_POSIX2_BC_SCALE_MAX                  = 99
	X_POSIX2_BC_STRING_MAX                 = 1000
	X_POSIX2_CHARCLASS_NAME_MAX            = 14
	X_POSIX2_COLL_WEIGHTS_MAX              = 2
	X_POSIX2_EQUIV_CLASS_MAX               = 2
	X_POSIX2_EXPR_NEST_MAX                 = 32
	X_POSIX2_LINE_MAX                      = 2048
	X_POSIX2_RE_DUP_MAX                    = 255
	X_POSIX_AIO_LISTIO_MAX                 = 2
	X_POSIX_AIO_MAX                        = 1
	X_POSIX_ARG_MAX                        = 4096
	X_POSIX_CHILD_MAX                      = 25
	X_POSIX_CLOCKRES_MIN                   = 20000000
	X_POSIX_DELAYTIMER_MAX                 = 32
	X_POSIX_HOST_NAME_MAX                  = 255
	X_POSIX_LINK_MAX                       = 8
	X_POSIX_LOGIN_NAME_MAX                 = 9
	X_POSIX_MAX_CANON                      = 255
	X_POSIX_MAX_INPUT                      = 255
	X_POSIX_MQ_OPEN_MAX                    = 8
	X_POSIX_MQ_PRIO_MAX                    = 32
	X_POSIX_NAME_MAX                       = 14
	X_POSIX_NGROUPS_MAX                    = 8
	X_POSIX_OPEN_MAX                       = 20
	X_POSIX_PATH_MAX                       = 256
	X_POSIX_PIPE_BUF                       = 512
	X_POSIX_RE_DUP_MAX                     = 255
	X_POSIX_RTSIG_MAX                      = 8
	X_POSIX_SEM_NSEMS_MAX                  = 256
	X_POSIX_SEM_VALUE_MAX                  = 32767
	X_POSIX_SIGQUEUE_MAX                   = 32
	X_POSIX_SSIZE_MAX                      = 32767
	X_POSIX_SS_REPL_MAX                    = 4
	X_POSIX_STREAM_MAX                     = 8
	X_POSIX_SYMLINK_MAX                    = 255
	X_POSIX_SYMLOOP_MAX                    = 8
	X_POSIX_THREAD_DESTRUCTOR_ITERATIONS   = 4
	X_POSIX_THREAD_KEYS_MAX                = 128
	X_POSIX_THREAD_THREADS_MAX             = 64
	X_POSIX_TIMER_MAX                      = 32
	X_POSIX_TRACE_EVENT_NAME_MAX           = 30
	X_POSIX_TRACE_NAME_MAX                 = 8
	X_POSIX_TRACE_SYS_MAX                  = 8
	X_POSIX_TRACE_USER_EVENT_MAX           = 32
	X_POSIX_TTY_NAME_MAX                   = 9
	X_POSIX_TZNAME_MAX                     = 6
	X_SYS_SYSLIMITS_H_                     = 0
	X_XOPEN_IOV_MAX                        = 16
	X_XOPEN_NAME_MAX                       = 255
	X_XOPEN_PATH_MAX                       = 1024
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

var _ int8