// Code generated by 'ccgo signal/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o signal/signal_openbsd_386.go -pkgname signal', DO NOT EDIT.

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
	BIG_ENDIAN               = 4321
	BUS_ADRALN               = 1
	BUS_ADRERR               = 2
	BUS_OBJERR               = 3
	BYTE_ORDER               = 1234
	CLD_CONTINUED            = 6
	CLD_DUMPED               = 3
	CLD_EXITED               = 1
	CLD_KILLED               = 2
	CLD_STOPPED              = 5
	CLD_TRAPPED              = 4
	CLK_TCK                  = 100
	CLOCKS_PER_SEC           = 100
	CLOCK_BOOTTIME           = 6
	CLOCK_MONOTONIC          = 3
	CLOCK_PROCESS_CPUTIME_ID = 2
	CLOCK_REALTIME           = 0
	CLOCK_THREAD_CPUTIME_ID  = 4
	CLOCK_UPTIME             = 5
	DST_AUST                 = 2
	DST_CAN                  = 6
	DST_EET                  = 5
	DST_MET                  = 4
	DST_NONE                 = 0
	DST_USA                  = 1
	DST_WET                  = 3
	EMT_TAGOVF               = 1
	FD_SETSIZE               = 1024
	FPE_FLTDIV               = 3
	FPE_FLTINV               = 7
	FPE_FLTOVF               = 4
	FPE_FLTRES               = 6
	FPE_FLTSUB               = 8
	FPE_FLTUND               = 5
	FPE_INTDIV               = 1
	FPE_INTOVF               = 2
	ILL_BADSTK               = 8
	ILL_COPROC               = 7
	ILL_ILLADR               = 3
	ILL_ILLOPC               = 1
	ILL_ILLOPN               = 2
	ILL_ILLTRP               = 4
	ILL_PRVOPC               = 5
	ILL_PRVREG               = 6
	ITIMER_PROF              = 2
	ITIMER_REAL              = 0
	ITIMER_VIRTUAL           = 1
	LITTLE_ENDIAN            = 1234
	MINSIGSTKSZ              = 12288
	NBBY                     = 8
	NSIG                     = 33
	NSIGBUS                  = 3
	NSIGCLD                  = 6
	NSIGEMT                  = 1
	NSIGFPE                  = 8
	NSIGILL                  = 8
	NSIGSEGV                 = 2
	NSIGTRAP                 = 2
	PDP_ENDIAN               = 3412
	SA_NOCLDSTOP             = 0x0008
	SA_NOCLDWAIT             = 0x0020
	SA_NODEFER               = 0x0010
	SA_ONSTACK               = 0x0001
	SA_RESETHAND             = 0x0004
	SA_RESTART               = 0x0002
	SA_SIGINFO               = 0x0040
	SEGV_ACCERR              = 2
	SEGV_MAPERR              = 1
	SIGABRT                  = 6
	SIGALRM                  = 14
	SIGBUS                   = 10
	SIGCHLD                  = 20
	SIGCONT                  = 19
	SIGEMT                   = 7
	SIGFPE                   = 8
	SIGHUP                   = 1
	SIGILL                   = 4
	SIGINFO                  = 29
	SIGINT                   = 2
	SIGIO                    = 23
	SIGIOT                   = 6
	SIGKILL                  = 9
	SIGPIPE                  = 13
	SIGPROF                  = 27
	SIGQUIT                  = 3
	SIGSEGV                  = 11
	SIGSTKSZ                 = 28672
	SIGSTOP                  = 17
	SIGSYS                   = 12
	SIGTERM                  = 15
	SIGTHR                   = 32
	SIGTRAP                  = 5
	SIGTSTP                  = 18
	SIGTTIN                  = 21
	SIGTTOU                  = 22
	SIGURG                   = 16
	SIGUSR1                  = 30
	SIGUSR2                  = 31
	SIGVTALRM                = 26
	SIGWINCH                 = 28
	SIGXCPU                  = 24
	SIGXFSZ                  = 25
	SIG_BLOCK                = 1
	SIG_SETMASK              = 3
	SIG_UNBLOCK              = 2
	SI_LWP                   = -1
	SI_MAXSZ                 = 128
	SI_NOINFO                = 32767
	SI_QUEUE                 = -2
	SI_TIMER                 = -3
	SI_USER                  = 0
	SS_DISABLE               = 0x0004
	SS_ONSTACK               = 0x0001
	SV_INTERRUPT             = 2
	SV_ONSTACK               = 1
	SV_RESETHAND             = 4
	TIMER_ABSTIME            = 0x1
	TIMER_RELTIME            = 0x0
	TIME_UTC                 = 1
	TRAP_BRKPT               = 1
	TRAP_TRACE               = 2
	X_BIG_ENDIAN             = 4321
	X_BYTE_ORDER             = 1234
	X_CLOCKID_T_DEFINED_     = 0
	X_CLOCK_T_DEFINED_       = 0
	X_FILE_OFFSET_BITS       = 64
	X_ILP32                  = 1
	X_INT16_T_DEFINED_       = 0
	X_INT32_T_DEFINED_       = 0
	X_INT64_T_DEFINED_       = 0
	X_INT8_T_DEFINED_        = 0
	X_LITTLE_ENDIAN          = 1234
	X_LOCALE_T_DEFINED_      = 0
	X_MACHINE_CDEFS_H_       = 0
	X_MACHINE_ENDIAN_H_      = 0
	X_MACHINE_SIGNAL_H_      = 0
	X_MACHINE__TYPES_H_      = 0
	X_MAX_PAGE_SHIFT         = 12
	X_NSIG                   = 33
	X_OFF_T_DEFINED_         = 0
	X_PDP_ENDIAN             = 3412
	X_PID_T_DEFINED_         = 0
	X_QUAD_HIGHWORD          = 1
	X_QUAD_LOWWORD           = 0
	X_SELECT_DEFINED_        = 0
	X_SIGSET_T_DEFINED_      = 0
	X_SIZE_T_DEFINED_        = 0
	X_SSIZE_T_DEFINED_       = 0
	X_STACKALIGNBYTES        = 15
	X_SYS_CDEFS_H_           = 0
	X_SYS_ENDIAN_H_          = 0
	X_SYS_SELECT_H_          = 0
	X_SYS_SIGINFO_H          = 0
	X_SYS_SIGNAL_H_          = 0
	X_SYS_TIME_H_            = 0
	X_SYS_TYPES_H_           = 0
	X_SYS__ENDIAN_H_         = 0
	X_SYS__TIME_H_           = 0
	X_SYS__TYPES_H_          = 0
	X_TIMER_T_DEFINED_       = 0
	X_TIMESPEC_DECLARED      = 0
	X_TIMEVAL_DECLARED       = 0
	X_TIME_H_                = 0
	X_TIME_T_DEFINED_        = 0
	X_UINT16_T_DEFINED_      = 0
	X_UINT32_T_DEFINED_      = 0
	X_UINT64_T_DEFINED_      = 0
	X_UINT8_T_DEFINED_       = 0
	X_USER_SIGNAL_H          = 0
	I386                     = 1
	Unix                     = 1
)

type Ptrdiff_t = int32

type Size_t = uint32

type Wchar_t = int32

type X__builtin_va_list = uintptr
type X__float128 = float64

type Sig_atomic_t = int32

type Sigcontext = struct {
	Fsc_gs      int32
	Fsc_fs      int32
	Fsc_es      int32
	Fsc_ds      int32
	Fsc_edi     int32
	Fsc_esi     int32
	Fsc_ebp     int32
	Fsc_ebx     int32
	Fsc_edx     int32
	Fsc_ecx     int32
	Fsc_eax     int32
	Fsc_eip     int32
	Fsc_cs      int32
	Fsc_eflags  int32
	Fsc_esp     int32
	Fsc_ss      int32
	Fsc_cookie  int32
	Fsc_mask    int32
	Fsc_trapno  int32
	Fsc_err     int32
	Fsc_fpstate uintptr
}

type Sigset_t = uint32

type Sigval = struct{ Fsival_int int32 }

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

type X__int_fast8_t = X__int32_t
type X__uint_fast8_t = X__uint32_t
type X__int_fast16_t = X__int32_t
type X__uint_fast16_t = X__uint32_t
type X__int_fast32_t = X__int32_t
type X__uint_fast32_t = X__uint32_t
type X__int_fast64_t = X__int64_t
type X__uint_fast64_t = X__uint64_t

type X__intptr_t = int32
type X__uintptr_t = uint32

type X__intmax_t = X__int64_t
type X__uintmax_t = X__uint64_t

type X__register_t = int32

type X__vaddr_t = uint32
type X__paddr_t = uint32
type X__vsize_t = uint32
type X__psize_t = uint32

type X__double_t = float64
type X__float_t = float64
type X__ptrdiff_t = int32
type X__size_t = uint32
type X__ssize_t = int32
type X__va_list = X__builtin_va_list

type X__wchar_t = int32
type X__wint_t = int32
type X__rune_t = int32
type X__wctrans_t = uintptr
type X__wctype_t = uintptr

type X__blkcnt_t = X__int64_t
type X__blksize_t = X__int32_t
type X__clock_t = X__int64_t
type X__clockid_t = X__int32_t
type X__cpuid_t = uint32
type X__dev_t = X__int32_t
type X__fixpt_t = X__uint32_t
type X__fsblkcnt_t = X__uint64_t
type X__fsfilcnt_t = X__uint64_t
type X__gid_t = X__uint32_t
type X__id_t = X__uint32_t
type X__in_addr_t = X__uint32_t
type X__in_port_t = X__uint16_t
type X__ino_t = X__uint64_t
type X__key_t = int32
type X__mode_t = X__uint32_t
type X__nlink_t = X__uint32_t
type X__off_t = X__int64_t
type X__pid_t = X__int32_t
type X__rlim_t = X__uint64_t
type X__sa_family_t = X__uint8_t
type X__segsz_t = X__int32_t
type X__socklen_t = X__uint32_t
type X__suseconds_t = int32
type X__time_t = X__int64_t
type X__timer_t = X__int32_t
type X__uid_t = X__uint32_t
type X__useconds_t = X__uint32_t

type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint32
	F__mbstate8  [128]int8
}

type U_char = uint8
type U_short = uint16
type U_int = uint32
type U_long = uint32

type Unchar = uint8
type Ushort = uint16
type Uint = uint32
type Ulong = uint32

type Cpuid_t = X__cpuid_t
type Register_t = X__register_t

type Int8_t = X__int8_t

type Uint8_t = X__uint8_t

type Int16_t = X__int16_t

type Uint16_t = X__uint16_t

type Int32_t = X__int32_t

type Uint32_t = X__uint32_t

type Int64_t = X__int64_t

type Uint64_t = X__uint64_t

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

type Quad_t = X__int64_t
type U_quad_t = X__uint64_t

type Vaddr_t = X__vaddr_t
type Paddr_t = X__paddr_t
type Vsize_t = X__vsize_t
type Psize_t = X__psize_t

type Blkcnt_t = X__blkcnt_t
type Blksize_t = X__blksize_t
type Caddr_t = uintptr
type Daddr32_t = X__int32_t
type Daddr_t = X__int64_t
type Dev_t = X__dev_t
type Fixpt_t = X__fixpt_t
type Gid_t = X__gid_t
type Id_t = X__id_t
type Ino_t = X__ino_t
type Key_t = X__key_t
type Mode_t = X__mode_t
type Nlink_t = X__nlink_t
type Rlim_t = X__rlim_t
type Segsz_t = X__segsz_t
type Uid_t = X__uid_t
type Useconds_t = X__useconds_t
type Suseconds_t = X__suseconds_t
type Fsblkcnt_t = X__fsblkcnt_t
type Fsfilcnt_t = X__fsfilcnt_t

type Clock_t = X__clock_t

type Clockid_t = X__clockid_t

type Pid_t = X__pid_t

type Ssize_t = X__ssize_t

type Time_t = X__time_t

type Timer_t = X__timer_t

type Off_t = X__off_t

type Timeval = struct {
	Ftv_sec  Time_t
	Ftv_usec Suseconds_t
}

type Timespec = struct {
	Ftv_sec  Time_t
	Ftv_nsec int32
}

type X__fd_mask = Uint32_t

type Fd_set1 = struct{ Ffds_bits [32]X__fd_mask }

type Fd_set = Fd_set1

type Timezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Itimerval = struct {
	Fit_interval struct {
		Ftv_sec  Time_t
		Ftv_usec Suseconds_t
	}
	Fit_value struct {
		Ftv_sec  Time_t
		Ftv_usec Suseconds_t
	}
}

type Clockinfo = struct {
	Fhz     int32
	Ftick   int32
	Fstathz int32
	Fprofhz int32
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

type Locale_t = uintptr

type Tm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	Ftm_gmtoff int32
	Ftm_zone   uintptr
}

type Siginfo_t = struct {
	Fsi_signo int32
	Fsi_code  int32
	Fsi_errno int32
	F_data    struct{ F_pad [29]int32 }
}

type Sigaction = struct {
	F__sigaction_u struct{ F__sa_handler uintptr }
	Fsa_mask       Sigset_t
	Fsa_flags      int32
}

type Sig_t = uintptr

type Sigvec = struct {
	Fsv_handler uintptr
	Fsv_mask    int32
	Fsv_flags   int32
}

type Sigaltstack = struct {
	Fss_sp    uintptr
	Fss_size  Size_t
	Fss_flags int32
}

type Stack_t = Sigaltstack

type Ucontext_t = Sigcontext

var _ int8