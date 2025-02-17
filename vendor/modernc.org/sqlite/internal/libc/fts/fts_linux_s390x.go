// Code generated by 'ccgo fts/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o fts/fts_linux_s390x.go -pkgname fts', DO NOT EDIT.

package fts

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
	ACCESSPERMS                  = 511
	ALLPERMS                     = 4095
	BIG_ENDIAN                   = 4321
	BYTE_ORDER                   = 4321
	DEFFILEMODE                  = 438
	FD_SETSIZE                   = 1024
	FTS_AGAIN                    = 1
	FTS_COMFOLLOW                = 0x0001
	FTS_D                        = 1
	FTS_DC                       = 2
	FTS_DEFAULT                  = 3
	FTS_DNR                      = 4
	FTS_DONTCHDIR                = 0x01
	FTS_DOT                      = 5
	FTS_DP                       = 6
	FTS_ERR                      = 7
	FTS_F                        = 8
	FTS_FOLLOW                   = 2
	FTS_INIT                     = 9
	FTS_LOGICAL                  = 0x0002
	FTS_NAMEONLY                 = 0x0100
	FTS_NOCHDIR                  = 0x0004
	FTS_NOINSTR                  = 3
	FTS_NOSTAT                   = 0x0008
	FTS_NS                       = 10
	FTS_NSOK                     = 11
	FTS_OPTIONMASK               = 0x00ff
	FTS_PHYSICAL                 = 0x0010
	FTS_ROOTLEVEL                = 0
	FTS_ROOTPARENTLEVEL          = -1
	FTS_SEEDOT                   = 0x0020
	FTS_SKIP                     = 4
	FTS_SL                       = 12
	FTS_SLNONE                   = 13
	FTS_STOP                     = 0x0200
	FTS_SYMFOLLOW                = 0x02
	FTS_W                        = 14
	FTS_WHITEOUT                 = 0x0080
	FTS_XDEV                     = 0x0040
	LITTLE_ENDIAN                = 1234
	PDP_ENDIAN                   = 3412
	S_BLKSIZE                    = 512
	S_IEXEC                      = 64
	S_IFBLK                      = 24576
	S_IFCHR                      = 8192
	S_IFDIR                      = 16384
	S_IFIFO                      = 4096
	S_IFLNK                      = 40960
	S_IFMT                       = 61440
	S_IFREG                      = 32768
	S_IFSOCK                     = 49152
	S_IREAD                      = 256
	S_IRGRP                      = 32
	S_IROTH                      = 4
	S_IRUSR                      = 256
	S_IRWXG                      = 56
	S_IRWXO                      = 7
	S_IRWXU                      = 448
	S_ISGID                      = 1024
	S_ISUID                      = 2048
	S_ISVTX                      = 512
	S_IWGRP                      = 16
	S_IWOTH                      = 2
	S_IWRITE                     = 128
	S_IWUSR                      = 128
	S_IXGRP                      = 8
	S_IXOTH                      = 1
	S_IXUSR                      = 64
	UTIME_NOW                    = 1073741823
	UTIME_OMIT                   = 1073741822
	X_ATFILE_SOURCE              = 1
	X_BITS_BYTESWAP_H            = 1
	X_BITS_ENDIANNESS_H          = 1
	X_BITS_ENDIAN_H              = 1
	X_BITS_PTHREADTYPES_ARCH_H   = 1
	X_BITS_PTHREADTYPES_COMMON_H = 1
	X_BITS_STAT_H                = 1
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
	X_FTS_H                      = 1
	X_GCC_SIZE_T                 = 0
	X_LP64                       = 1
	X_MKNOD_VER                  = 0
	X_MKNOD_VER_LINUX            = 0
	X_POSIX_C_SOURCE             = 200809
	X_POSIX_SOURCE               = 1
	X_RWLOCK_INTERNAL_H          = 0
	X_SIZET_                     = 0
	X_SIZE_T                     = 0
	X_SIZE_T_                    = 0
	X_SIZE_T_DECLARED            = 0
	X_SIZE_T_DEFINED             = 0
	X_SIZE_T_DEFINED_            = 0
	X_STATBUF_ST_BLKSIZE         = 0
	X_STATBUF_ST_NSEC            = 0
	X_STATBUF_ST_RDEV            = 0
	X_STAT_VER                   = 1
	X_STAT_VER_KERNEL            = 0
	X_STAT_VER_LINUX             = 1
	X_STDC_PREDEF_H              = 1
	X_STRUCT_TIMESPEC            = 1
	X_SYS_CDEFS_H                = 1
	X_SYS_SELECT_H               = 1
	X_SYS_SIZE_T_H               = 0
	X_SYS_STAT_H                 = 1
	X_SYS_TYPES_H                = 1
	X_THREAD_MUTEX_INTERNAL_H    = 1
	X_THREAD_SHARED_TYPES_H      = 1
	X_T_SIZE                     = 0
	X_T_SIZE_                    = 0
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

type X__int_least8_t = X__int8_t
type X__uint_least8_t = X__uint8_t
type X__int_least16_t = X__int16_t
type X__uint_least16_t = X__uint16_t
type X__int_least32_t = X__int32_t
type X__uint_least32_t = X__uint32_t
type X__int_least64_t = X__int64_t
type X__uint_least64_t = X__uint64_t

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
type X__nlink_t = uint64
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

type X__daddr_t = int32
type X__key_t = int32

type X__clockid_t = int32

type X__timer_t = uintptr

type X__blksize_t = int64

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

type X__loff_t = X__off64_t
type X__caddr_t = uintptr

type X__intptr_t = int64

type X__socklen_t = uint32

type X__sig_atomic_t = int32

type U_char = X__u_char
type U_short = X__u_short
type U_int = X__u_int
type U_long = X__u_long
type Quad_t = X__quad_t
type U_quad_t = X__u_quad_t
type Fsid_t = X__fsid_t
type Loff_t = X__loff_t

type Ino_t = X__ino64_t

type Dev_t = X__dev_t

type Gid_t = X__gid_t

type Mode_t = X__mode_t

type Nlink_t = X__nlink_t

type Uid_t = X__uid_t

type Off_t = X__off64_t

type Pid_t = X__pid_t

type Id_t = X__id_t

type Ssize_t = X__ssize_t

type Daddr_t = X__daddr_t
type Caddr_t = X__caddr_t

type Key_t = X__key_t

type Clock_t = X__clock_t

type Clockid_t = X__clockid_t

type Time_t = X__time_t

type Timer_t = X__timer_t

type Ulong = uint64
type Ushort = uint16
type Uint = uint32

type Int8_t = X__int8_t
type Int16_t = X__int16_t
type Int32_t = X__int32_t
type Int64_t = X__int64_t

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

type Register_t = int32

type X__sigset_t = struct{ F__val [16]uint64 }

type Sigset_t = X__sigset_t

type Timeval = struct {
	Ftv_sec  X__time_t
	Ftv_usec X__suseconds_t
}

type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec X__syscall_slong_t
}

type Suseconds_t = X__suseconds_t

type X__fd_mask = int64

type Fd_set = struct{ F__fds_bits [16]X__fd_mask }

type Fd_mask = X__fd_mask

type Blksize_t = X__blksize_t

type Blkcnt_t = X__blkcnt64_t
type Fsblkcnt_t = X__fsblkcnt64_t
type Fsfilcnt_t = X__fsfilcnt64_t

type X__pthread_internal_list = struct {
	F__prev uintptr
	F__next uintptr
}

type X__pthread_list_t = X__pthread_internal_list

type X__pthread_internal_slist = struct{ F__next uintptr }

type X__pthread_slist_t = X__pthread_internal_slist

type X__pthread_mutex_s = struct {
	F__lock    int32
	F__count   uint32
	F__owner   int32
	F__nusers  uint32
	F__kind    int32
	F__spins   int16
	F__elision int16
	F__list    X__pthread_list_t
}

type X__pthread_rwlock_arch_t = struct {
	F__readers       uint32
	F__writers       uint32
	F__wrphase_futex uint32
	F__writers_futex uint32
	F__pad3          uint32
	F__pad4          uint32
	F__cur_writer    int32
	F__shared        int32
	F__pad1          uint64
	F__pad2          uint64
	F__flags         uint32
	F__ccgo_pad1     [4]byte
}

type X__pthread_cond_s = struct {
	F__0            struct{ F__wseq uint64 }
	F__8            struct{ F__g1_start uint64 }
	F__g_refs       [2]uint32
	F__g_size       [2]uint32
	F__g1_orig_size uint32
	F__wrefs        uint32
	F__g_signals    [2]uint32
}

type Pthread_t = uint64

type Pthread_mutexattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]uint8
}

type Pthread_condattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]uint8
}

type Pthread_key_t = uint32

type Pthread_once_t = int32

type Pthread_attr_t1 = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [56]uint8
}

type Pthread_attr_t = Pthread_attr_t1

type Pthread_mutex_t = struct{ F__data X__pthread_mutex_s }

type Pthread_cond_t = struct{ F__data X__pthread_cond_s }

type Pthread_rwlock_t = struct{ F__data X__pthread_rwlock_arch_t }

type Pthread_rwlockattr_t = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [8]uint8
}

type Pthread_spinlock_t = int32

type Pthread_barrier_t = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [32]uint8
}

type Pthread_barrierattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [4]uint8
}

type Stat = struct {
	Fst_dev            X__dev_t
	Fst_ino            X__ino_t
	Fst_nlink          X__nlink_t
	Fst_mode           X__mode_t
	Fst_uid            X__uid_t
	Fst_gid            X__gid_t
	F__glibc_reserved0 int32
	Fst_rdev           X__dev_t
	Fst_size           X__off_t
	Fst_atim           struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
	Fst_mtim struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
	Fst_ctim struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
	Fst_blksize       X__blksize_t
	Fst_blocks        X__blkcnt_t
	F__glibc_reserved [3]int64
}

type X_ftsent = struct {
	Ffts_cycle   uintptr
	Ffts_parent  uintptr
	Ffts_link    uintptr
	Ffts_number  int64
	Ffts_pointer uintptr
	Ffts_accpath uintptr
	Ffts_path    uintptr
	Ffts_errno   int32
	Ffts_symfd   int32
	Ffts_pathlen uint16
	Ffts_namelen uint16
	F__ccgo_pad1 [4]byte
	Ffts_ino     Ino_t
	Ffts_dev     Dev_t
	Ffts_nlink   Nlink_t
	Ffts_level   int16
	Ffts_info    uint16
	Ffts_flags   uint16
	Ffts_instr   uint16
	Ffts_statp   uintptr
	Ffts_name    [1]uint8
	F__ccgo_pad2 [7]byte
}

type FTS = struct {
	Ffts_cur     uintptr
	Ffts_child   uintptr
	Ffts_array   uintptr
	Ffts_dev     Dev_t
	Ffts_path    uintptr
	Ffts_rfd     int32
	Ffts_pathlen int32
	Ffts_nitems  int32
	F__ccgo_pad1 [4]byte
	Ffts_compar  uintptr
	Ffts_options int32
	F__ccgo_pad2 [4]byte
}

type FTSENT = X_ftsent

var _ uint8
