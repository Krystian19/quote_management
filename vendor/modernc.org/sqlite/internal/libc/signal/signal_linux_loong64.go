// Code generated by 'ccgo signal/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o signal/signal_linux_loong64.go -pkgname signal', DO NOT EDIT.

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
	CONTEXT_INFO_ALIGN           = 16
	FPU_CTX_ALIGN                = 8
	FPU_CTX_MAGIC                = 0x46505501
	LARCH_NGREG                  = 32
	LARCH_REG_A0                 = 4
	LARCH_REG_NARGS              = 8
	LARCH_REG_RA                 = 1
	LARCH_REG_S0                 = 23
	LARCH_REG_S1                 = 24
	LARCH_REG_S2                 = 25
	LARCH_REG_SP                 = 3
	LASX_CTX_ALIGN               = 32
	LASX_CTX_MAGIC               = 0x41535801
	LSX_CTX_ALIGN                = 16
	LSX_CTX_MAGIC                = 0x53580001
	MINSIGSTKSZ                  = 4096
	NSIG                         = 65
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
	SC_ADDRERR_RD                = 1073741824
	SC_ADDRERR_WR                = 2147483648
	SC_USED_FP                   = 1
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
	X_ABILP64                    = 3
	X_ASM_GENERIC_INT_LL64_H     = 0
	X_ASM_GENERIC_TYPES_H        = 0
	X_ASM_SIGCONTEXT_H           = 0
	X_ATFILE_SOURCE              = 1
	X_BITS_ATOMIC_WIDE_COUNTER_H = 0
	X_BITS_ENDIANNESS_H          = 1
	X_BITS_ENDIAN_H              = 1
	X_BITS_PTHREADTYPES_ARCH_H   = 1
	X_BITS_PTHREADTYPES_COMMON_H = 1
	X_BITS_SIGACTION_H           = 1
	X_BITS_SIGCONTEXT_H          = 1
	X_BITS_SIGEVENT_CONSTS_H     = 1
	X_BITS_SIGINFO_ARCH_H        = 1
	X_BITS_SIGINFO_CONSTS_H      = 1
	X_BITS_SIGNUM_ARCH_H         = 1
	X_BITS_SIGNUM_GENERIC_H      = 1
	X_BITS_SIGSTACK_H            = 1
	X_BITS_SIGTHREAD_H           = 1
	X_BITS_SS_FLAGS_H            = 1
	X_BITS_TIME64_H              = 1
	X_BITS_TYPESIZES_H           = 1
	X_BITS_TYPES_H               = 1
	X_BSD_SIZE_T_                = 0
	X_BSD_SIZE_T_DEFINED_        = 0
	X_DEFAULT_SOURCE             = 1
	X_FEATURES_H                 = 1
	X_FILE_OFFSET_BITS           = 64
	X_GCC_SIZE_T                 = 0
	X_LINUX_POSIX_TYPES_H        = 0
	X_LINUX_STDDEF_H             = 0
	X_LINUX_TYPES_H              = 0
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
	X_NSIG                       = 65
	X_POSIX_C_SOURCE             = 200809
	X_POSIX_SOURCE               = 1
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
	X_SYS_SIZE_T_H               = 0
	X_SYS_UCONTEXT_H             = 1
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
	SEGV_MTEAERR = 8
	SEGV_MTESERR = 9
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

type Sig_atomic_t = int32

type X__sigset_t = struct{ F__val [16]uint64 }

type Sigset_t = X__sigset_t

type Pid_t = int32
type Uid_t = uint32

type Time_t = int64

type Timespec = struct {
	Ftv_sec  int64
	Ftv_nsec int64
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
	F__size      [56]int8
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

type Sig_t = uintptr

type Sigaction = struct {
	F__sigaction_handler struct{ Fsa_handler uintptr }
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

type X__kernel_long_t = int64
type X__kernel_ulong_t = uint64

type X__kernel_ino_t = uint64

type X__kernel_mode_t = uint32

type X__kernel_pid_t = int32

type X__kernel_ipc_pid_t = int32

type X__kernel_uid_t = uint32
type X__kernel_gid_t = uint32

type X__kernel_suseconds_t = int64

type X__kernel_daddr_t = int32

type X__kernel_uid32_t = uint32
type X__kernel_gid32_t = uint32

type X__kernel_old_uid_t = uint32
type X__kernel_old_gid_t = uint32

type X__kernel_old_dev_t = uint32

type X__kernel_size_t = uint64
type X__kernel_ssize_t = int64
type X__kernel_ptrdiff_t = int64

type X__kernel_fsid_t = struct{ Fval [2]int32 }

type X__kernel_off_t = int64
type X__kernel_loff_t = int64
type X__kernel_old_time_t = int64
type X__kernel_time_t = int64
type X__kernel_time64_t = int64
type X__kernel_clock_t = int64
type X__kernel_timer_t = int32
type X__kernel_clockid_t = int32
type X__kernel_caddr_t = uintptr
type X__kernel_uid16_t = uint16
type X__kernel_gid16_t = uint16

type X__le16 = uint16
type X__be16 = uint16
type X__le32 = uint32
type X__be32 = uint32
type X__le64 = uint64
type X__be64 = uint64

type X__sum16 = uint16
type X__wsum = uint32

type X__poll_t = uint32

type Sigcontext = struct {
	F__ccgo_pad1 [0]uint64
	Fsc_pc       uint64
	Fsc_regs     [32]uint64
	Fsc_flags    uint32
	F__ccgo_pad2 [4]byte
}

type Sctx_info = struct {
	Fmagic   uint32
	Fsize    uint32
	Fpadding uint64
}

type Fpu_context = struct {
	Fregs        [32]uint64
	Ffcc         uint64
	Ffcsr        uint32
	F__ccgo_pad1 [4]byte
}

type Lsx_context = struct {
	Fregs        [64]uint64
	Ffcc         uint64
	Ffcsr        uint32
	F__ccgo_pad1 [4]byte
}

type Lasx_context = struct {
	Fregs        [128]uint64
	Ffcc         uint64
	Ffcsr        uint32
	F__ccgo_pad1 [4]byte
}

type Stack_t = struct {
	Fss_sp       uintptr
	Fss_flags    int32
	F__ccgo_pad1 [4]byte
	Fss_size     uint64
}

type Greg_t = uint64

type Gregset_t = [32]uint64

type Mcontext_t1 = struct {
	F__ccgo_pad1 [0]uint64
	F__pc        uint64
	F__gregs     [32]uint64
	F__flags     uint32
	F__ccgo_pad2 [4]byte
}

type Mcontext_t = Mcontext_t1

type Ucontext_t1 = struct {
	F__uc_flags  uint64
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

var _ int8