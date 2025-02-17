// Code generated by 'ccgo netinet/in/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o netinet/in/in_freebsd_386.go -pkgname in', DO NOT EDIT.

package in

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
	BIG_ENDIAN                    = 4321
	BYTE_ORDER                    = 1234
	ICMP6_FILTER                  = 18
	ICMPV6CTL_ND6_ONLINKNSRFC4861 = 47
	INET6_ADDRSTRLEN              = 46
	INET_ADDRSTRLEN               = 16
	IN_CLASSA_HOST                = 0x00ffffff
	IN_CLASSA_MAX                 = 128
	IN_CLASSA_NET                 = 0xff000000
	IN_CLASSA_NSHIFT              = 24
	IN_CLASSB_HOST                = 0x0000ffff
	IN_CLASSB_MAX                 = 65536
	IN_CLASSB_NET                 = 0xffff0000
	IN_CLASSB_NSHIFT              = 16
	IN_CLASSC_HOST                = 0x000000ff
	IN_CLASSC_NET                 = 0xffffff00
	IN_CLASSC_NSHIFT              = 8
	IN_CLASSD_HOST                = 0x0fffffff
	IN_CLASSD_NET                 = 0xf0000000
	IN_CLASSD_NSHIFT              = 28
	IN_HISTORICAL_NETS            = 0
	IN_LOOPBACKNET                = 127
	IN_NETMASK_DEFAULT            = 0xffffff00
	IPCTL_ACCEPTSOURCEROUTE       = 13
	IPCTL_DEFTTL                  = 3
	IPCTL_DIRECTEDBROADCAST       = 9
	IPCTL_FASTFORWARDING          = 14
	IPCTL_FORWARDING              = 1
	IPCTL_GIF_TTL                 = 16
	IPCTL_INTRDQDROPS             = 18
	IPCTL_INTRDQMAXLEN            = 17
	IPCTL_INTRQDROPS              = 11
	IPCTL_INTRQMAXLEN             = 10
	IPCTL_SENDREDIRECTS           = 2
	IPCTL_SOURCEROUTE             = 8
	IPCTL_STATS                   = 12
	IPPORT_EPHEMERALFIRST         = 10000
	IPPORT_EPHEMERALLAST          = 65535
	IPPORT_HIFIRSTAUTO            = 49152
	IPPORT_HILASTAUTO             = 65535
	IPPORT_MAX                    = 65535
	IPPORT_RESERVED               = 1024
	IPPORT_RESERVEDSTART          = 600
	IPPROTO_3PC                   = 34
	IPPROTO_ADFS                  = 68
	IPPROTO_AH                    = 51
	IPPROTO_AHIP                  = 61
	IPPROTO_APES                  = 99
	IPPROTO_ARGUS                 = 13
	IPPROTO_AX25                  = 93
	IPPROTO_BHA                   = 49
	IPPROTO_BLT                   = 30
	IPPROTO_BRSATMON              = 76
	IPPROTO_CARP                  = 112
	IPPROTO_CFTP                  = 62
	IPPROTO_CHAOS                 = 16
	IPPROTO_CMTP                  = 38
	IPPROTO_CPHB                  = 73
	IPPROTO_CPNX                  = 72
	IPPROTO_DCCP                  = 33
	IPPROTO_DDP                   = 37
	IPPROTO_DGP                   = 86
	IPPROTO_DIVERT                = 258
	IPPROTO_DONE                  = 257
	IPPROTO_DSTOPTS               = 60
	IPPROTO_EGP                   = 8
	IPPROTO_EMCON                 = 14
	IPPROTO_ENCAP                 = 98
	IPPROTO_EON                   = 80
	IPPROTO_ESP                   = 50
	IPPROTO_ETHERIP               = 97
	IPPROTO_FRAGMENT              = 44
	IPPROTO_GGP                   = 3
	IPPROTO_GMTP                  = 100
	IPPROTO_GRE                   = 47
	IPPROTO_HELLO                 = 63
	IPPROTO_HIP                   = 139
	IPPROTO_HMP                   = 20
	IPPROTO_HOPOPTS               = 0
	IPPROTO_ICMP                  = 1
	IPPROTO_ICMPV6                = 58
	IPPROTO_IDP                   = 22
	IPPROTO_IDPR                  = 35
	IPPROTO_IDRP                  = 45
	IPPROTO_IGMP                  = 2
	IPPROTO_IGP                   = 85
	IPPROTO_IGRP                  = 88
	IPPROTO_IL                    = 40
	IPPROTO_INLSP                 = 52
	IPPROTO_INP                   = 32
	IPPROTO_IP                    = 0
	IPPROTO_IPCOMP                = 108
	IPPROTO_IPCV                  = 71
	IPPROTO_IPEIP                 = 94
	IPPROTO_IPIP                  = 4
	IPPROTO_IPPC                  = 67
	IPPROTO_IPV4                  = 4
	IPPROTO_IPV6                  = 41
	IPPROTO_IRTP                  = 28
	IPPROTO_KRYPTOLAN             = 65
	IPPROTO_LARP                  = 91
	IPPROTO_LEAF1                 = 25
	IPPROTO_LEAF2                 = 26
	IPPROTO_MAX                   = 256
	IPPROTO_MEAS                  = 19
	IPPROTO_MH                    = 135
	IPPROTO_MHRP                  = 48
	IPPROTO_MICP                  = 95
	IPPROTO_MOBILE                = 55
	IPPROTO_MPLS                  = 137
	IPPROTO_MTP                   = 92
	IPPROTO_MUX                   = 18
	IPPROTO_ND                    = 77
	IPPROTO_NHRP                  = 54
	IPPROTO_NONE                  = 59
	IPPROTO_NSP                   = 31
	IPPROTO_NVPII                 = 11
	IPPROTO_OLD_DIVERT            = 254
	IPPROTO_OSPFIGP               = 89
	IPPROTO_PFSYNC                = 240
	IPPROTO_PGM                   = 113
	IPPROTO_PIGP                  = 9
	IPPROTO_PIM                   = 103
	IPPROTO_PRM                   = 21
	IPPROTO_PUP                   = 12
	IPPROTO_PVP                   = 75
	IPPROTO_RAW                   = 255
	IPPROTO_RCCMON                = 10
	IPPROTO_RDP                   = 27
	IPPROTO_RESERVED_253          = 253
	IPPROTO_RESERVED_254          = 254
	IPPROTO_ROUTING               = 43
	IPPROTO_RSVP                  = 46
	IPPROTO_RVD                   = 66
	IPPROTO_SATEXPAK              = 64
	IPPROTO_SATMON                = 69
	IPPROTO_SCCSP                 = 96
	IPPROTO_SCTP                  = 132
	IPPROTO_SDRP                  = 42
	IPPROTO_SEND                  = 259
	IPPROTO_SHIM6                 = 140
	IPPROTO_SKIP                  = 57
	IPPROTO_SPACER                = 32767
	IPPROTO_SRPC                  = 90
	IPPROTO_ST                    = 7
	IPPROTO_SVMTP                 = 82
	IPPROTO_SWIPE                 = 53
	IPPROTO_TCF                   = 87
	IPPROTO_TCP                   = 6
	IPPROTO_TLSP                  = 56
	IPPROTO_TP                    = 29
	IPPROTO_TPXX                  = 39
	IPPROTO_TRUNK1                = 23
	IPPROTO_TRUNK2                = 24
	IPPROTO_TTP                   = 84
	IPPROTO_UDP                   = 17
	IPPROTO_UDPLITE               = 136
	IPPROTO_VINES                 = 83
	IPPROTO_VISA                  = 70
	IPPROTO_VMTP                  = 81
	IPPROTO_WBEXPAK               = 79
	IPPROTO_WBMON                 = 78
	IPPROTO_WSN                   = 74
	IPPROTO_XNET                  = 15
	IPPROTO_XTP                   = 36
	IPV6CTL_ACCEPT_RTADV          = 12
	IPV6CTL_ADDRCTLPOLICY         = 38
	IPV6CTL_AUTO_FLOWLABEL        = 17
	IPV6CTL_AUTO_LINKLOCAL        = 35
	IPV6CTL_DAD_COUNT             = 16
	IPV6CTL_DEFHLIM               = 3
	IPV6CTL_DEFMCASTHLIM          = 18
	IPV6CTL_FORWARDING            = 1
	IPV6CTL_FORWSRCRT             = 5
	IPV6CTL_GIF_HLIM              = 19
	IPV6CTL_HDRNESTLIMIT          = 15
	IPV6CTL_INTRDQMAXLEN          = 52
	IPV6CTL_INTRQMAXLEN           = 51
	IPV6CTL_KAME_VERSION          = 20
	IPV6CTL_LOG_INTERVAL          = 14
	IPV6CTL_MAXFRAGBUCKETSIZE     = 54
	IPV6CTL_MAXFRAGPACKETS        = 9
	IPV6CTL_MAXFRAGS              = 41
	IPV6CTL_MAXFRAGSPERPACKET     = 53
	IPV6CTL_MAXID                 = 55
	IPV6CTL_MCAST_PMTU            = 44
	IPV6CTL_MRTPROTO              = 8
	IPV6CTL_MRTSTATS              = 7
	IPV6CTL_NORBIT_RAIF           = 49
	IPV6CTL_NO_RADR               = 48
	IPV6CTL_PREFER_TEMPADDR       = 37
	IPV6CTL_RFC6204W3             = 50
	IPV6CTL_RIP6STATS             = 36
	IPV6CTL_RR_PRUNE              = 22
	IPV6CTL_SENDREDIRECTS         = 2
	IPV6CTL_SOURCECHECK           = 10
	IPV6CTL_SOURCECHECK_LOGINT    = 11
	IPV6CTL_STATS                 = 6
	IPV6CTL_STEALTH               = 45
	IPV6CTL_TEMPPLTIME            = 33
	IPV6CTL_TEMPVLTIME            = 34
	IPV6CTL_USETEMPADDR           = 32
	IPV6CTL_USE_DEFAULTZONE       = 39
	IPV6CTL_USE_DEPRECATED        = 21
	IPV6CTL_V6ONLY                = 24
	IPV6PORT_ANONMAX              = 65535
	IPV6PORT_ANONMIN              = 49152
	IPV6PORT_RESERVED             = 1024
	IPV6PORT_RESERVEDMAX          = 1023
	IPV6PORT_RESERVEDMIN          = 600
	IPV6PROTO_MAXID               = 104
	IPV6_AUTOFLOWLABEL            = 59
	IPV6_BINDANY                  = 64
	IPV6_BINDMULTI                = 65
	IPV6_BINDV6ONLY               = 27
	IPV6_CHECKSUM                 = 26
	IPV6_DEFAULT_MULTICAST_HOPS   = 1
	IPV6_DEFAULT_MULTICAST_LOOP   = 1
	IPV6_DONTFRAG                 = 62
	IPV6_DSTOPTS                  = 50
	IPV6_FLOWID                   = 67
	IPV6_FLOWTYPE                 = 68
	IPV6_FW_ADD                   = 30
	IPV6_FW_DEL                   = 31
	IPV6_FW_FLUSH                 = 32
	IPV6_FW_GET                   = 34
	IPV6_FW_ZERO                  = 33
	IPV6_HOPLIMIT                 = 47
	IPV6_HOPOPTS                  = 49
	IPV6_IPSEC_POLICY             = 28
	IPV6_JOIN_GROUP               = 12
	IPV6_LEAVE_GROUP              = 13
	IPV6_MAX_GROUP_SRC_FILTER     = 512
	IPV6_MAX_MEMBERSHIPS          = 4095
	IPV6_MAX_SOCK_SRC_FILTER      = 128
	IPV6_MSFILTER                 = 74
	IPV6_MULTICAST_HOPS           = 10
	IPV6_MULTICAST_IF             = 9
	IPV6_MULTICAST_LOOP           = 11
	IPV6_NEXTHOP                  = 48
	IPV6_ORIGDSTADDR              = 72
	IPV6_PATHMTU                  = 44
	IPV6_PKTINFO                  = 46
	IPV6_PORTRANGE                = 14
	IPV6_PORTRANGE_DEFAULT        = 0
	IPV6_PORTRANGE_HIGH           = 1
	IPV6_PORTRANGE_LOW            = 2
	IPV6_PREFER_TEMPADDR          = 63
	IPV6_RECVDSTOPTS              = 40
	IPV6_RECVFLOWID               = 70
	IPV6_RECVHOPLIMIT             = 37
	IPV6_RECVHOPOPTS              = 39
	IPV6_RECVORIGDSTADDR          = 72
	IPV6_RECVPATHMTU              = 43
	IPV6_RECVPKTINFO              = 36
	IPV6_RECVRSSBUCKETID          = 71
	IPV6_RECVRTHDR                = 38
	IPV6_RECVTCLASS               = 57
	IPV6_RSSBUCKETID              = 69
	IPV6_RSS_LISTEN_BUCKET        = 66
	IPV6_RTHDR                    = 51
	IPV6_RTHDRDSTOPTS             = 35
	IPV6_RTHDR_LOOSE              = 0
	IPV6_RTHDR_STRICT             = 1
	IPV6_RTHDR_TYPE_0             = 0
	IPV6_SOCKOPT_RESERVED1        = 3
	IPV6_TCLASS                   = 61
	IPV6_UNICAST_HOPS             = 4
	IPV6_USE_MIN_MTU              = 42
	IPV6_V6ONLY                   = 27
	IPV6_VLAN_PCP                 = 75
	IP_ADD_MEMBERSHIP             = 12
	IP_ADD_SOURCE_MEMBERSHIP      = 70
	IP_BINDANY                    = 24
	IP_BINDMULTI                  = 25
	IP_BLOCK_SOURCE               = 72
	IP_DEFAULT_MULTICAST_LOOP     = 1
	IP_DEFAULT_MULTICAST_TTL      = 1
	IP_DONTFRAG                   = 67
	IP_DROP_MEMBERSHIP            = 13
	IP_DROP_SOURCE_MEMBERSHIP     = 71
	IP_DUMMYNET3                  = 49
	IP_DUMMYNET_CONFIGURE         = 60
	IP_DUMMYNET_DEL               = 61
	IP_DUMMYNET_FLUSH             = 62
	IP_DUMMYNET_GET               = 64
	IP_FLOWID                     = 90
	IP_FLOWTYPE                   = 91
	IP_FW3                        = 48
	IP_FW_ADD                     = 50
	IP_FW_DEL                     = 51
	IP_FW_FLUSH                   = 52
	IP_FW_GET                     = 54
	IP_FW_NAT_CFG                 = 56
	IP_FW_NAT_DEL                 = 57
	IP_FW_NAT_GET_CONFIG          = 58
	IP_FW_NAT_GET_LOG             = 59
	IP_FW_RESETLOG                = 55
	IP_FW_TABLE_ADD               = 40
	IP_FW_TABLE_DEL               = 41
	IP_FW_TABLE_FLUSH             = 42
	IP_FW_TABLE_GETSIZE           = 43
	IP_FW_TABLE_LIST              = 44
	IP_FW_ZERO                    = 53
	IP_HDRINCL                    = 2
	IP_IPSEC_POLICY               = 21
	IP_MAX_GROUP_SRC_FILTER       = 512
	IP_MAX_MEMBERSHIPS            = 4095
	IP_MAX_SOCK_MUTE_FILTER       = 128
	IP_MAX_SOCK_SRC_FILTER        = 128
	IP_MINTTL                     = 66
	IP_MSFILTER                   = 74
	IP_MULTICAST_IF               = 9
	IP_MULTICAST_LOOP             = 11
	IP_MULTICAST_TTL              = 10
	IP_MULTICAST_VIF              = 14
	IP_ONESBCAST                  = 23
	IP_OPTIONS                    = 1
	IP_ORIGDSTADDR                = 27
	IP_PORTRANGE                  = 19
	IP_PORTRANGE_DEFAULT          = 0
	IP_PORTRANGE_HIGH             = 1
	IP_PORTRANGE_LOW              = 2
	IP_RECVDSTADDR                = 7
	IP_RECVFLOWID                 = 93
	IP_RECVIF                     = 20
	IP_RECVOPTS                   = 5
	IP_RECVORIGDSTADDR            = 27
	IP_RECVRETOPTS                = 6
	IP_RECVRSSBUCKETID            = 94
	IP_RECVTOS                    = 68
	IP_RECVTTL                    = 65
	IP_RETOPTS                    = 8
	IP_RSSBUCKETID                = 92
	IP_RSS_LISTEN_BUCKET          = 26
	IP_RSVP_OFF                   = 16
	IP_RSVP_ON                    = 15
	IP_RSVP_VIF_OFF               = 18
	IP_RSVP_VIF_ON                = 17
	IP_SENDSRCADDR                = 7
	IP_TOS                        = 3
	IP_TTL                        = 4
	IP_UNBLOCK_SOURCE             = 73
	IP_VLAN_PCP                   = 75
	LITTLE_ENDIAN                 = 1234
	MCAST_BLOCK_SOURCE            = 84
	MCAST_EXCLUDE                 = 2
	MCAST_INCLUDE                 = 1
	MCAST_JOIN_GROUP              = 80
	MCAST_JOIN_SOURCE_GROUP       = 82
	MCAST_LEAVE_GROUP             = 81
	MCAST_LEAVE_SOURCE_GROUP      = 83
	MCAST_UNBLOCK_SOURCE          = 85
	MCAST_UNDEFINED               = 0
	PDP_ENDIAN                    = 3412
	SIN6_LEN                      = 0
	X_BIG_ENDIAN                  = 4321
	X_BYTEORDER_FUNC_DEFINED      = 0
	X_BYTEORDER_PROTOTYPED        = 0
	X_BYTE_ORDER                  = 1234
	X_FILE_OFFSET_BITS            = 64
	X_ILP32                       = 1
	X_IN_ADDR_T_DECLARED          = 0
	X_IN_PORT_T_DECLARED          = 0
	X_LITTLE_ENDIAN               = 1234
	X_MACHINE_ENDIAN_H_           = 0
	X_MACHINE__LIMITS_H_          = 0
	X_MACHINE__TYPES_H_           = 0
	X_NETINET6_IN6_H_             = 0
	X_NETINET_IN_H_               = 0
	X_Nonnull                     = 0
	X_Null_unspecified            = 0
	X_Nullable                    = 0
	X_PDP_ENDIAN                  = 3412
	X_QUAD_HIGHWORD               = 1
	X_QUAD_LOWWORD                = 0
	X_SA_FAMILY_T_DECLARED        = 0
	X_SIZE_T_DECLARED             = 0
	X_SOCKLEN_T_DECLARED          = 0
	X_SS_MAXSIZE                  = 128
	X_STRUCT_IN_ADDR_DECLARED     = 0
	X_SYS_CDEFS_H_                = 0
	X_SYS__ENDIAN_H_              = 0
	X_SYS__SOCKADDR_STORAGE_H_    = 0
	X_SYS__TYPES_H_               = 0
	X_UINT16_T_DECLARED           = 0
	X_UINT32_T_DECLARED           = 0
	X_UINT8_T_DECLARED            = 0
	I386                          = 1
	Unix                          = 1
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

type Uint8_t = X__uint8_t

type Uint16_t = X__uint16_t

type Uint32_t = X__uint32_t

type In_addr_t = Uint32_t

type In_port_t = Uint16_t

type Sa_family_t = X__sa_family_t

type In_addr = struct{ Fs_addr In_addr_t }

type Socklen_t = X__socklen_t

type Sockaddr_storage = struct {
	Fss_len     uint8
	Fss_family  Sa_family_t
	F__ss_pad1  [6]int8
	F__ss_align X__int64_t
	F__ss_pad2  [112]int8
}

type Sockaddr_in = struct {
	Fsin_len    Uint8_t
	Fsin_family Sa_family_t
	Fsin_port   In_port_t
	Fsin_addr   struct{ Fs_addr In_addr_t }
	Fsin_zero   [8]int8
}

type Ip_mreq = struct {
	Fimr_multiaddr struct{ Fs_addr In_addr_t }
	Fimr_interface struct{ Fs_addr In_addr_t }
}

type Ip_mreqn = struct {
	Fimr_multiaddr struct{ Fs_addr In_addr_t }
	Fimr_address   struct{ Fs_addr In_addr_t }
	Fimr_ifindex   int32
}

type Ip_mreq_source = struct {
	Fimr_multiaddr  struct{ Fs_addr In_addr_t }
	Fimr_sourceaddr struct{ Fs_addr In_addr_t }
	Fimr_interface  struct{ Fs_addr In_addr_t }
}

type Group_req = struct {
	Fgr_interface Uint32_t
	Fgr_group     struct {
		Fss_len     uint8
		Fss_family  Sa_family_t
		F__ss_pad1  [6]int8
		F__ss_align X__int64_t
		F__ss_pad2  [112]int8
	}
}

type Group_source_req = struct {
	Fgsr_interface Uint32_t
	Fgsr_group     struct {
		Fss_len     uint8
		Fss_family  Sa_family_t
		F__ss_pad1  [6]int8
		F__ss_align X__int64_t
		F__ss_pad2  [112]int8
	}
	Fgsr_source struct {
		Fss_len     uint8
		Fss_family  Sa_family_t
		F__ss_pad1  [6]int8
		F__ss_align X__int64_t
		F__ss_pad2  [112]int8
	}
}

type X__msfilterreq = struct {
	Fmsfr_ifindex Uint32_t
	Fmsfr_fmode   Uint32_t
	Fmsfr_nsrcs   Uint32_t
	Fmsfr_group   struct {
		Fss_len     uint8
		Fss_family  Sa_family_t
		F__ss_pad1  [6]int8
		F__ss_align X__int64_t
		F__ss_pad2  [112]int8
	}
	Fmsfr_srcs uintptr
}

type In6_addr = struct {
	F__u6_addr struct {
		F__ccgo_pad1 [0]uint32
		F__u6_addr8  [16]Uint8_t
	}
}

type Sockaddr_in6 = struct {
	Fsin6_len      Uint8_t
	Fsin6_family   Sa_family_t
	Fsin6_port     In_port_t
	Fsin6_flowinfo Uint32_t
	Fsin6_addr     struct {
		F__u6_addr struct {
			F__ccgo_pad1 [0]uint32
			F__u6_addr8  [16]Uint8_t
		}
	}
	Fsin6_scope_id Uint32_t
}

type Route_in6 = struct {
	Fro_nh      uintptr
	Fro_lle     uintptr
	Fro_prepend uintptr
	Fro_plen    Uint16_t
	Fro_flags   Uint16_t
	Fro_mtu     Uint16_t
	Fspare      Uint16_t
	Fro_dst     struct {
		Fsin6_len      Uint8_t
		Fsin6_family   Sa_family_t
		Fsin6_port     In_port_t
		Fsin6_flowinfo Uint32_t
		Fsin6_addr     struct {
			F__u6_addr struct {
				F__ccgo_pad1 [0]uint32
				F__u6_addr8  [16]Uint8_t
			}
		}
		Fsin6_scope_id Uint32_t
	}
}

type Ipv6_mreq = struct {
	Fipv6mr_multiaddr struct {
		F__u6_addr struct {
			F__ccgo_pad1 [0]uint32
			F__u6_addr8  [16]Uint8_t
		}
	}
	Fipv6mr_interface uint32
}

type In6_pktinfo = struct {
	Fipi6_addr struct {
		F__u6_addr struct {
			F__ccgo_pad1 [0]uint32
			F__u6_addr8  [16]Uint8_t
		}
	}
	Fipi6_ifindex uint32
}

type Ip6_mtuinfo = struct {
	Fip6m_addr struct {
		Fsin6_len      Uint8_t
		Fsin6_family   Sa_family_t
		Fsin6_port     In_port_t
		Fsin6_flowinfo Uint32_t
		Fsin6_addr     struct {
			F__u6_addr struct {
				F__ccgo_pad1 [0]uint32
				F__u6_addr8  [16]Uint8_t
			}
		}
		Fsin6_scope_id Uint32_t
	}
	Fip6m_mtu Uint32_t
}

var _ int8
