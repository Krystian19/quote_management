// Code generated by 'ccgo sys/socket/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o sys/socket/socket_netbsd_arm.go -pkgname socket', DO NOT EDIT.

package socket

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
	AF_APPLETALK              = 16
	AF_ARP                    = 28
	AF_BLUETOOTH              = 31
	AF_CAN                    = 35
	AF_CCITT                  = 10
	AF_CHAOS                  = 5
	AF_CNT                    = 21
	AF_COIP                   = 20
	AF_DATAKIT                = 9
	AF_DECnet                 = 12
	AF_DLI                    = 13
	AF_E164                   = 26
	AF_ECMA                   = 8
	AF_ETHER                  = 36
	AF_HYLINK                 = 15
	AF_IEEE80211              = 32
	AF_IMPLINK                = 3
	AF_INET                   = 2
	AF_INET6                  = 24
	AF_IPX                    = 23
	AF_ISDN                   = 26
	AF_ISO                    = 7
	AF_LAT                    = 14
	AF_LINK                   = 18
	AF_LOCAL                  = 1
	AF_MAX                    = 37
	AF_MPLS                   = 33
	AF_NATM                   = 27
	AF_NS                     = 6
	AF_OROUTE                 = 17
	AF_OSI                    = 7
	AF_PUP                    = 4
	AF_ROUTE                  = 34
	AF_SNA                    = 11
	AF_UNIX                   = 1
	AF_UNSPEC                 = 0
	MSG_BCAST                 = 0x0100
	MSG_CMSG_CLOEXEC          = 0x0800
	MSG_CONTROLMBUF           = 0x2000000
	MSG_CTRUNC                = 0x0020
	MSG_DONTROUTE             = 0x0004
	MSG_DONTWAIT              = 0x0080
	MSG_EOR                   = 0x0008
	MSG_IOVUSRSPACE           = 0x4000000
	MSG_LENUSRSPACE           = 0x8000000
	MSG_MCAST                 = 0x0200
	MSG_NAMEMBUF              = 0x1000000
	MSG_NBIO                  = 0x1000
	MSG_NOSIGNAL              = 0x0400
	MSG_NOTIFICATION          = 0x4000
	MSG_OOB                   = 0x0001
	MSG_PEEK                  = 0x0002
	MSG_TRUNC                 = 0x0010
	MSG_USERFLAGS             = 0x0ffffff
	MSG_WAITALL               = 0x0040
	MSG_WAITFORONE            = 0x2000
	NET_RT_DUMP               = 1
	NET_RT_FLAGS              = 2
	NET_RT_IFLIST             = 6
	NET_RT_OIFLIST            = 5
	NET_RT_OOIFLIST           = 4
	NET_RT_OOOIFLIST          = 3
	PCB_ALL                   = 0
	PCB_SLOP                  = 20
	PF_APPLETALK              = 16
	PF_ARP                    = 28
	PF_BLUETOOTH              = 31
	PF_CAN                    = 35
	PF_CCITT                  = 10
	PF_CHAOS                  = 5
	PF_CNT                    = 21
	PF_COIP                   = 20
	PF_DATAKIT                = 9
	PF_DECnet                 = 12
	PF_DLI                    = 13
	PF_E164                   = 26
	PF_ECMA                   = 8
	PF_ETHER                  = 36
	PF_HYLINK                 = 15
	PF_IMPLINK                = 3
	PF_INET                   = 2
	PF_INET6                  = 24
	PF_IPX                    = 23
	PF_ISDN                   = 26
	PF_ISO                    = 7
	PF_KEY                    = 29
	PF_LAT                    = 14
	PF_LINK                   = 18
	PF_LOCAL                  = 1
	PF_MAX                    = 37
	PF_MPLS                   = 33
	PF_NATM                   = 27
	PF_NS                     = 6
	PF_OROUTE                 = 17
	PF_OSI                    = 7
	PF_PIP                    = 25
	PF_PUP                    = 4
	PF_ROUTE                  = 34
	PF_RTIP                   = 22
	PF_SNA                    = 11
	PF_UNIX                   = 1
	PF_UNSPEC                 = 0
	PF_XTP                    = 19
	SCM_CREDS                 = 0x10
	SCM_RIGHTS                = 0x01
	SCM_TIMESTAMP             = 0x08
	SHUT_RD                   = 0
	SHUT_RDWR                 = 2
	SHUT_WR                   = 1
	SOCK_CLOEXEC              = 0x10000000
	SOCK_CONN_DGRAM           = 6
	SOCK_DCCP                 = 6
	SOCK_DGRAM                = 2
	SOCK_FLAGS_MASK           = 0xf0000000
	SOCK_NONBLOCK             = 0x20000000
	SOCK_NOSIGPIPE            = 0x40000000
	SOCK_RAW                  = 3
	SOCK_RDM                  = 4
	SOCK_SEQPACKET            = 5
	SOCK_STREAM               = 1
	SOL_SOCKET                = 0xffff
	SOMAXCONN                 = 128
	SO_ACCEPTCONN             = 0x0002
	SO_ACCEPTFILTER           = 0x1000
	SO_BROADCAST              = 0x0020
	SO_DEBUG                  = 0x0001
	SO_DEFOPTS                = 27645
	SO_DONTROUTE              = 0x0010
	SO_ERROR                  = 0x1007
	SO_KEEPALIVE              = 0x0008
	SO_LINGER                 = 0x0080
	SO_NOHEADER               = 0x100a
	SO_NOSIGPIPE              = 0x0800
	SO_OOBINLINE              = 0x0100
	SO_OVERFLOWED             = 0x1009
	SO_RCVBUF                 = 0x1002
	SO_RCVLOWAT               = 0x1004
	SO_RCVTIMEO               = 0x100c
	SO_RERROR                 = 0x4000
	SO_REUSEADDR              = 0x0004
	SO_REUSEPORT              = 0x0200
	SO_SNDBUF                 = 0x1001
	SO_SNDLOWAT               = 0x1003
	SO_SNDTIMEO               = 0x100b
	SO_TIMESTAMP              = 0x2000
	SO_TYPE                   = 0x1008
	SO_USELOOPBACK            = 0x0040
	UIO_MAXIOV                = 1024
	X_ARM_ARCH_4T             = 0
	X_ARM_ARCH_5              = 0
	X_ARM_ARCH_5T             = 0
	X_ARM_ARCH_6              = 0
	X_ARM_ARCH_7              = 0
	X_ARM_ARCH_DWORD_OK       = 0
	X_ARM_ARCH_T2             = 0
	X_ARM_CDEFS_H_            = 0
	X_ARM_INT_TYPES_H_        = 0
	X_FILE_OFFSET_BITS        = 64
	X_NETBSD_SOURCE           = 1
	X_SS_MAXSIZE              = 128
	X_SYS_ANSI_H_             = 0
	X_SYS_CDEFS_ELF_H_        = 0
	X_SYS_CDEFS_H_            = 0
	X_SYS_COMMON_ANSI_H_      = 0
	X_SYS_COMMON_INT_TYPES_H_ = 0
	X_SYS_SIGTYPES_H_         = 0
	X_SYS_SOCKET_H_           = 0
	X_SYS_UIO_H_              = 0
	Pseudo_AF_HDRCMPLT        = 30
	Pseudo_AF_KEY             = 29
	Pseudo_AF_PIP             = 25
	Pseudo_AF_RTIP            = 22
	Pseudo_AF_XTP             = 19
)

const (
	UIO_READ  = 0
	UIO_WRITE = 1
)

const (
	UIO_USERSPACE = 0
	UIO_SYSSPACE  = 1
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

type X__intptr_t = int32
type X__uintptr_t = uint32

type X__caddr_t = uintptr
type X__gid_t = X__uint32_t
type X__in_addr_t = X__uint32_t
type X__in_port_t = X__uint16_t
type X__mode_t = X__uint32_t
type X__off_t = X__int64_t
type X__pid_t = X__int32_t
type X__sa_family_t = X__uint8_t
type X__socklen_t = uint32
type X__uid_t = X__uint32_t
type X__fsblkcnt_t = X__uint64_t
type X__fsfilcnt_t = X__uint64_t
type X__wctrans_t = uintptr
type X__wctype_t = uintptr

type X__mbstate_t = struct {
	F__mbstateL  X__int64_t
	F__ccgo_pad1 [120]byte
}

type X__va_list = X__builtin_va_list

type Sa_family_t = X__sa_family_t

type Socklen_t = X__socklen_t

type Ssize_t = int32

type Iovec = struct {
	Fiov_base uintptr
	Fiov_len  Size_t
}

type Off_t = X__off_t

type Sigset_t = struct{ F__bits [4]X__uint32_t }

type Sigaltstack = struct {
	Fss_sp    uintptr
	Fss_size  Size_t
	Fss_flags int32
}

type Stack_t = Sigaltstack

type Linger = struct {
	Fl_onoff  int32
	Fl_linger int32
}

type Accept_filter_arg = struct {
	Faf_name [16]uint8
	Faf_arg  [240]uint8
}

type Sockaddr = struct {
	Fsa_len    X__uint8_t
	Fsa_family X__sa_family_t
	Fsa_data   [14]uint8
}

type Sockaddr_storage = struct {
	Fss_len     X__uint8_t
	Fss_family  X__sa_family_t
	F__ss_pad1  [6]uint8
	F__ss_align X__int64_t
	F__ss_pad2  [112]uint8
}

type Pid_t = X__pid_t

type Gid_t = X__gid_t

type Uid_t = X__uid_t

type Sockcred = struct {
	Fsc_pid     X__pid_t
	Fsc_uid     X__uid_t
	Fsc_euid    X__uid_t
	Fsc_gid     X__gid_t
	Fsc_egid    X__gid_t
	Fsc_ngroups int32
	Fsc_groups  [1]X__gid_t
}

type Kinfo_pcb = struct {
	Fki_pcbaddr  X__uint64_t
	Fki_ppcbaddr X__uint64_t
	Fki_sockaddr X__uint64_t
	Fki_family   X__uint32_t
	Fki_type     X__uint32_t
	Fki_protocol X__uint32_t
	Fki_pflags   X__uint32_t
	Fki_sostate  X__uint32_t
	Fki_prstate  X__uint32_t
	Fki_tstate   X__int32_t
	Fki_tflags   X__uint32_t
	Fki_rcvq     X__uint64_t
	Fki_sndq     X__uint64_t
	Fki_s        struct {
		F_kis_src struct {
			Fsa_len    X__uint8_t
			Fsa_family X__sa_family_t
			Fsa_data   [14]uint8
		}
		F__ccgo_pad1 [248]byte
	}
	Fki_d struct {
		F_kid_dst struct {
			Fsa_len    X__uint8_t
			Fsa_family X__sa_family_t
			Fsa_data   [14]uint8
		}
		F__ccgo_pad1 [248]byte
	}
	Fki_inode   X__uint64_t
	Fki_vnode   X__uint64_t
	Fki_conn    X__uint64_t
	Fki_refs    X__uint64_t
	Fki_nextref X__uint64_t
}

type Msghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    X__socklen_t
	Fmsg_iov        uintptr
	Fmsg_iovlen     int32
	Fmsg_control    uintptr
	Fmsg_controllen X__socklen_t
	Fmsg_flags      int32
}

type Mmsghdr = struct {
	Fmsg_hdr struct {
		Fmsg_name       uintptr
		Fmsg_namelen    X__socklen_t
		Fmsg_iov        uintptr
		Fmsg_iovlen     int32
		Fmsg_control    uintptr
		Fmsg_controllen X__socklen_t
		Fmsg_flags      int32
	}
	Fmsg_len uint32
}

type Cmsghdr = struct {
	Fcmsg_len   X__socklen_t
	Fcmsg_level int32
	Fcmsg_type  int32
}

var _ uint8