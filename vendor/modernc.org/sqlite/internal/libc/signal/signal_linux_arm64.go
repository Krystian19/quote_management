// Code generated by 'ccgo signal/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o signal/signal_linux_arm64.go -pkgname signal', DO NOT EDIT.

package signal

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
	ELF_PRARGSZ                  = 80
	ESR_MAGIC                    = 0x45535201
	EXTRA_MAGIC                  = 0x45585401
	FD_SETSIZE                   = 1024
	FPSIMD_MAGIC                 = 0x46508001
	LITTLE_ENDIAN                = 1234
	MINSIGSTKSZ                  = 5120
	NSIG                         = 65
	PDP_ENDIAN                   = 3412
	SA_INTERRUPT                 = 0x20000000
	SA_NOCLDSTOP                 = 1
	SA_NOCLDWAIT                 = 2
	SA_NODEFER                   = 0x40000000
	SA_NOMASK                    = 1073741824
	SA_ONESHOT                   = 2147483648
	SA_ONSTACK                   = 0x08000000
	SA_RESETHAND                 = 0x80000000
	SA_RESTART                   = 0x10000000
	SA_SIGINFO                   = 4
	SA_STACK                     = 134217728
	SIGABRT                      = 6
	SIGALRM                      = 14
	SIGBUS                       = 7
	SIGCHLD                      = 17
	SIGCLD                       = 17
	SIGCONT                      = 18
	SIGFPE                       = 8
	SIGHUP                       = 1
	SIGILL                       = 4
	SIGINT                       = 2
	SIGIO                        = 29
	SIGIOT                       = 6
	SIGKILL                      = 9
	SIGPIPE                      = 13
	SIGPOLL                      = 29
	SIGPROF                      = 27
	SIGPWR                       = 30
	SIGQUIT                      = 3
	SIGSEGV                      = 11
	SIGSTKFLT                    = 16
	SIGSTKSZ                     = 16384
	SIGSTOP                      = 19
	SIGSYS                       = 31
	SIGTERM                      = 15
	SIGTRAP                      = 5
	SIGTSTP                      = 20
	SIGTTIN                      = 21
	SIGTTOU                      = 22
	SIGURG                       = 23
	SIGUSR1                      = 10
	SIGUSR2                      = 12
	SIGVTALRM                    = 26
	SIGWINCH                     = 28
	SIGXCPU                      = 24
	SIGXFSZ                      = 25
	SIG_BLOCK                    = 0
	SIG_SETMASK                  = 2
	SIG_UNBLOCK                  = 1
	SVE_MAGIC                    = 0x53564501
	SVE_NUM_PREGS                = 16
	SVE_NUM_ZREGS                = 32
	SVE_VL_MAX                   = 8192
	SVE_VL_MIN                   = 16
	SVE_VQ_BYTES                 = 16
	SVE_VQ_MAX                   = 512
	SVE_VQ_MIN                   = 1
	X_ASM_GENERIC_INT_LL64_H     = 0
	X_ASM_GENERIC_TYPES_H        = 0
	X_ATFILE_SOURCE              = 1
	X_BITS_BYTESWAP_H            = 1
	X_BITS_ENDIANNESS_H          = 1
	X_BITS_ENDIAN_H              = 1
	X_BITS_PTHREADTYPES_ARCH_H   = 1
	X_BITS_PTHREADTYPES_COMMON_H = 1
	X_BITS_SIGACTION_H           = 1
	X_BITS_SIGCONTEXT_H          = 1
	X_BITS_SIGEVENT_CONSTS_H     = 1
	X_BITS_SIGINFO_ARCH_H        = 1
	X_BITS_SIGINFO_CONSTS_H      = 1
	X_BITS_SIGNUM_GENERIC_H      = 1
	X_BITS_SIGNUM_H              = 1
	X_BITS_SIGSTACK_H            = 1
	X_BITS_SIGTHREAD_H           = 1
	X_BITS_SS_FLAGS_H            = 1
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
	X_LINUX_POSIX_TYPES_H        = 0
	X_LINUX_TYPES_H              = 0
	X_LP64                       = 1
	X_NSIG                       = 65
	X_POSIX_C_SOURCE             = 200809
	X_POSIX_SOURCE               = 1
	X_RWLOCK_INTERNAL_H          = 0
	X_SIGNAL_H                   = 0
	X_SIZET_                     = 0
	X_SIZE_T                     = 0
	X_SIZE_T_                    = 0
	X_SIZE_T_DECLARED            = 0
	X_SIZE_T_DEFINED             = 0
	X_SIZE_T_DEFINED_            = 0
	X_STDC_PREDEF_H              = 1
	X_STRUCT_TIMESPEC            = 1
	X_SYS_CDEFS_H                = 1
	X_SYS_PROCFS_H               = 1
	X_SYS_SELECT_H               = 1
	X_SYS_SIZE_T_H               = 0
	X_SYS_TIME_H                 = 1
	X_SYS_TYPES_H                = 1
	X_SYS_UCONTEXT_H             = 1
	X_SYS_USER_H                 = 1
	X_THREAD_MUTEX_INTERNAL_H    = 1
	X_THREAD_SHARED_TYPES_H      = 1
	X_T_SIZE                     = 0
	X_T_SIZE_                    = 0
	Linux                        = 1
	Unix                         = 1
)

const (
	SIGEV_SIGNAL = 0
	SIGEV_NONE   = 1
	SIGEV_THREAD = 2

	SIGEV_THREAD_ID = 4
)

const (
	SEGV_MAPERR  = 1
	SEGV_ACCERR  = 2
	SEGV_BNDERR  = 3
	SEGV_PKUERR  = 4
	SEGV_ACCADI  = 5
	SEGV_ADIDERR = 6
	SEGV_ADIPERR = 7
)

const (
	BUS_ADRALN    = 1
	BUS_ADRERR    = 2
	BUS_OBJERR    = 3
	BUS_MCEERR_AR = 4
	BUS_MCEERR_AO = 5
)

const (
	CLD_EXITED    = 1
	CLD_KILLED    = 2
	CLD_DUMPED    = 3
	CLD_TRAPPED   = 4
	CLD_STOPPED   = 5
	CLD_CONTINUED = 6
)

const (
	POLL_IN  = 1
	POLL_OUT = 2
	POLL_MSG = 3
	POLL_ERR = 4
	POLL_PRI = 5
	POLL_HUP = 6
)

const (
	SI_ASYNCNL  = -60
	SI_DETHREAD = -7

	SI_TKILL   = -6
	SI_SIGIO   = -5
	SI_ASYNCIO = -4
	SI_MESGQ   = -3
	SI_TIMER   = -2
	SI_QUEUE   = -1
	SI_USER    = 0
	SI_KERNEL  = 128
)

const (
	ILL_ILLOPC   = 1
	ILL_ILLOPN   = 2
	ILL_ILLADR   = 3
	ILL_ILLTRP   = 4
	ILL_PRVOPC   = 5
	ILL_PRVREG   = 6
	ILL_COPROC   = 7
	ILL_BADSTK   = 8
	ILL_BADIADDR = 9
)

const (
	FPE_INTDIV   = 1
	FPE_INTOVF   = 2
	FPE_FLTDIV   = 3
	FPE_FLTOVF   = 4
	FPE_FLTUND   = 5
	FPE_FLTRES   = 6
	FPE_FLTINV   = 7
	FPE_FLTSUB   = 8
	FPE_FLTUNK   = 14
	FPE_CONDTRAP = 15
)

const (
	SS_ONSTACK = 1
	SS_DISABLE = 2
)

const (
	ITIMER_REAL = 0

	ITIMER_VIRTUAL = 1

	ITIMER_PROF = 2
)

type Ptrdiff_t = int64

type Size_t = uint64

type Wchar_t = uint32

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

type X__loff_t = X__off64_t
type X__caddr_t = uintptr

type X__intptr_t = int64

type X__socklen_t = uint32

type X__sig_atomic_t = int32

type Sig_atomic_t = X__sig_atomic_t

type X__sigset_t = struct{ F__val [16]uint64 }

type Sigset_t = X__sigset_t

type Pid_t = X__pid_t
type Uid_t = X__uid_t

type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec X__syscall_slong_t
}

type Sigval = struct {
	F__ccgo_pad1 [0]uint64
	Fsival_int   int32
	F__ccgo_pad2 [4]byte
}

type X__sigval_t = Sigval

type Siginfo_t = struct {
	Fsi_signo  int32
	Fsi_errno  int32
	Fsi_code   int32
	F__pad0    int32
	F_sifields struct {
		F__ccgo_pad1 [0]uint64
		F_pad        [28]int32
	}
}

type Sigval_t = X__sigval_t

type Pthread_attr_t1 = struct {
	F__ccgo_pad1 [0]uint64
	F__size      [64]uint8
}

type Pthread_attr_t = Pthread_attr_t1

type Sigevent = struct {
	Fsigev_value  X__sigval_t
	Fsigev_signo  int32
	Fsigev_notify int32
	F_sigev_un    struct {
		F__ccgo_pad1 [0]uint64
		F_pad        [12]int32
	}
}

type Sigevent_t = Sigevent

type X__sighandler_t = uintptr

type Sig_t = X__sighandler_t

type Sigaction = struct {
	F__sigaction_handler struct{ Fsa_handler X__sighandler_t }
	Fsa_mask             X__sigset_t
	Fsa_flags            int32
	F__ccgo_pad1         [4]byte
	Fsa_restorer         uintptr
}

type X__s8 = int8
type X__u8 = uint8

type X__s16 = int16
type X__u16 = uint16

type X__s32 = int32
type X__u32 = uint32

type X__s64 = int64
type X__u64 = uint64

type X__kernel_fd_set = struct{ Ffds_bits [16]uint64 }

type X__kernel_sighandler_t = uintptr

type X__kernel_key_t = int32
type X__kernel_mqd_t = int32

type X__kernel_old_uid_t = uint16
type X__kernel_old_gid_t = uint16

type X__kernel_long_t = int64
type X__kernel_ulong_t = uint64

type X__kernel_ino_t = X__kernel_ulong_t

type X__kernel_mode_t = uint32

type X__kernel_pid_t = int32

type X__kernel_ipc_pid_t = int32

type X__kernel_uid_t = uint32
type X__kernel_gid_t = uint32

type X__kernel_suseconds_t = X__kernel_long_t

type X__kernel_daddr_t = int32

type X__kernel_uid32_t = uint32
type X__kernel_gid32_t = uint32

type X__kernel_old_dev_t = uint32

type X__kernel_size_t = X__kernel_ulong_t
type X__kernel_ssize_t = X__kernel_long_t
type X__kernel_ptrdiff_t = X__kernel_long_t

type X__kernel_fsid_t = struct{ Fval [2]int32 }

type X__kernel_off_t = X__kernel_long_t
type X__kernel_loff_t = int64
type X__kernel_old_time_t = X__kernel_long_t
type X__kernel_time_t = X__kernel_long_t
type X__kernel_time64_t = int64
type X__kernel_clock_t = X__kernel_long_t
type X__kernel_timer_t = int32
type X__kernel_clockid_t = int32
type X__kernel_caddr_t = uintptr
type X__kernel_uid16_t = uint16
type X__kernel_gid16_t = uint16

type X__le16 = X__u16
type X__be16 = X__u16
type X__le32 = X__u32
type X__be32 = X__u32
type X__le64 = X__u64
type X__be64 = X__u64

type X__sum16 = X__u16
type X__wsum = X__u32

type X__poll_t = uint32

type Sigcontext = struct {
	Ffault_address X__u64
	Fregs          [31]X__u64
	Fsp            X__u64
	Fpc            X__u64
	Fpstate        X__u64
	F__reserved    [4096]X__u8
}

type X_aarch64_ctx = struct {
	Fmagic X__u32
	Fsize  X__u32
}

type Fpsimd_context = struct {
	Fhead struct {
		Fmagic X__u32
		Fsize  X__u32
	}
	Ffpsr  X__u32
	Ffpcr  X__u32
	Fvregs [32]X__uint128_t
}

type Esr_context = struct {
	Fhead struct {
		Fmagic X__u32
		Fsize  X__u32
	}
	Fesr X__u64
}

type Extra_context = struct {
	Fhead struct {
		Fmagic X__u32
		Fsize  X__u32
	}
	Fdatap      X__u64
	Fsize       X__u32
	F__reserved [3]X__u32
}

type Sve_context = struct {
	Fhead struct {
		Fmagic X__u32
		Fsize  X__u32
	}
	Fvl         X__u16
	F__reserved [3]X__u16
}

type Stack_t = struct {
	Fss_sp       uintptr
	Fss_flags    int32
	F__ccgo_pad1 [4]byte
	Fss_size     Size_t
}

type Time_t = X__time_t

type Timeval = struct {
	Ftv_sec  X__time_t
	Ftv_usec X__suseconds_t
}

type Suseconds_t = X__suseconds_t

type X__fd_mask = int64

type Fd_set = struct{ F__fds_bits [16]X__fd_mask }

type Fd_mask = X__fd_mask

type Timezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Itimerval = struct {
	Fit_interval struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
	Fit_value struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
}

type X__itimer_which_t = int32

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

type Off_t = X__off64_t

type Id_t = X__id_t

type Ssize_t = X__ssize_t

type Daddr_t = X__daddr_t
type Caddr_t = X__caddr_t

type Key_t = X__key_t

type Clock_t = X__clock_t

type Clockid_t = X__clockid_t

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
	F__size      [8]uint8
}

type Pthread_condattr_t = struct {
	F__ccgo_pad1 [0]uint32
	F__size      [8]uint8
}

type Pthread_key_t = uint32

type Pthread_once_t = int32

type Pthread_mutex_t = struct {
	F__data      X__pthread_mutex_s
	F__ccgo_pad1 [8]byte
}

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
	F__size      [8]uint8
}

type User_regs_struct = struct {
	Fregs   [31]uint64
	Fsp     uint64
	Fpc     uint64
	Fpstate uint64
}

type User_fpsimd_struct = struct {
	Fvregs [32]X__uint128_t
	Ffpsr  uint32
	Ffpcr  uint32
}

type Elf_greg_t = X__uint64_t

type Elf_gregset_t = [34]Elf_greg_t

type Elf_fpregset_t = User_fpsimd_struct

type X__pr_uid_t = uint32
type X__pr_gid_t = uint32

type Elf_siginfo = struct {
	Fsi_signo int32
	Fsi_code  int32
	Fsi_errno int32
}

type Elf_prstatus = struct {
	Fpr_info struct {
		Fsi_signo int32
		Fsi_code  int32
		Fsi_errno int32
	}
	Fpr_cursig   int16
	F__ccgo_pad1 [2]byte
	Fpr_sigpend  uint64
	Fpr_sighold  uint64
	Fpr_pid      X__pid_t
	Fpr_ppid     X__pid_t
	Fpr_pgrp     X__pid_t
	Fpr_sid      X__pid_t
	Fpr_utime    struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
	Fpr_stime struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
	Fpr_cutime struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
	Fpr_cstime struct {
		Ftv_sec  X__time_t
		Ftv_usec X__suseconds_t
	}
	Fpr_reg      Elf_gregset_t
	Fpr_fpvalid  int32
	F__ccgo_pad2 [4]byte
}

type Elf_prpsinfo = struct {
	Fpr_state    uint8
	Fpr_sname    uint8
	Fpr_zomb     uint8
	Fpr_nice     uint8
	F__ccgo_pad1 [4]byte
	Fpr_flag     uint64
	Fpr_uid      X__pr_uid_t
	Fpr_gid      X__pr_gid_t
	Fpr_pid      int32
	Fpr_ppid     int32
	Fpr_pgrp     int32
	Fpr_sid      int32
	Fpr_fname    [16]uint8
	Fpr_psargs   [80]uint8
}

type Psaddr_t = uintptr

type X__prgregset_t = Elf_gregset_t
type X__prfpregset_t = Elf_fpregset_t

type Prgregset_t = X__prgregset_t
type Prfpregset_t = X__prfpregset_t

type Lwpid_t = X__pid_t

type Prstatus_t = Elf_prstatus
type Prpsinfo_t = Elf_prpsinfo

type Greg_t = Elf_greg_t

type Gregset_t = Elf_gregset_t

type Fpregset_t = Elf_fpregset_t

type Mcontext_t = struct {
	Ffault_address uint64
	Fregs          [31]uint64
	Fsp            uint64
	Fpc            uint64
	Fpstate        uint64
	F__reserved    [4096]uint8
}

type Ucontext_t1 = struct {
	Fuc_flags    uint64
	Fuc_link     uintptr
	Fuc_stack    Stack_t
	Fuc_sigmask  Sigset_t
	Fuc_mcontext Mcontext_t
}

type Ucontext_t = Ucontext_t1

type Sigstack = struct {
	Fss_sp       uintptr
	Fss_onstack  int32
	F__ccgo_pad1 [4]byte
}

var _ uint8
