// Code generated by 'ccgo termios/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o termios/termios_darwin_arm64.go -pkgname termios', DO NOT EDIT.

package termios

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
	ALTWERASE                              = 0x00000200
	B0                                     = 0
	B110                                   = 110
	B115200                                = 115200
	B1200                                  = 1200
	B134                                   = 134
	B14400                                 = 14400
	B150                                   = 150
	B1800                                  = 1800
	B19200                                 = 19200
	B200                                   = 200
	B230400                                = 230400
	B2400                                  = 2400
	B28800                                 = 28800
	B300                                   = 300
	B38400                                 = 38400
	B4800                                  = 4800
	B50                                    = 50
	B57600                                 = 57600
	B600                                   = 600
	B7200                                  = 7200
	B75                                    = 75
	B76800                                 = 76800
	B9600                                  = 9600
	BRKINT                                 = 0x00000002
	BS0                                    = 0x00000000
	BS1                                    = 0x00008000
	BSDLY                                  = 0x00008000
	CBRK                                   = 255
	CCAR_OFLOW                             = 0x00100000
	CCTS_OFLOW                             = 0x00010000
	CDISCARD                               = 15
	CDSR_OFLOW                             = 0x00080000
	CDSUSP                                 = 25
	CDTR_IFLOW                             = 0x00040000
	CEOF                                   = 4
	CEOL                                   = 0xff
	CEOT                                   = 4
	CERASE                                 = 0177
	CFLUSH                                 = 15
	CIGNORE                                = 0x00000001
	CINTR                                  = 3
	CKILL                                  = 21
	CLNEXT                                 = 22
	CLOCAL                                 = 0x00008000
	CMIN                                   = 1
	CQUIT                                  = 034
	CR0                                    = 0x00000000
	CR1                                    = 0x00001000
	CR2                                    = 0x00002000
	CR3                                    = 0x00003000
	CRDLY                                  = 0x00003000
	CREAD                                  = 0x00000800
	CREPRINT                               = 18
	CRPRNT                                 = 18
	CRTSCTS                                = 196608
	CRTS_IFLOW                             = 0x00020000
	CS5                                    = 0x00000000
	CS6                                    = 0x00000100
	CS7                                    = 0x00000200
	CS8                                    = 0x00000300
	CSIZE                                  = 0x00000300
	CSTART                                 = 17
	CSTATUS                                = 20
	CSTOP                                  = 19
	CSTOPB                                 = 0x00000400
	CSUSP                                  = 26
	CTIME                                  = 0
	CWERASE                                = 23
	ECHO                                   = 0x00000008
	ECHOCTL                                = 0x00000040
	ECHOE                                  = 0x00000002
	ECHOK                                  = 0x00000004
	ECHOKE                                 = 0x00000001
	ECHONL                                 = 0x00000010
	ECHOPRT                                = 0x00000020
	EXTA                                   = 19200
	EXTB                                   = 38400
	EXTPROC                                = 0x00000800
	FF0                                    = 0x00000000
	FF1                                    = 0x00004000
	FFDLY                                  = 0x00004000
	FLUSHO                                 = 0x00800000
	HUPCL                                  = 0x00004000
	ICANON                                 = 0x00000100
	ICRNL                                  = 0x00000100
	IEXTEN                                 = 0x00000400
	IGNBRK                                 = 0x00000001
	IGNCR                                  = 0x00000080
	IGNPAR                                 = 0x00000004
	IMAXBEL                                = 0x00002000
	INLCR                                  = 0x00000040
	INPCK                                  = 0x00000010
	IOCPARM_MASK                           = 0x1fff
	IOCPARM_MAX                            = 8192
	ISIG                                   = 0x00000080
	ISTRIP                                 = 0x00000020
	IUTF8                                  = 0x00004000
	IXANY                                  = 0x00000800
	IXOFF                                  = 0x00000400
	IXON                                   = 0x00000200
	MDMBUF                                 = 0x00100000
	NCCS                                   = 20
	NL0                                    = 0x00000000
	NL1                                    = 0x00000100
	NL2                                    = 0x00000200
	NL3                                    = 0x00000300
	NLDLY                                  = 0x00000300
	NOFLSH                                 = 0x80000000
	NOKERNINFO                             = 0x02000000
	OCRNL                                  = 0x00000010
	OFDEL                                  = 0x00020000
	OFILL                                  = 0x00000080
	ONLCR                                  = 0x00000002
	ONLRET                                 = 0x00000040
	ONOCR                                  = 0x00000020
	ONOEOT                                 = 0x00000008
	OPOST                                  = 0x00000001
	OXTABS                                 = 0x00000004
	PARENB                                 = 0x00001000
	PARMRK                                 = 0x00000008
	PARODD                                 = 0x00002000
	PENDIN                                 = 0x20000000
	PPPDISC                                = 5
	SLIPDISC                               = 4
	TAB0                                   = 0x00000000
	TAB1                                   = 0x00000400
	TAB2                                   = 0x00000800
	TAB3                                   = 0x00000004
	TABDLY                                 = 0x00000c04
	TABLDISC                               = 3
	TCIFLUSH                               = 1
	TCIOFF                                 = 3
	TCIOFLUSH                              = 3
	TCION                                  = 4
	TCOFLUSH                               = 2
	TCOOFF                                 = 1
	TCOON                                  = 2
	TCSADRAIN                              = 1
	TCSAFLUSH                              = 2
	TCSANOW                                = 0
	TCSASOFT                               = 0x10
	TIOCM_CAR                              = 0100
	TIOCM_CD                               = 64
	TIOCM_CTS                              = 0040
	TIOCM_DSR                              = 0400
	TIOCM_DTR                              = 0002
	TIOCM_LE                               = 0001
	TIOCM_RI                               = 128
	TIOCM_RNG                              = 0200
	TIOCM_RTS                              = 0004
	TIOCM_SR                               = 0020
	TIOCM_ST                               = 0010
	TIOCPKT_DATA                           = 0x00
	TIOCPKT_DOSTOP                         = 0x20
	TIOCPKT_FLUSHREAD                      = 0x01
	TIOCPKT_FLUSHWRITE                     = 0x02
	TIOCPKT_IOCTL                          = 0x40
	TIOCPKT_NOSTOP                         = 0x10
	TIOCPKT_START                          = 0x08
	TIOCPKT_STOP                           = 0x04
	TOSTOP                                 = 0x00400000
	TTYDEF_CFLAG                           = 19200
	TTYDEF_IFLAG                           = 11010
	TTYDEF_LFLAG                           = 1483
	TTYDEF_OFLAG                           = 3
	TTYDEF_SPEED                           = 9600
	TTYDISC                                = 0
	VDISCARD                               = 15
	VDSUSP                                 = 11
	VEOF                                   = 0
	VEOL                                   = 1
	VEOL2                                  = 2
	VERASE                                 = 3
	VINTR                                  = 8
	VKILL                                  = 5
	VLNEXT                                 = 14
	VMIN                                   = 16
	VQUIT                                  = 9
	VREPRINT                               = 6
	VSTART                                 = 12
	VSTATUS                                = 18
	VSTOP                                  = 13
	VSUSP                                  = 10
	VT0                                    = 0x00000000
	VT1                                    = 0x00010000
	VTDLY                                  = 0x00010000
	VTIME                                  = 17
	VWERASE                                = 4
	X_BSD_ARM__TYPES_H_                    = 0
	X_BSD_MACHINE__TYPES_H_                = 0
	X_CDEFS_H_                             = 0
	X_DARWIN_FEATURE_64_BIT_INODE          = 1
	X_DARWIN_FEATURE_ONLY_64_BIT_INODE     = 1
	X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
	X_DARWIN_FEATURE_ONLY_VERS_1050        = 1
	X_DARWIN_FEATURE_UNIX_CONFORMANCE      = 3
	X_FILE_OFFSET_BITS                     = 64
	X_FORTIFY_SOURCE                       = 2
	X_LP64                                 = 1
	X_Nonnull                              = 0
	X_Null_unspecified                     = 0
	X_Nullable                             = 0
	X_PID_T                                = 0
	X_SYS_IOCCOM_H_                        = 0
	X_SYS_TERMIOS_H_                       = 0
	X_SYS_TTYCOM_H_                        = 0
	X_SYS_TTYDEFAULTS_H_                   = 0
	X_SYS__PTHREAD_TYPES_H_                = 0
	X_SYS__TYPES_H_                        = 0
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

var X__darwin_check_fd_set_overflow uintptr

type Tcflag_t = uint64
type Cc_t = uint8
type Speed_t = uint64

type Termios = struct {
	Fc_iflag     Tcflag_t
	Fc_oflag     Tcflag_t
	Fc_cflag     Tcflag_t
	Fc_lflag     Tcflag_t
	Fc_cc        [20]Cc_t
	F__ccgo_pad1 [4]byte
	Fc_ispeed    Speed_t
	Fc_ospeed    Speed_t
}

type X__int8_t = int8
type X__uint8_t = uint8
type X__int16_t = int16
type X__uint16_t = uint16
type X__int32_t = int32
type X__uint32_t = uint32
type X__int64_t = int64
type X__uint64_t = uint64

type X__darwin_intptr_t = int64
type X__darwin_natural_t = uint32

type X__darwin_ct_rune_t = int32

type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
}

type X__darwin_mbstate_t = X__mbstate_t

type X__darwin_ptrdiff_t = int64

type X__darwin_size_t = uint64

type X__darwin_va_list = X__builtin_va_list

type X__darwin_wchar_t = int32

type X__darwin_rune_t = X__darwin_wchar_t

type X__darwin_wint_t = int32

type X__darwin_clock_t = uint64
type X__darwin_socklen_t = X__uint32_t
type X__darwin_ssize_t = int64
type X__darwin_time_t = int64

type X__darwin_blkcnt_t = X__int64_t
type X__darwin_blksize_t = X__int32_t
type X__darwin_dev_t = X__int32_t
type X__darwin_fsblkcnt_t = uint32
type X__darwin_fsfilcnt_t = uint32
type X__darwin_gid_t = X__uint32_t
type X__darwin_id_t = X__uint32_t
type X__darwin_ino64_t = X__uint64_t
type X__darwin_ino_t = X__darwin_ino64_t
type X__darwin_mach_port_name_t = X__darwin_natural_t
type X__darwin_mach_port_t = X__darwin_mach_port_name_t
type X__darwin_mode_t = X__uint16_t
type X__darwin_off_t = X__int64_t
type X__darwin_pid_t = X__int32_t
type X__darwin_sigset_t = X__uint32_t
type X__darwin_suseconds_t = X__int32_t
type X__darwin_uid_t = X__uint32_t
type X__darwin_useconds_t = X__uint32_t
type X__darwin_uuid_t = [16]uint8
type X__darwin_uuid_string_t = [37]int8

type X__darwin_pthread_handler_rec = struct {
	F__routine uintptr
	F__arg     uintptr
	F__next    uintptr
}

type X_opaque_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type X_opaque_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type X_opaque_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type X_opaque_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type X_opaque_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type X_opaque_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type X_opaque_pthread_t = struct {
	F__sig           int64
	F__cleanup_stack uintptr
	F__opaque        [8176]int8
}

type X__darwin_pthread_attr_t = X_opaque_pthread_attr_t
type X__darwin_pthread_cond_t = X_opaque_pthread_cond_t
type X__darwin_pthread_condattr_t = X_opaque_pthread_condattr_t
type X__darwin_pthread_key_t = uint64
type X__darwin_pthread_mutex_t = X_opaque_pthread_mutex_t
type X__darwin_pthread_mutexattr_t = X_opaque_pthread_mutexattr_t
type X__darwin_pthread_once_t = X_opaque_pthread_once_t
type X__darwin_pthread_rwlock_t = X_opaque_pthread_rwlock_t
type X__darwin_pthread_rwlockattr_t = X_opaque_pthread_rwlockattr_t
type X__darwin_pthread_t = uintptr

type Winsize = struct {
	Fws_row    uint16
	Fws_col    uint16
	Fws_xpixel uint16
	Fws_ypixel uint16
}

type X__darwin_nl_item = int32
type X__darwin_wctrans_t = int32
type X__darwin_wctype_t = X__uint32_t

type Pid_t = X__darwin_pid_t

var _ int8