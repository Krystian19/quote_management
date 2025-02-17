// Code generated by 'ccgo limits/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o limits/limits_illumos_amd64.go -pkgname limits', DO NOT EDIT.

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
	ARG_MAX                              = 2096640
	BC_BASE_MAX                          = 99
	BC_DIM_MAX                           = 2048
	BC_SCALE_MAX                         = 99
	BC_STRING_MAX                        = 1000
	CHARCLASS_NAME_MAX                   = 14
	CHAR_BIT                             = 8
	CHAR_MAX                             = 127
	CHAR_MIN                             = -128
	CHILD_MAX                            = 25
	COLL_WEIGHTS_MAX                     = 10
	DBL_DIG                              = 15
	DBL_MAX                              = 1.7976931348623157081452e+308
	DBL_MIN                              = 2.2250738585072013830903e-308
	EXPR_NEST_MAX                        = 32
	FCHR_MAX                             = 1048576
	FLT_DIG                              = 6
	FLT_MAX                              = 3.4028234663852885981170e+38
	FLT_MIN                              = 1.1754943508222875079688e-38
	INT16_MAX                            = 32767
	INT16_MIN                            = -32768
	INT32_MAX                            = 2147483647
	INT32_MIN                            = -2147483648
	INT64_MAX                            = 9223372036854775807
	INT64_MIN                            = -9223372036854775808
	INT8_MAX                             = 127
	INT8_MIN                             = -128
	INTMAX_MAX                           = 9223372036854775807
	INTMAX_MIN                           = -9223372036854775808
	INTPTR_MAX                           = 9223372036854775807
	INTPTR_MIN                           = -9223372036854775808
	INT_FAST16_MAX                       = 2147483647
	INT_FAST16_MIN                       = -2147483648
	INT_FAST32_MAX                       = 2147483647
	INT_FAST32_MIN                       = -2147483648
	INT_FAST64_MAX                       = 9223372036854775807
	INT_FAST64_MIN                       = -9223372036854775808
	INT_FAST8_MAX                        = 127
	INT_FAST8_MIN                        = -128
	INT_LEAST16_MAX                      = 32767
	INT_LEAST16_MIN                      = -32768
	INT_LEAST32_MAX                      = 2147483647
	INT_LEAST32_MIN                      = -2147483648
	INT_LEAST64_MAX                      = 9223372036854775807
	INT_LEAST64_MIN                      = -9223372036854775808
	INT_LEAST8_MAX                       = 127
	INT_LEAST8_MIN                       = -128
	INT_MAX                              = 2147483647
	INT_MIN                              = -2147483648
	IOV_MAX                              = 1024
	LINE_MAX                             = 2048
	LLONG_MAX                            = 9223372036854775807
	LLONG_MIN                            = -9223372036854775808
	LOGIN_NAME_MAX                       = 33
	LOGIN_NAME_MAX_TRAD                  = 9
	LOGNAME_MAX                          = 32
	LOGNAME_MAX_TRAD                     = 8
	LONG_BIT                             = 64
	LONG_LONG_MAX                        = 9223372036854775807
	LONG_LONG_MIN                        = -9223372036854775808
	LONG_MAX                             = 9223372036854775807
	LONG_MIN                             = -9223372036854775808
	MAX_CANON                            = 256
	MAX_INPUT                            = 512
	MB_LEN_MAX                           = 5
	NAME_MAX                             = 255
	NGROUPS_MAX                          = 16
	NL_ARGMAX                            = 9
	NL_LANGMAX                           = 14
	NL_MSGMAX                            = 32767
	NL_NMAX                              = 1
	NL_SETMAX                            = 255
	NL_TEXTMAX                           = 2048
	NZERO                                = 20
	OPEN_MAX                             = 256
	PASS_MAX                             = 256
	PATH_MAX                             = 1024
	PID_MAX                              = 999999
	PIPE_BUF                             = 5120
	PIPE_MAX                             = 5120
	PTRDIFF_MAX                          = 9223372036854775807
	PTRDIFF_MIN                          = -9223372036854775808
	RE_DUP_MAX                           = 255
	SCHAR_MAX                            = 127
	SCHAR_MIN                            = -128
	SHRT_MAX                             = 32767
	SHRT_MIN                             = -32768
	SIG_ATOMIC_MAX                       = 2147483647
	SIG_ATOMIC_MIN                       = -2147483648
	SIZE_MAX                             = 18446744073709551615
	SSIZE_MAX                            = 9223372036854775807
	STD_BLK                              = 1024
	SYMLINK_MAX                          = 1024
	SYSPID_MAX                           = 1
	SYS_NMLN                             = 257
	TMP_MAX                              = 17576
	TTYNAME_MAX                          = 128
	UCHAR_MAX                            = 255
	UID_MAX                              = 2147483647
	UINT16_MAX                           = 65535
	UINT32_MAX                           = 4294967295
	UINT64_MAX                           = 18446744073709551615
	UINT8_MAX                            = 255
	UINTMAX_MAX                          = 18446744073709551615
	UINTPTR_MAX                          = 18446744073709551615
	UINT_FAST16_MAX                      = 4294967295
	UINT_FAST32_MAX                      = 4294967295
	UINT_FAST64_MAX                      = 18446744073709551615
	UINT_FAST8_MAX                       = 255
	UINT_LEAST16_MAX                     = 65535
	UINT_LEAST32_MAX                     = 4294967295
	UINT_LEAST64_MAX                     = 18446744073709551615
	UINT_LEAST8_MAX                      = 255
	UINT_MAX                             = 4294967295
	ULLONG_MAX                           = 18446744073709551615
	ULONG_LONG_MAX                       = 18446744073709551615
	ULONG_MAX                            = 18446744073709551615
	USHRT_MAX                            = 65535
	USI_MAX                              = 4294967295
	WCHAR_MAX                            = 2147483647
	WCHAR_MIN                            = -2147483648
	WINT_MAX                             = 2147483647
	WINT_MIN                             = -2147483648
	WORD_BIT                             = 32
	X_ALIGNMENT_REQUIRED                 = 1
	X_ARG_MAX32                          = 1048320
	X_ARG_MAX64                          = 2096640
	X_BIT_FIELDS_LTOH                    = 0
	X_BOOL_ALIGNMENT                     = 1
	X_CHAR_ALIGNMENT                     = 1
	X_CHAR_IS_SIGNED                     = 0
	X_CLOCK_T                            = 0
	X_DMA_USES_PHYSADDR                  = 0
	X_DONT_USE_1275_GENERIC_NAMES        = 0
	X_DOUBLE_ALIGNMENT                   = 8
	X_DOUBLE_COMPLEX_ALIGNMENT           = 8
	X_DTRACE_VERSION                     = 1
	X_FILE_OFFSET_BITS                   = 64
	X_FIRMWARE_NEEDS_FDISK               = 0
	X_FLOAT_ALIGNMENT                    = 4
	X_FLOAT_COMPLEX_ALIGNMENT            = 4
	X_GCC_LIMITS_H_                      = 0
	X_HAVE_CPUID_INSN                    = 0
	X_IEEE_754                           = 0
	X_INT_ALIGNMENT                      = 4
	X_ISO_CPP_14882_1998                 = 0
	X_ISO_C_9899_1999                    = 0
	X_ISO_C_9899_2011                    = 0
	X_ISO_LIMITS_ISO_H                   = 0
	X_LARGEFILE64_SOURCE                 = 1
	X_LARGEFILE_SOURCE                   = 1
	X_LIMITS_H                           = 0
	X_LIMITS_H___                        = 0
	X_LITTLE_ENDIAN                      = 0
	X_LONGLONG_TYPE                      = 0
	X_LONG_ALIGNMENT                     = 8
	X_LONG_DOUBLE_ALIGNMENT              = 16
	X_LONG_DOUBLE_COMPLEX_ALIGNMENT      = 16
	X_LONG_LONG_ALIGNMENT                = 8
	X_LONG_LONG_ALIGNMENT_32             = 4
	X_LONG_LONG_LTOH                     = 0
	X_LP64                               = 1
	X_MAX_ALIGNMENT                      = 16
	X_MULTI_DATAMODEL                    = 0
	X_NORETURN_KYWD                      = 0
	X_PASS_MAX                           = 256
	X_PASS_MAX_XPG                       = 8
	X_POINTER_ALIGNMENT                  = 8
	X_POSIX2_BC_BASE_MAX                 = 99
	X_POSIX2_BC_DIM_MAX                  = 2048
	X_POSIX2_BC_SCALE_MAX                = 99
	X_POSIX2_BC_STRING_MAX               = 1000
	X_POSIX2_CHARCLASS_NAME_MAX          = 14
	X_POSIX2_COLL_WEIGHTS_MAX            = 2
	X_POSIX2_EXPR_NEST_MAX               = 32
	X_POSIX2_LINE_MAX                    = 2048
	X_POSIX2_RE_DUP_MAX                  = 255
	X_POSIX_AIO_LISTIO_MAX               = 2
	X_POSIX_AIO_MAX                      = 1
	X_POSIX_ARG_MAX                      = 4096
	X_POSIX_CHILD_MAX                    = 6
	X_POSIX_CLOCKRES_MIN                 = 20000000
	X_POSIX_DELAYTIMER_MAX               = 32
	X_POSIX_HOST_NAME_MAX                = 255
	X_POSIX_LINK_MAX                     = 8
	X_POSIX_LOGIN_NAME_MAX               = 9
	X_POSIX_MAX_CANON                    = 255
	X_POSIX_MAX_INPUT                    = 255
	X_POSIX_MQ_OPEN_MAX                  = 8
	X_POSIX_MQ_PRIO_MAX                  = 32
	X_POSIX_NAME_MAX                     = 14
	X_POSIX_NGROUPS_MAX                  = 0
	X_POSIX_OPEN_MAX                     = 16
	X_POSIX_PATH_MAX                     = 255
	X_POSIX_PIPE_BUF                     = 512
	X_POSIX_RE_DUP_MAX                   = 255
	X_POSIX_RTSIG_MAX                    = 8
	X_POSIX_SEM_NSEMS_MAX                = 256
	X_POSIX_SEM_VALUE_MAX                = 32767
	X_POSIX_SIGQUEUE_MAX                 = 32
	X_POSIX_SSIZE_MAX                    = 32767
	X_POSIX_STREAM_MAX                   = 8
	X_POSIX_SYMLINK_MAX                  = 255
	X_POSIX_SYMLOOP_MAX                  = 8
	X_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
	X_POSIX_THREAD_KEYS_MAX              = 128
	X_POSIX_THREAD_THREADS_MAX           = 64
	X_POSIX_TIMER_MAX                    = 32
	X_POSIX_TTY_NAME_MAX                 = 9
	X_POSIX_TZNAME_MAX                   = 3
	X_PSM_MODULES                        = 0
	X_RESTRICT_KYWD                      = 0
	X_RTC_CONFIG                         = 0
	X_SHORT_ALIGNMENT                    = 2
	X_SOFT_HOSTID                        = 0
	X_STACK_GROWS_DOWNWARD               = 0
	X_STDC_C11                           = 0
	X_STDC_C99                           = 0
	X_SUNOS_VTOC_16                      = 0
	X_SYS_CCOMPILE_H                     = 0
	X_SYS_FEATURE_TESTS_H                = 0
	X_SYS_INT_LIMITS_H                   = 0
	X_SYS_ISA_DEFS_H                     = 0
	X_SYS_LIMITS_H                       = 0
	X_XOPEN_IOV_MAX                      = 16
	X_XOPEN_NAME_MAX                     = 255
	X_XOPEN_PATH_MAX                     = 1024
	X_XOPEN_VERSION                      = 3
	Sun                                  = 1
	Unix                                 = 1
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

type Clock_t = int64

var _ int8
