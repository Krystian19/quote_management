// Code generated by 'ccgo unistd/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o unistd/unistd_freebsd_386.go -pkgname unistd', DO NOT EDIT.

package unistd

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
	BIG_ENDIAN                          = 4321
	BYTE_ORDER                          = 1234
	CLOSE_RANGE_CLOEXEC                 = 4
	FD_SETSIZE                          = 1024
	F_LOCK                              = 1
	F_OK                                = 0
	F_TEST                              = 3
	F_TLOCK                             = 2
	F_ULOCK                             = 0
	LITTLE_ENDIAN                       = 1234
	L_INCR                              = 1
	L_SET                               = 0
	L_XTND                              = 2
	PDP_ENDIAN                          = 3412
	RFCENVG                             = 2048
	RFCFDG                              = 4096
	RFCNAMEG                            = 1024
	RFENVG                              = 2
	RFFDG                               = 4
	RFFLAGS                             = 2416930932
	RFHIGHPID                           = 262144
	RFKERNELONLY                        = 268828672
	RFLINUXTHPN                         = 65536
	RFMEM                               = 32
	RFNAMEG                             = 1
	RFNOTEG                             = 8
	RFNOWAIT                            = 64
	RFPPWAIT                            = 2147483648
	RFPROC                              = 16
	RFPROCDESC                          = 268435456
	RFSIGSHARE                          = 16384
	RFSPAWN                             = 2147483648
	RFSTOPPED                           = 131072
	RFTHREAD                            = 8192
	RFTSIGMASK                          = 0xFF
	RFTSIGSHIFT                         = 20
	RFTSIGZMB                           = 524288
	R_OK                                = 0x04
	SEEK_CUR                            = 1
	SEEK_DATA                           = 3
	SEEK_END                            = 2
	SEEK_HOLE                           = 4
	SEEK_SET                            = 0
	STDERR_FILENO                       = 2
	STDIN_FILENO                        = 0
	STDOUT_FILENO                       = 1
	SWAPOFF_FORCE                       = 0x00000001
	W_OK                                = 0x02
	X_OK                                = 0x01
	X_ACCMODE_T_DECLARED                = 0
	X_BIG_ENDIAN                        = 4321
	X_BLKCNT_T_DECLARED                 = 0
	X_BLKSIZE_T_DECLARED                = 0
	X_BYTE_ORDER                        = 1234
	X_CAP_IOCTL_T_DECLARED              = 0
	X_CAP_RIGHTS_T_DECLARED             = 0
	X_CLOCKID_T_DECLARED                = 0
	X_CLOCK_T_DECLARED                  = 0
	X_CS_PATH                           = 1
	X_CS_POSIX_V6_ILP32_OFF32_CFLAGS    = 2
	X_CS_POSIX_V6_ILP32_OFF32_LDFLAGS   = 3
	X_CS_POSIX_V6_ILP32_OFF32_LIBS      = 4
	X_CS_POSIX_V6_ILP32_OFFBIG_CFLAGS   = 5
	X_CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS  = 6
	X_CS_POSIX_V6_ILP32_OFFBIG_LIBS     = 7
	X_CS_POSIX_V6_LP64_OFF64_CFLAGS     = 8
	X_CS_POSIX_V6_LP64_OFF64_LDFLAGS    = 9
	X_CS_POSIX_V6_LP64_OFF64_LIBS       = 10
	X_CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS   = 11
	X_CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS  = 12
	X_CS_POSIX_V6_LPBIG_OFFBIG_LIBS     = 13
	X_CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 14
	X_DEV_T_DECLARED                    = 0
	X_FFLAGS_T_DECLARED                 = 0
	X_FILE_OFFSET_BITS                  = 64
	X_FSBLKCNT_T_DECLARED               = 0
	X_FTRUNCATE_DECLARED                = 0
	X_GETOPT_DECLARED                   = 0
	X_GID_T_DECLARED                    = 0
	X_ID_T_DECLARED                     = 0
	X_ILP32                             = 1
	X_INO_T_DECLARED                    = 0
	X_INT16_T_DECLARED                  = 0
	X_INT32_T_DECLARED                  = 0
	X_INT64_T_DECLARED                  = 0
	X_INT8_T_DECLARED                   = 0
	X_INTMAX_T_DECLARED                 = 0
	X_INTPTR_T_DECLARED                 = 0
	X_IN_ADDR_T_DECLARED                = 0
	X_IN_PORT_T_DECLARED                = 0
	X_KEY_T_DECLARED                    = 0
	X_LITTLE_ENDIAN                     = 1234
	X_LSEEK_DECLARED                    = 0
	X_LWPID_T_DECLARED                  = 0
	X_MACHINE_ENDIAN_H_                 = 0
	X_MACHINE__LIMITS_H_                = 0
	X_MACHINE__TYPES_H_                 = 0
	X_MKDTEMP_DECLARED                  = 0
	X_MKNOD_DECLARED                    = 0
	X_MKSTEMP_DECLARED                  = 0
	X_MKTEMP_DECLARED                   = 0
	X_MMAP_DECLARED                     = 0
	X_MODE_T_DECLARED                   = 0
	X_MQD_T_DECLARED                    = 0
	X_NLINK_T_DECLARED                  = 0
	X_Nonnull                           = 0
	X_Null_unspecified                  = 0
	X_Nullable                          = 0
	X_OFF64_T_DECLARED                  = 0
	X_OFF_T_DECLARED                    = 0
	X_OPTRESET_DECLARED                 = 0
	X_PC_ACL_EXTENDED                   = 59
	X_PC_ACL_NFS4                       = 64
	X_PC_ACL_PATH_MAX                   = 60
	X_PC_ALLOC_SIZE_MIN                 = 10
	X_PC_ASYNC_IO                       = 53
	X_PC_CAP_PRESENT                    = 61
	X_PC_CHOWN_RESTRICTED               = 7
	X_PC_FILESIZEBITS                   = 12
	X_PC_INF_PRESENT                    = 62
	X_PC_LINK_MAX                       = 1
	X_PC_MAC_PRESENT                    = 63
	X_PC_MAX_CANON                      = 2
	X_PC_MAX_INPUT                      = 3
	X_PC_MIN_HOLE_SIZE                  = 21
	X_PC_NAME_MAX                       = 4
	X_PC_NO_TRUNC                       = 8
	X_PC_PATH_MAX                       = 5
	X_PC_PIPE_BUF                       = 6
	X_PC_PRIO_IO                        = 54
	X_PC_REC_INCR_XFER_SIZE             = 14
	X_PC_REC_MAX_XFER_SIZE              = 15
	X_PC_REC_MIN_XFER_SIZE              = 16
	X_PC_REC_XFER_ALIGN                 = 17
	X_PC_SYMLINK_MAX                    = 18
	X_PC_SYNC_IO                        = 55
	X_PC_VDISABLE                       = 9
	X_PDP_ENDIAN                        = 3412
	X_PID_T_DECLARED                    = 0
	X_POSIX2_CHAR_TERM                  = 1
	X_POSIX2_C_BIND                     = 200112
	X_POSIX2_C_DEV                      = -1
	X_POSIX2_FORT_DEV                   = -1
	X_POSIX2_FORT_RUN                   = 200112
	X_POSIX2_LOCALEDEF                  = -1
	X_POSIX2_PBS                        = -1
	X_POSIX2_PBS_ACCOUNTING             = -1
	X_POSIX2_PBS_CHECKPOINT             = -1
	X_POSIX2_PBS_LOCATE                 = -1
	X_POSIX2_PBS_MESSAGE                = -1
	X_POSIX2_PBS_TRACK                  = -1
	X_POSIX2_SW_DEV                     = -1
	X_POSIX2_UPE                        = 200112
	X_POSIX2_VERSION                    = 199212
	X_POSIX_ADVISORY_INFO               = 200112
	X_POSIX_ASYNCHRONOUS_IO             = 200112
	X_POSIX_BARRIERS                    = 200112
	X_POSIX_CHOWN_RESTRICTED            = 1
	X_POSIX_CLOCK_SELECTION             = -1
	X_POSIX_CPUTIME                     = 200112
	X_POSIX_FSYNC                       = 200112
	X_POSIX_IPV6                        = 0
	X_POSIX_JOB_CONTROL                 = 1
	X_POSIX_MAPPED_FILES                = 200112
	X_POSIX_MEMLOCK                     = -1
	X_POSIX_MEMLOCK_RANGE               = 200112
	X_POSIX_MEMORY_PROTECTION           = 200112
	X_POSIX_MESSAGE_PASSING             = 200112
	X_POSIX_MONOTONIC_CLOCK             = 200112
	X_POSIX_NO_TRUNC                    = 1
	X_POSIX_PRIORITIZED_IO              = -1
	X_POSIX_PRIORITY_SCHEDULING         = 0
	X_POSIX_RAW_SOCKETS                 = 200112
	X_POSIX_READER_WRITER_LOCKS         = 200112
	X_POSIX_REALTIME_SIGNALS            = 200112
	X_POSIX_REGEXP                      = 1
	X_POSIX_SEMAPHORES                  = 200112
	X_POSIX_SHARED_MEMORY_OBJECTS       = 200112
	X_POSIX_SHELL                       = 1
	X_POSIX_SPAWN                       = 200112
	X_POSIX_SPIN_LOCKS                  = 200112
	X_POSIX_SPORADIC_SERVER             = -1
	X_POSIX_SYNCHRONIZED_IO             = -1
	X_POSIX_THREADS                     = 200112
	X_POSIX_THREAD_ATTR_STACKADDR       = 200112
	X_POSIX_THREAD_ATTR_STACKSIZE       = 200112
	X_POSIX_THREAD_CPUTIME              = 200112
	X_POSIX_THREAD_PRIORITY_SCHEDULING  = 200112
	X_POSIX_THREAD_PRIO_INHERIT         = 200112
	X_POSIX_THREAD_PRIO_PROTECT         = 200112
	X_POSIX_THREAD_PROCESS_SHARED       = 200112
	X_POSIX_THREAD_SAFE_FUNCTIONS       = -1
	X_POSIX_THREAD_SPORADIC_SERVER      = -1
	X_POSIX_TIMEOUTS                    = 200112
	X_POSIX_TIMERS                      = 200112
	X_POSIX_TRACE                       = -1
	X_POSIX_TRACE_EVENT_FILTER          = -1
	X_POSIX_TRACE_INHERIT               = -1
	X_POSIX_TRACE_LOG                   = -1
	X_POSIX_TYPED_MEMORY_OBJECTS        = -1
	X_POSIX_VDISABLE                    = 0xff
	X_POSIX_VERSION                     = 200112
	X_PTHREAD_T_DECLARED                = 0
	X_QUAD_HIGHWORD                     = 1
	X_QUAD_LOWWORD                      = 0
	X_RLIM_T_DECLARED                   = 0
	X_SC_2_CHAR_TERM                    = 20
	X_SC_2_C_BIND                       = 18
	X_SC_2_C_DEV                        = 19
	X_SC_2_FORT_DEV                     = 21
	X_SC_2_FORT_RUN                     = 22
	X_SC_2_LOCALEDEF                    = 23
	X_SC_2_PBS                          = 59
	X_SC_2_PBS_ACCOUNTING               = 60
	X_SC_2_PBS_CHECKPOINT               = 61
	X_SC_2_PBS_LOCATE                   = 62
	X_SC_2_PBS_MESSAGE                  = 63
	X_SC_2_PBS_TRACK                    = 64
	X_SC_2_SW_DEV                       = 24
	X_SC_2_UPE                          = 25
	X_SC_2_VERSION                      = 17
	X_SC_ADVISORY_INFO                  = 65
	X_SC_AIO_LISTIO_MAX                 = 42
	X_SC_AIO_MAX                        = 43
	X_SC_AIO_PRIO_DELTA_MAX             = 44
	X_SC_ARG_MAX                        = 1
	X_SC_ASYNCHRONOUS_IO                = 28
	X_SC_ATEXIT_MAX                     = 107
	X_SC_BARRIERS                       = 66
	X_SC_BC_BASE_MAX                    = 9
	X_SC_BC_DIM_MAX                     = 10
	X_SC_BC_SCALE_MAX                   = 11
	X_SC_BC_STRING_MAX                  = 12
	X_SC_CHILD_MAX                      = 2
	X_SC_CLK_TCK                        = 3
	X_SC_CLOCK_SELECTION                = 67
	X_SC_COLL_WEIGHTS_MAX               = 13
	X_SC_CPUSET_SIZE                    = 122
	X_SC_CPUTIME                        = 68
	X_SC_DELAYTIMER_MAX                 = 45
	X_SC_EXPR_NEST_MAX                  = 14
	X_SC_FILE_LOCKING                   = 69
	X_SC_FSYNC                          = 38
	X_SC_GETGR_R_SIZE_MAX               = 70
	X_SC_GETPW_R_SIZE_MAX               = 71
	X_SC_HOST_NAME_MAX                  = 72
	X_SC_IOV_MAX                        = 56
	X_SC_IPV6                           = 118
	X_SC_JOB_CONTROL                    = 6
	X_SC_LINE_MAX                       = 15
	X_SC_LOGIN_NAME_MAX                 = 73
	X_SC_MAPPED_FILES                   = 29
	X_SC_MEMLOCK                        = 30
	X_SC_MEMLOCK_RANGE                  = 31
	X_SC_MEMORY_PROTECTION              = 32
	X_SC_MESSAGE_PASSING                = 33
	X_SC_MONOTONIC_CLOCK                = 74
	X_SC_MQ_OPEN_MAX                    = 46
	X_SC_MQ_PRIO_MAX                    = 75
	X_SC_NGROUPS_MAX                    = 4
	X_SC_NPROCESSORS_CONF               = 57
	X_SC_NPROCESSORS_ONLN               = 58
	X_SC_OPEN_MAX                       = 5
	X_SC_PAGESIZE                       = 47
	X_SC_PAGE_SIZE                      = 47
	X_SC_PHYS_PAGES                     = 121
	X_SC_PRIORITIZED_IO                 = 34
	X_SC_PRIORITY_SCHEDULING            = 35
	X_SC_RAW_SOCKETS                    = 119
	X_SC_READER_WRITER_LOCKS            = 76
	X_SC_REALTIME_SIGNALS               = 36
	X_SC_REGEXP                         = 77
	X_SC_RE_DUP_MAX                     = 16
	X_SC_RTSIG_MAX                      = 48
	X_SC_SAVED_IDS                      = 7
	X_SC_SEMAPHORES                     = 37
	X_SC_SEM_NSEMS_MAX                  = 49
	X_SC_SEM_VALUE_MAX                  = 50
	X_SC_SHARED_MEMORY_OBJECTS          = 39
	X_SC_SHELL                          = 78
	X_SC_SIGQUEUE_MAX                   = 51
	X_SC_SPAWN                          = 79
	X_SC_SPIN_LOCKS                     = 80
	X_SC_SPORADIC_SERVER                = 81
	X_SC_STREAM_MAX                     = 26
	X_SC_SYMLOOP_MAX                    = 120
	X_SC_SYNCHRONIZED_IO                = 40
	X_SC_THREADS                        = 96
	X_SC_THREAD_ATTR_STACKADDR          = 82
	X_SC_THREAD_ATTR_STACKSIZE          = 83
	X_SC_THREAD_CPUTIME                 = 84
	X_SC_THREAD_DESTRUCTOR_ITERATIONS   = 85
	X_SC_THREAD_KEYS_MAX                = 86
	X_SC_THREAD_PRIORITY_SCHEDULING     = 89
	X_SC_THREAD_PRIO_INHERIT            = 87
	X_SC_THREAD_PRIO_PROTECT            = 88
	X_SC_THREAD_PROCESS_SHARED          = 90
	X_SC_THREAD_SAFE_FUNCTIONS          = 91
	X_SC_THREAD_SPORADIC_SERVER         = 92
	X_SC_THREAD_STACK_MIN               = 93
	X_SC_THREAD_THREADS_MAX             = 94
	X_SC_TIMEOUTS                       = 95
	X_SC_TIMERS                         = 41
	X_SC_TIMER_MAX                      = 52
	X_SC_TRACE                          = 97
	X_SC_TRACE_EVENT_FILTER             = 98
	X_SC_TRACE_INHERIT                  = 99
	X_SC_TRACE_LOG                      = 100
	X_SC_TTY_NAME_MAX                   = 101
	X_SC_TYPED_MEMORY_OBJECTS           = 102
	X_SC_TZNAME_MAX                     = 27
	X_SC_V6_ILP32_OFF32                 = 103
	X_SC_V6_ILP32_OFFBIG                = 104
	X_SC_V6_LP64_OFF64                  = 105
	X_SC_V6_LPBIG_OFFBIG                = 106
	X_SC_VERSION                        = 8
	X_SC_XOPEN_CRYPT                    = 108
	X_SC_XOPEN_ENH_I18N                 = 109
	X_SC_XOPEN_LEGACY                   = 110
	X_SC_XOPEN_REALTIME                 = 111
	X_SC_XOPEN_REALTIME_THREADS         = 112
	X_SC_XOPEN_SHM                      = 113
	X_SC_XOPEN_STREAMS                  = 114
	X_SC_XOPEN_UNIX                     = 115
	X_SC_XOPEN_VERSION                  = 116
	X_SC_XOPEN_XCU_VERSION              = 117
	X_SELECT_DECLARED                   = 0
	X_SIGSET_T_DECLARED                 = 0
	X_SIG_MAXSIG                        = 128
	X_SIG_WORDS                         = 4
	X_SIZE_T_DECLARED                   = 0
	X_SSIZE_T_DECLARED                  = 0
	X_SUSECONDS_T_DECLARED              = 0
	X_SWAB_DECLARED                     = 0
	X_SYS_CDEFS_H_                      = 0
	X_SYS_SELECT_H_                     = 0
	X_SYS_TIMESPEC_H_                   = 0
	X_SYS_TYPES_H_                      = 0
	X_SYS_UNISTD_H_                     = 0
	X_SYS__ENDIAN_H_                    = 0
	X_SYS__PTHREADTYPES_H_              = 0
	X_SYS__SIGSET_H_                    = 0
	X_SYS__STDINT_H_                    = 0
	X_SYS__TIMESPEC_H_                  = 0
	X_SYS__TIMEVAL_H_                   = 0
	X_SYS__TYPES_H_                     = 0
	X_TIMER_T_DECLARED                  = 0
	X_TIME_T_DECLARED                   = 0
	X_TRUNCATE_DECLARED                 = 0
	X_UID_T_DECLARED                    = 0
	X_UINT16_T_DECLARED                 = 0
	X_UINT32_T_DECLARED                 = 0
	X_UINT64_T_DECLARED                 = 0
	X_UINT8_T_DECLARED                  = 0
	X_UINTMAX_T_DECLARED                = 0
	X_UINTPTR_T_DECLARED                = 0
	X_UNISTD_H_                         = 0
	X_USECONDS_T_DECLARED               = 0
	X_V6_ILP32_OFF32                    = -1
	X_V6_ILP32_OFFBIG                   = 0
	X_V6_LP64_OFF64                     = 0
	X_V6_LPBIG_OFFBIG                   = -1
	X_XOPEN_CRYPT                       = -1
	X_XOPEN_ENH_I18N                    = -1
	X_XOPEN_LEGACY                      = -1
	X_XOPEN_REALTIME                    = -1
	X_XOPEN_REALTIME_THREADS            = -1
	X_XOPEN_SHM                         = 1
	X_XOPEN_STREAMS                     = -1
	X_XOPEN_UNIX                        = -1
	I386                                = 1
	Unix                                = 1
)

type Ptrdiff_t = int32

type Size_t = uint32

type Wchar_t = int32

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

type X__clock_t = uint32
type X__critical_t = X__int32_t
type X__double_t = float64
type X__float_t = float64
type X__intfptr_t = X__int32_t
type X__intptr_t = X__int32_t
type X__intmax_t = X__int64_t
type X__int_fast8_t = X__int32_t
type X__int_fast16_t = X__int32_t
type X__int_fast32_t = X__int32_t
type X__int_fast64_t = X__int64_t
type X__int_least8_t = X__int8_t
type X__int_least16_t = X__int16_t
type X__int_least32_t = X__int32_t
type X__int_least64_t = X__int64_t
type X__ptrdiff_t = X__int32_t
type X__register_t = X__int32_t
type X__segsz_t = X__int32_t
type X__size_t = X__uint32_t
type X__ssize_t = X__int32_t
type X__time_t = X__int32_t
type X__uintfptr_t = X__uint32_t
type X__uintptr_t = X__uint32_t
type X__uintmax_t = X__uint64_t
type X__uint_fast8_t = X__uint32_t
type X__uint_fast16_t = X__uint32_t
type X__uint_fast32_t = X__uint32_t
type X__uint_fast64_t = X__uint64_t
type X__uint_least8_t = X__uint8_t
type X__uint_least16_t = X__uint16_t
type X__uint_least32_t = X__uint32_t
type X__uint_least64_t = X__uint64_t
type X__u_register_t = X__uint32_t
type X__vm_offset_t = X__uint32_t
type X__vm_paddr_t = X__uint64_t
type X__vm_size_t = X__uint32_t
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
type X__key_t = int32
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
type X__suseconds_t = int32
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
	F__ccgo_pad1 [0]uint32
	F__mbstate8  [128]int8
}

type X__rman_res_t = X__uintmax_t

type X__va_list = X__builtin_va_list
type X__gnuc_va_list = X__va_list
type Pthread_once = struct {
	Fstate int32
	Fmutex Pthread_mutex_t
}

type Pthread_t = uintptr
type Pthread_attr_t = uintptr
type Pthread_mutex_t = uintptr
type Pthread_mutexattr_t = uintptr
type Pthread_cond_t = uintptr
type Pthread_condattr_t = uintptr
type Pthread_key_t = int32
type Pthread_once_t = Pthread_once
type Pthread_rwlock_t = uintptr
type Pthread_rwlockattr_t = uintptr
type Pthread_barrier_t = uintptr
type Pthread_barrierattr_t = uintptr
type Pthread_spinlock_t = uintptr

type Pthread_addr_t = uintptr
type Pthread_startroutine_t = uintptr

type U_char = uint8
type U_short = uint16
type U_int = uint32
type U_long = uint32
type Ushort = uint16
type Uint = uint32

type Int8_t = X__int8_t

type Int16_t = X__int16_t

type Int32_t = X__int32_t

type Int64_t = X__int64_t

type Uint8_t = X__uint8_t

type Uint16_t = X__uint16_t

type Uint32_t = X__uint32_t

type Uint64_t = X__uint64_t

type Intptr_t = X__intptr_t
type Uintptr_t = X__uintptr_t
type Intmax_t = X__intmax_t
type Uintmax_t = X__uintmax_t

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

type U_quad_t = X__uint64_t
type Quad_t = X__int64_t
type Qaddr_t = uintptr

type Caddr_t = uintptr
type C_caddr_t = uintptr

type Blksize_t = X__blksize_t

type Cpuwhich_t = X__cpuwhich_t
type Cpulevel_t = X__cpulevel_t
type Cpusetid_t = X__cpusetid_t

type Blkcnt_t = X__blkcnt_t

type Clock_t = X__clock_t

type Clockid_t = X__clockid_t

type Critical_t = X__critical_t
type Daddr_t = X__daddr_t

type Dev_t = X__dev_t

type Fflags_t = X__fflags_t

type Fixpt_t = X__fixpt_t

type Fsblkcnt_t = X__fsblkcnt_t
type Fsfilcnt_t = X__fsfilcnt_t

type Gid_t = X__gid_t

type In_addr_t = X__uint32_t

type In_port_t = X__uint16_t

type Id_t = X__id_t

type Ino_t = X__ino_t

type Key_t = X__key_t

type Lwpid_t = X__lwpid_t

type Mode_t = X__mode_t

type Accmode_t = X__accmode_t

type Nlink_t = X__nlink_t

type Off_t = X__off_t

type Off64_t = X__off64_t

type Pid_t = X__pid_t

type Register_t = X__register_t

type Rlim_t = X__rlim_t

type Sbintime_t = X__int64_t

type Segsz_t = X__segsz_t

type Ssize_t = X__ssize_t

type Suseconds_t = X__suseconds_t

type Time_t = X__time_t

type Timer_t = X__timer_t

type Mqd_t = X__mqd_t

type U_register_t = X__u_register_t

type Uid_t = X__uid_t

type Useconds_t = X__useconds_t

type Cap_ioctl_t = uint32

type Kpaddr_t = X__uint64_t
type Kvaddr_t = X__uint64_t
type Ksize_t = X__uint64_t
type Kssize_t = X__int64_t

type Vm_offset_t = X__vm_offset_t
type Vm_ooffset_t = X__uint64_t
type Vm_paddr_t = X__vm_paddr_t
type Vm_pindex_t = X__uint64_t
type Vm_size_t = X__vm_size_t

type Rman_res_t = X__rman_res_t

type X__sigset = struct{ F__bits [4]X__uint32_t }

type X__sigset_t = X__sigset

type Timeval = struct {
	Ftv_sec  Time_t
	Ftv_usec Suseconds_t
}

type Timespec = struct {
	Ftv_sec  Time_t
	Ftv_nsec int32
}

type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  Time_t
		Ftv_nsec int32
	}
	Fit_value struct {
		Ftv_sec  Time_t
		Ftv_nsec int32
	}
}

type X__fd_mask = uint32
type Fd_mask = X__fd_mask

type Sigset_t = X__sigset_t

type Fd_set1 = struct{ F__fds_bits [32]X__fd_mask }

type Fd_set = Fd_set1

type Crypt_data = struct {
	Finitialized int32
	F__buf       [256]int8
}

var _ int8
