// Code generated by 'ccgo stdlib/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o stdlib/stdlib_windows_386.go -pkgname stdlib', DO NOT EDIT.

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
	CHAR_BIT                                        = 8
	CHAR_MAX                                        = 127
	CHAR_MIN                                        = -128
	DUMMYSTRUCTNAME                                 = 0
	DUMMYSTRUCTNAME1                                = 0
	DUMMYSTRUCTNAME2                                = 0
	DUMMYSTRUCTNAME3                                = 0
	DUMMYSTRUCTNAME4                                = 0
	DUMMYSTRUCTNAME5                                = 0
	DUMMYUNIONNAME                                  = 0
	DUMMYUNIONNAME1                                 = 0
	DUMMYUNIONNAME2                                 = 0
	DUMMYUNIONNAME3                                 = 0
	DUMMYUNIONNAME4                                 = 0
	DUMMYUNIONNAME5                                 = 0
	DUMMYUNIONNAME6                                 = 0
	DUMMYUNIONNAME7                                 = 0
	DUMMYUNIONNAME8                                 = 0
	DUMMYUNIONNAME9                                 = 0
	EXIT_FAILURE                                    = 1
	EXIT_SUCCESS                                    = 0
	INT_MAX                                         = 2147483647
	INT_MIN                                         = -2147483648
	LLONG_MAX                                       = 9223372036854775807
	LLONG_MIN                                       = -9223372036854775808
	LONG_LONG_MAX                                   = 9223372036854775807
	LONG_LONG_MIN                                   = -9223372036854775808
	LONG_MAX                                        = 2147483647
	LONG_MIN                                        = -2147483648
	MB_LEN_MAX                                      = 5
	MINGW_DDK_H                                     = 0
	MINGW_HAS_DDK_H                                 = 1
	MINGW_HAS_SECURE_API                            = 1
	MINGW_SDK_INIT                                  = 0
	PATH_MAX                                        = 260
	RAND_MAX                                        = 0x7fff
	SCHAR_MAX                                       = 127
	SCHAR_MIN                                       = -128
	SHRT_MAX                                        = 32767
	SHRT_MIN                                        = -32768
	SIZE_MAX                                        = 4294967295
	SSIZE_MAX                                       = 2147483647
	UCHAR_MAX                                       = 255
	UINT_MAX                                        = 4294967295
	ULLONG_MAX                                      = 18446744073709551615
	ULONG_LONG_MAX                                  = 18446744073709551615
	ULONG_MAX                                       = 4294967295
	UNALIGNED                                       = 0
	USE___UUIDOF                                    = 0
	USHRT_MAX                                       = 65535
	WIN32                                           = 1
	WINNT                                           = 1
	X_AGLOBAL                                       = 0
	X_ALLOCA_S_HEAP_MARKER                          = 0xDDDD
	X_ALLOCA_S_MARKER_SIZE                          = 8
	X_ALLOCA_S_STACK_MARKER                         = 0xCCCC
	X_ALLOCA_S_THRESHOLD                            = 1024
	X_ANONYMOUS_STRUCT                              = 0
	X_ANONYMOUS_UNION                               = 0
	X_ARGMAX                                        = 100
	X_CALL_REPORTFAULT                              = 0x2
	X_CONST_RETURN                                  = 0
	X_CRTNOALIAS                                    = 0
	X_CRTRESTRICT                                   = 0
	X_CRT_ABS_DEFINED                               = 0
	X_CRT_ALGO_DEFINED                              = 0
	X_CRT_ALLOCATION_DEFINED                        = 0
	X_CRT_ALTERNATIVE_IMPORTED                      = 0
	X_CRT_ATOF_DEFINED                              = 0
	X_CRT_DOUBLE_DEC                                = 0
	X_CRT_ERRNO_DEFINED                             = 0
	X_CRT_MANAGED_HEAP_DEPRECATE                    = 0
	X_CRT_PACKING                                   = 8
	X_CRT_PERROR_DEFINED                            = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES          = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY   = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES        = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT  = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY = 0
	X_CRT_SWAB_DEFINED                              = 0
	X_CRT_SYSTEM_DEFINED                            = 0
	X_CRT_TERMINATE_DEFINED                         = 0
	X_CRT_USE_WINAPI_FAMILY_DESKTOP_APP             = 0
	X_CRT_WPERROR_DEFINED                           = 0
	X_CRT_WSYSTEM_DEFINED                           = 0
	X_CVTBUFSIZE                                    = 349
	X_DIV_T_DEFINED                                 = 0
	X_DLL                                           = 0
	X_ERRCODE_DEFINED                               = 0
	X_FILE_OFFSET_BITS                              = 64
	X_FREEA_INLINE                                  = 0
	X_FREEENTRY                                     = 0
	X_GCC_LIMITS_H_                                 = 0
	X_HEAPBADBEGIN                                  = -3
	X_HEAPBADNODE                                   = -4
	X_HEAPBADPTR                                    = -6
	X_HEAPEMPTY                                     = -1
	X_HEAPEND                                       = -5
	X_HEAPINFO_DEFINED                              = 0
	X_HEAPOK                                        = -2
	X_HEAP_MAXREQ                                   = 0xFFFFFFE0
	X_I16_MAX                                       = 32767
	X_I16_MIN                                       = -32768
	X_I32_MAX                                       = 2147483647
	X_I32_MIN                                       = -2147483648
	X_I64_MAX                                       = 9223372036854775807
	X_I64_MIN                                       = -9223372036854775808
	X_I8_MAX                                        = 127
	X_I8_MIN                                        = -128
	X_ILP32                                         = 1
	X_INC_CORECRT                                   = 0
	X_INC_CORECRT_WSTDLIB                           = 0
	X_INC_CRTDEFS                                   = 0
	X_INC_CRTDEFS_MACRO                             = 0
	X_INC_LIMITS                                    = 0
	X_INC_MINGW_SECAPI                              = 0
	X_INC_STDLIB                                    = 0
	X_INC_STDLIB_S                                  = 0
	X_INC_VADEFS                                    = 0
	X_INC__MINGW_H                                  = 0
	X_INT128_DEFINED                                = 0
	X_INTEGRAL_MAX_BITS                             = 64
	X_INTPTR_T_DEFINED                              = 0
	X_LIMITS_H___                                   = 0
	X_MALLOC_H_                                     = 0
	X_MAX_DIR                                       = 256
	X_MAX_DRIVE                                     = 3
	X_MAX_ENV                                       = 32767
	X_MAX_EXT                                       = 256
	X_MAX_FNAME                                     = 256
	X_MAX_PATH                                      = 260
	X_MAX_WAIT_MALLOC_CRT                           = 60000
	X_MM_MALLOC_H_INCLUDED                          = 0
	X_MT                                            = 0
	X_M_IX86                                        = 600
	X_ONEXIT_T_DEFINED                              = 0
	X_OUT_TO_DEFAULT                                = 0
	X_OUT_TO_MSGBOX                                 = 2
	X_OUT_TO_STDERR                                 = 1
	X_PGLOBAL                                       = 0
	X_PTRDIFF_T_                                    = 0
	X_PTRDIFF_T_DEFINED                             = 0
	X_QSORT_S_DEFINED                               = 0
	X_REPORT_ERRMODE                                = 3
	X_RSIZE_T_DEFINED                               = 0
	X_SECURECRT_FILL_BUFFER_PATTERN                 = 0xFD
	X_SIZE_T_DEFINED                                = 0
	X_SSIZE_T_DEFINED                               = 0
	X_TAGLC_ID_DEFINED                              = 0
	X_THREADLOCALEINFO                              = 0
	X_TIME32_T_DEFINED                              = 0
	X_TIME64_T_DEFINED                              = 0
	X_TIME_T_DEFINED                                = 0
	X_UI16_MAX                                      = 0xffff
	X_UI32_MAX                                      = 0xffffffff
	X_UI64_MAX                                      = 0xffffffffffffffff
	X_UI8_MAX                                       = 0xff
	X_UINTPTR_T_DEFINED                             = 0
	X_USEDENTRY                                     = 1
	X_USE_32BIT_TIME_T                              = 0
	X_VA_LIST_DEFINED                               = 0
	X_W64                                           = 0
	X_WCHAR_T_DEFINED                               = 0
	X_WCTYPE_T_DEFINED                              = 0
	X_WIN32                                         = 1
	X_WIN32_WINNT                                   = 0x502
	X_WINT_T                                        = 0
	X_WRITE_ABORT_MSG                               = 0x1
	X_WSTDLIBP_DEFINED                              = 0
	X_WSTDLIB_DEFINED                               = 0
	X_X86_                                          = 1
	I386                                            = 1
)

type Ptrdiff_t = int32

type Size_t = uint32

type Wchar_t = uint16

type X__builtin_va_list = uintptr
type X__float128 = float64

type Va_list = X__builtin_va_list

type X__gnuc_va_list = X__builtin_va_list

type Ssize_t = int32

type Rsize_t = Size_t

type Intptr_t = int32

type Uintptr_t = uint32

type Wint_t = uint16
type Wctype_t = uint16

type Errno_t = int32

type X__time32_t = int32

type X__time64_t = int64

type Time_t = X__time32_t

type Threadlocaleinfostruct = struct {
	Frefcount      int32
	Flc_codepage   uint32
	Flc_collate_cp uint32
	Flc_handle     [6]uint32
	Flc_id         [6]LC_ID
	Flc_category   [6]struct {
		Flocale    uintptr
		Fwlocale   uintptr
		Frefcount  uintptr
		Fwrefcount uintptr
	}
	Flc_clike            int32
	Fmb_cur_max          int32
	Flconv_intl_refcount uintptr
	Flconv_num_refcount  uintptr
	Flconv_mon_refcount  uintptr
	Flconv               uintptr
	Fctype1_refcount     uintptr
	Fctype1              uintptr
	Fpctype              uintptr
	Fpclmap              uintptr
	Fpcumap              uintptr
	Flc_time_curr        uintptr
}

type Pthreadlocinfo = uintptr
type Pthreadmbcinfo = uintptr

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
}

type X_locale_tstruct = Localeinfo_struct
type X_locale_t = uintptr

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
}

type LC_ID = TagLC_ID
type LPLC_ID = uintptr

type Threadlocinfo = Threadlocaleinfostruct

type X_onexit_t = uintptr

type X_div_t = struct {
	Fquot int32
	Frem  int32
}

type Div_t = X_div_t

type X_ldiv_t = struct {
	Fquot int32
	Frem  int32
}

type Ldiv_t = X_ldiv_t

type X_LDOUBLE = struct{ Fld [10]uint8 }

type X_CRT_DOUBLE = struct{ Fx float64 }

type X_CRT_FLOAT = struct{ Ff float32 }

type X_LONGDOUBLE = struct{ Fx float64 }

type X_LDBL12 = struct{ Fld12 [12]uint8 }

type X_purecall_handler = uintptr

type X_invalid_parameter_handler = uintptr

type Lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type X_heapinfo = struct {
	F_pentry  uintptr
	F_size    Size_t
	F_useflag int32
}

type X_HEAPINFO = X_heapinfo

var _ int8