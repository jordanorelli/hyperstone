// Code generated by protoc-gen-go.
// source: network_connection.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ENetworkDisconnectionReason int32

const (
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID                     ENetworkDisconnectionReason = 0
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SHUTDOWN                    ENetworkDisconnectionReason = 1
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_USER          ENetworkDisconnectionReason = 2
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_SERVER        ENetworkDisconnectionReason = 3
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOST                        ENetworkDisconnectionReason = 4
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_OVERFLOW                    ENetworkDisconnectionReason = 5
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_BANNED                ENetworkDisconnectionReason = 6
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_INUSE                 ENetworkDisconnectionReason = 7
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_TICKET                ENetworkDisconnectionReason = 8
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGON                 ENetworkDisconnectionReason = 9
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHCANCELLED         ENetworkDisconnectionReason = 10
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED       ENetworkDisconnectionReason = 11
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHINVALID           ENetworkDisconnectionReason = 12
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VACBANSTATE           ENetworkDisconnectionReason = 13
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE   ENetworkDisconnectionReason = 14
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT    ENetworkDisconnectionReason = 15
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DROPPED               ENetworkDisconnectionReason = 16
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_OWNERSHIP             ENetworkDisconnectionReason = 17
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVERINFO_OVERFLOW         ENetworkDisconnectionReason = 18
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TICKMSG_OVERFLOW            ENetworkDisconnectionReason = 19
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW     ENetworkDisconnectionReason = 20
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW        ENetworkDisconnectionReason = 21
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW         ENetworkDisconnectionReason = 22
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW          ENetworkDisconnectionReason = 23
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTOVERFLOW            ENetworkDisconnectionReason = 24
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTERROR               ENetworkDisconnectionReason = 25
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RELIABLEOVERFLOW            ENetworkDisconnectionReason = 26
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADDELTATICK                ENetworkDisconnectionReason = 27
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOMORESPLITS                ENetworkDisconnectionReason = 28
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TIMEDOUT                    ENetworkDisconnectionReason = 29
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECTED                ENetworkDisconnectionReason = 30
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LEAVINGSPLIT                ENetworkDisconnectionReason = 31
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIFFERENTCLASSTABLES        ENetworkDisconnectionReason = 32
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADRELAYPASSWORD            ENetworkDisconnectionReason = 33
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADSPECTATORPASSWORD        ENetworkDisconnectionReason = 34
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVRESTRICTED              ENetworkDisconnectionReason = 35
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOSPECTATORS                ENetworkDisconnectionReason = 36
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVUNAVAILABLE             ENetworkDisconnectionReason = 37
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVSTOP                    ENetworkDisconnectionReason = 38
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKED                      ENetworkDisconnectionReason = 39
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BANADDED                    ENetworkDisconnectionReason = 40
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKBANADDED                ENetworkDisconnectionReason = 41
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVDIRECT                  ENetworkDisconnectionReason = 42
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA      ENetworkDisconnectionReason = 43
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_MISMATCH         ENetworkDisconnectionReason = 44
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_USERCMD                     ENetworkDisconnectionReason = 45
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REJECTED_BY_GAME            ENetworkDisconnectionReason = 46
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR         ENetworkDisconnectionReason = 47
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR       ENetworkDisconnectionReason = 48
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BAD_SERVER_PASSWORD         ENetworkDisconnectionReason = 49
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION  ENetworkDisconnectionReason = 50
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECTION_FAILURE          ENetworkDisconnectionReason = 51
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS      ENetworkDisconnectionReason = 52
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RECONNECTION                ENetworkDisconnectionReason = 53
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPSHUTDOWN                ENetworkDisconnectionReason = 54
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPDEACTIVATE              ENetworkDisconnectionReason = 55
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HOST_ENDGAME                ENetworkDisconnectionReason = 56
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE     ENetworkDisconnectionReason = 57
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CREATE_SERVER_FAILED        ENetworkDisconnectionReason = 58
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_EXITING                     ENetworkDisconnectionReason = 59
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE      ENetworkDisconnectionReason = 60
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY ENetworkDisconnectionReason = 61
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL     ENetworkDisconnectionReason = 62
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP    ENetworkDisconnectionReason = 63
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_NO_MAP               ENetworkDisconnectionReason = 64
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP        ENetworkDisconnectionReason = 65
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM       ENetworkDisconnectionReason = 66
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_MISC             ENetworkDisconnectionReason = 67
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT   ENetworkDisconnectionReason = 68
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_SHUTDOWN             ENetworkDisconnectionReason = 69
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW   ENetworkDisconnectionReason = 70
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE         ENetworkDisconnectionReason = 71
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT    ENetworkDisconnectionReason = 72
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_INCOMPATIBLE         ENetworkDisconnectionReason = 73
)

var ENetworkDisconnectionReason_name = map[int32]string{
	0:  "NETWORK_DISCONNECT_INVALID",
	1:  "NETWORK_DISCONNECT_SHUTDOWN",
	2:  "NETWORK_DISCONNECT_DISCONNECT_BY_USER",
	3:  "NETWORK_DISCONNECT_DISCONNECT_BY_SERVER",
	4:  "NETWORK_DISCONNECT_LOST",
	5:  "NETWORK_DISCONNECT_OVERFLOW",
	6:  "NETWORK_DISCONNECT_STEAM_BANNED",
	7:  "NETWORK_DISCONNECT_STEAM_INUSE",
	8:  "NETWORK_DISCONNECT_STEAM_TICKET",
	9:  "NETWORK_DISCONNECT_STEAM_LOGON",
	10: "NETWORK_DISCONNECT_STEAM_AUTHCANCELLED",
	11: "NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED",
	12: "NETWORK_DISCONNECT_STEAM_AUTHINVALID",
	13: "NETWORK_DISCONNECT_STEAM_VACBANSTATE",
	14: "NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE",
	15: "NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT",
	16: "NETWORK_DISCONNECT_STEAM_DROPPED",
	17: "NETWORK_DISCONNECT_STEAM_OWNERSHIP",
	18: "NETWORK_DISCONNECT_SERVERINFO_OVERFLOW",
	19: "NETWORK_DISCONNECT_TICKMSG_OVERFLOW",
	20: "NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW",
	21: "NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW",
	22: "NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW",
	23: "NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW",
	24: "NETWORK_DISCONNECT_SNAPSHOTOVERFLOW",
	25: "NETWORK_DISCONNECT_SNAPSHOTERROR",
	26: "NETWORK_DISCONNECT_RELIABLEOVERFLOW",
	27: "NETWORK_DISCONNECT_BADDELTATICK",
	28: "NETWORK_DISCONNECT_NOMORESPLITS",
	29: "NETWORK_DISCONNECT_TIMEDOUT",
	30: "NETWORK_DISCONNECT_DISCONNECTED",
	31: "NETWORK_DISCONNECT_LEAVINGSPLIT",
	32: "NETWORK_DISCONNECT_DIFFERENTCLASSTABLES",
	33: "NETWORK_DISCONNECT_BADRELAYPASSWORD",
	34: "NETWORK_DISCONNECT_BADSPECTATORPASSWORD",
	35: "NETWORK_DISCONNECT_HLTVRESTRICTED",
	36: "NETWORK_DISCONNECT_NOSPECTATORS",
	37: "NETWORK_DISCONNECT_HLTVUNAVAILABLE",
	38: "NETWORK_DISCONNECT_HLTVSTOP",
	39: "NETWORK_DISCONNECT_KICKED",
	40: "NETWORK_DISCONNECT_BANADDED",
	41: "NETWORK_DISCONNECT_KICKBANADDED",
	42: "NETWORK_DISCONNECT_HLTVDIRECT",
	43: "NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA",
	44: "NETWORK_DISCONNECT_PURESERVER_MISMATCH",
	45: "NETWORK_DISCONNECT_USERCMD",
	46: "NETWORK_DISCONNECT_REJECTED_BY_GAME",
	47: "NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR",
	48: "NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR",
	49: "NETWORK_DISCONNECT_BAD_SERVER_PASSWORD",
	50: "NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION",
	51: "NETWORK_DISCONNECT_CONNECTION_FAILURE",
	52: "NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS",
	53: "NETWORK_DISCONNECT_RECONNECTION",
	54: "NETWORK_DISCONNECT_LOOPSHUTDOWN",
	55: "NETWORK_DISCONNECT_LOOPDEACTIVATE",
	56: "NETWORK_DISCONNECT_HOST_ENDGAME",
	57: "NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE",
	58: "NETWORK_DISCONNECT_CREATE_SERVER_FAILED",
	59: "NETWORK_DISCONNECT_EXITING",
	60: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE",
	61: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY",
	62: "NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL",
	63: "NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP",
	64: "NETWORK_DISCONNECT_CLIENT_NO_MAP",
	65: "NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP",
	66: "NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM",
	67: "NETWORK_DISCONNECT_STEAM_DENY_MISC",
	68: "NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT",
	69: "NETWORK_DISCONNECT_SERVER_SHUTDOWN",
	70: "NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW",
	71: "NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE",
	72: "NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT",
	73: "NETWORK_DISCONNECT_SERVER_INCOMPATIBLE",
}
var ENetworkDisconnectionReason_value = map[string]int32{
	"NETWORK_DISCONNECT_INVALID":                     0,
	"NETWORK_DISCONNECT_SHUTDOWN":                    1,
	"NETWORK_DISCONNECT_DISCONNECT_BY_USER":          2,
	"NETWORK_DISCONNECT_DISCONNECT_BY_SERVER":        3,
	"NETWORK_DISCONNECT_LOST":                        4,
	"NETWORK_DISCONNECT_OVERFLOW":                    5,
	"NETWORK_DISCONNECT_STEAM_BANNED":                6,
	"NETWORK_DISCONNECT_STEAM_INUSE":                 7,
	"NETWORK_DISCONNECT_STEAM_TICKET":                8,
	"NETWORK_DISCONNECT_STEAM_LOGON":                 9,
	"NETWORK_DISCONNECT_STEAM_AUTHCANCELLED":         10,
	"NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED":       11,
	"NETWORK_DISCONNECT_STEAM_AUTHINVALID":           12,
	"NETWORK_DISCONNECT_STEAM_VACBANSTATE":           13,
	"NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE":   14,
	"NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT":    15,
	"NETWORK_DISCONNECT_STEAM_DROPPED":               16,
	"NETWORK_DISCONNECT_STEAM_OWNERSHIP":             17,
	"NETWORK_DISCONNECT_SERVERINFO_OVERFLOW":         18,
	"NETWORK_DISCONNECT_TICKMSG_OVERFLOW":            19,
	"NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW":     20,
	"NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW":        21,
	"NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW":         22,
	"NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW":          23,
	"NETWORK_DISCONNECT_SNAPSHOTOVERFLOW":            24,
	"NETWORK_DISCONNECT_SNAPSHOTERROR":               25,
	"NETWORK_DISCONNECT_RELIABLEOVERFLOW":            26,
	"NETWORK_DISCONNECT_BADDELTATICK":                27,
	"NETWORK_DISCONNECT_NOMORESPLITS":                28,
	"NETWORK_DISCONNECT_TIMEDOUT":                    29,
	"NETWORK_DISCONNECT_DISCONNECTED":                30,
	"NETWORK_DISCONNECT_LEAVINGSPLIT":                31,
	"NETWORK_DISCONNECT_DIFFERENTCLASSTABLES":        32,
	"NETWORK_DISCONNECT_BADRELAYPASSWORD":            33,
	"NETWORK_DISCONNECT_BADSPECTATORPASSWORD":        34,
	"NETWORK_DISCONNECT_HLTVRESTRICTED":              35,
	"NETWORK_DISCONNECT_NOSPECTATORS":                36,
	"NETWORK_DISCONNECT_HLTVUNAVAILABLE":             37,
	"NETWORK_DISCONNECT_HLTVSTOP":                    38,
	"NETWORK_DISCONNECT_KICKED":                      39,
	"NETWORK_DISCONNECT_BANADDED":                    40,
	"NETWORK_DISCONNECT_KICKBANADDED":                41,
	"NETWORK_DISCONNECT_HLTVDIRECT":                  42,
	"NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA":      43,
	"NETWORK_DISCONNECT_PURESERVER_MISMATCH":         44,
	"NETWORK_DISCONNECT_USERCMD":                     45,
	"NETWORK_DISCONNECT_REJECTED_BY_GAME":            46,
	"NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR":         47,
	"NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR":       48,
	"NETWORK_DISCONNECT_BAD_SERVER_PASSWORD":         49,
	"NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION":  50,
	"NETWORK_DISCONNECT_CONNECTION_FAILURE":          51,
	"NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS":      52,
	"NETWORK_DISCONNECT_RECONNECTION":                53,
	"NETWORK_DISCONNECT_LOOPSHUTDOWN":                54,
	"NETWORK_DISCONNECT_LOOPDEACTIVATE":              55,
	"NETWORK_DISCONNECT_HOST_ENDGAME":                56,
	"NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE":     57,
	"NETWORK_DISCONNECT_CREATE_SERVER_FAILED":        58,
	"NETWORK_DISCONNECT_EXITING":                     59,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE":      60,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY": 61,
	"NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL":     62,
	"NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP":    63,
	"NETWORK_DISCONNECT_CLIENT_NO_MAP":               64,
	"NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP":        65,
	"NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM":       66,
	"NETWORK_DISCONNECT_STEAM_DENY_MISC":             67,
	"NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT":   68,
	"NETWORK_DISCONNECT_SERVER_SHUTDOWN":             69,
	"NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW":   70,
	"NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE":         71,
	"NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT":    72,
	"NETWORK_DISCONNECT_SERVER_INCOMPATIBLE":         73,
}

func (x ENetworkDisconnectionReason) Enum() *ENetworkDisconnectionReason {
	p := new(ENetworkDisconnectionReason)
	*p = x
	return p
}
func (x ENetworkDisconnectionReason) String() string {
	return proto.EnumName(ENetworkDisconnectionReason_name, int32(x))
}
func (x *ENetworkDisconnectionReason) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ENetworkDisconnectionReason_value, data, "ENetworkDisconnectionReason")
	if err != nil {
		return err
	}
	*x = ENetworkDisconnectionReason(value)
	return nil
}
func (ENetworkDisconnectionReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

var E_NetworkConnectionToken = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50500,
	Name:          "dota.network_connection_token",
	Tag:           "bytes,50500,opt,name=network_connection_token",
}

func init() {
	proto.RegisterEnum("dota.ENetworkDisconnectionReason", ENetworkDisconnectionReason_name, ENetworkDisconnectionReason_value)
	proto.RegisterExtension(E_NetworkConnectionToken)
}

func init() { proto.RegisterFile("network_connection.proto", fileDescriptor32) }

var fileDescriptor32 = []byte{
	// 1903 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x98, 0xeb, 0x76, 0xdb, 0xc6,
	0x11, 0xc7, 0xe3, 0x3a, 0xbd, 0x04, 0xbd, 0x6d, 0x91, 0xb6, 0xb6, 0xa1, 0x58, 0x57, 0x5f, 0x24,
	0x5f, 0xe4, 0xa4, 0xf7, 0xa6, 0x49, 0x9a, 0x25, 0xb0, 0x22, 0x61, 0xe1, 0x16, 0xec, 0x92, 0xb6,
	0x7a, 0x72, 0x0e, 0x0e, 0x4c, 0xac, 0x24, 0xd4, 0x24, 0x80, 0x02, 0xa0, 0x1c, 0x7f, 0xeb, 0xe7,
	0x7e, 0xe9, 0x3b, 0xf4, 0x59, 0xfa, 0x08, 0x7d, 0xa0, 0xce, 0x02, 0x24, 0x68, 0xe6, 0xec, 0x22,
	0xfa, 0x64, 0xc8, 0x5c, 0xfc, 0x30, 0x3b, 0x3b, 0xf3, 0x9f, 0x99, 0xd5, 0x6e, 0x67, 0xbc, 0x7e,
	0x93, 0x97, 0xaf, 0xa3, 0x69, 0x9e, 0x65, 0x7c, 0x5a, 0xa7, 0x79, 0x76, 0x5c, 0x94, 0x79, 0x9d,
	0xeb, 0xef, 0x27, 0x79, 0x1d, 0x1b, 0xbb, 0x17, 0x79, 0x7e, 0x31, 0xe3, 0xcf, 0x9a, 0xff, 0x7b,
	0xb5, 0x38, 0x7f, 0x96, 0xf0, 0x6a, 0x5a, 0xa6, 0x45, 0x9d, 0x97, 0xed, 0xba, 0x47, 0xff, 0x7e,
	0xaa, 0x6d, 0x11, 0xaf, 0xa5, 0x58, 0x69, 0xb5, 0xe6, 0x84, 0x3c, 0xae, 0xf2, 0x4c, 0xdf, 0xd6,
	0x0c, 0x8f, 0xb0, 0x17, 0x7e, 0x78, 0x1a, 0x59, 0x36, 0x35, 0x7d, 0xcf, 0x23, 0x26, 0x8b, 0x6c,
	0x6f, 0x82, 0x1d, 0xdb, 0x42, 0xef, 0xe9, 0x3b, 0xda, 0x96, 0xe4, 0x77, 0x3a, 0x1a, 0x33, 0xcb,
	0x7f, 0xe1, 0xa1, 0x1b, 0xfa, 0x89, 0x76, 0x5f, 0xb2, 0xe0, 0x9d, 0xc7, 0xc1, 0x59, 0x34, 0xa6,
	0x24, 0x44, 0xdf, 0x33, 0xb6, 0xfe, 0xf3, 0xbf, 0xdb, 0xb7, 0x0e, 0x86, 0xf1, 0x9c, 0x8f, 0xed,
	0x68, 0x6d, 0x4c, 0x34, 0xae, 0x78, 0xa9, 0x3f, 0xd7, 0x1e, 0x7e, 0x27, 0x07, 0x30, 0x13, 0x20,
	0xdd, 0x34, 0xee, 0x02, 0xe9, 0x8e, 0x84, 0x44, 0x79, 0x79, 0x05, 0xac, 0x81, 0x76, 0x4b, 0xc2,
	0x72, 0x7c, 0xca, 0xd0, 0xfb, 0xc6, 0x7d, 0x78, 0x77, 0x4f, 0xf2, 0xae, 0xd9, 0xb9, 0xc6, 0xc9,
	0xab, 0x1a, 0xec, 0x91, 0x6d, 0xdc, 0x87, 0xef, 0x9f, 0x38, 0xfe, 0x0b, 0xf4, 0x7d, 0xe3, 0x08,
	0x38, 0xf7, 0x7b, 0x39, 0x3e, 0xd8, 0x72, 0x3e, 0xcb, 0xdf, 0xe8, 0xb6, 0xb6, 0x23, 0x73, 0x22,
	0x23, 0xd8, 0x8d, 0x06, 0x18, 0xfe, 0xb2, 0xd0, 0x0f, 0x8c, 0x7b, 0xc0, 0xdb, 0x95, 0xed, 0xa9,
	0xe6, 0xf1, 0xdc, 0xb6, 0x06, 0x31, 0xfc, 0x99, 0xe8, 0x43, 0x6d, 0x5b, 0x89, 0xb2, 0x3d, 0x70,
	0x34, 0xfa, 0xa1, 0x71, 0x00, 0xa4, 0x1d, 0x35, 0xc9, 0xce, 0xc0, 0xe1, 0x00, 0x52, 0xdb, 0xc4,
	0x6c, 0xf3, 0x94, 0x30, 0xf4, 0x23, 0x63, 0x1f, 0x48, 0xdb, 0x2a, 0x12, 0x4b, 0xa7, 0xaf, 0x79,
	0xad, 0x93, 0x1e, 0x8b, 0x1c, 0x7f, 0xe8, 0x7b, 0xe8, 0x03, 0x63, 0x0f, 0x38, 0x77, 0x55, 0x1c,
	0x27, 0xbf, 0x80, 0x40, 0x74, 0xb5, 0x07, 0x4a, 0x0c, 0x1e, 0xb3, 0x91, 0x89, 0x3d, 0x93, 0x38,
	0x0e, 0xb8, 0x4a, 0xbb, 0x0e, 0xce, 0xd7, 0x0e, 0x7b, 0x71, 0xd8, 0x09, 0x09, 0xb6, 0xce, 0xc0,
	0x61, 0x16, 0xfa, 0xf1, 0x75, 0x80, 0xa7, 0xda, 0xbd, 0x5e, 0xe0, 0x2a, 0x65, 0x7e, 0x72, 0x1d,
	0x98, 0xdd, 0x03, 0x9b, 0x60, 0x13, 0x62, 0x82, 0x32, 0xcc, 0x08, 0xfa, 0xa9, 0xb1, 0x03, 0xb0,
	0x2d, 0x15, 0x0c, 0x96, 0xea, 0x54, 0x7b, 0xd2, 0xe7, 0xfe, 0x21, 0xb1, 0x20, 0x2e, 0x22, 0xe2,
	0x50, 0xf2, 0x62, 0x44, 0x42, 0x82, 0x7e, 0xd6, 0x6f, 0x5f, 0x1b, 0x1c, 0x4c, 0x7b, 0xdc, 0x67,
	0x5f, 0x64, 0x8e, 0x88, 0x79, 0x0a, 0x61, 0xe2, 0x12, 0xcb, 0x1f, 0x33, 0xf4, 0xf3, 0xfe, 0x90,
	0x63, 0xe9, 0x9c, 0xfb, 0x8b, 0x1a, 0x76, 0xbd, 0xab, 0xa4, 0x5a, 0xa1, 0x1f, 0x04, 0x70, 0x16,
	0xa8, 0x1f, 0x65, 0x95, 0x79, 0x51, 0x40, 0x1a, 0xb8, 0xda, 0xbe, 0x12, 0x05, 0xba, 0x44, 0x42,
	0x3a, 0xb2, 0x03, 0xf4, 0x0b, 0x65, 0xb2, 0x37, 0x30, 0xff, 0x4d, 0xc6, 0xcb, 0xea, 0x32, 0x2d,
	0xf4, 0xb1, 0x3c, 0xf8, 0x1a, 0xb9, 0xb1, 0xbd, 0x13, 0x7f, 0x9d, 0xf7, 0xba, 0x32, 0xef, 0x5b,
	0xed, 0xb1, 0xb3, 0xf3, 0xbc, 0xcb, 0xfb, 0x53, 0xed, 0x40, 0x82, 0x15, 0xd9, 0xe5, 0xd2, 0xe1,
	0x9a, 0xf9, 0xa1, 0x32, 0xcf, 0x44, 0x8a, 0xb9, 0xbc, 0xaa, 0xe2, 0x0b, 0xae, 0x9f, 0x69, 0x8f,
	0xa4, 0x5b, 0x06, 0x03, 0x87, 0x0c, 0x0f, 0x1c, 0xb2, 0xc1, 0xfc, 0xa5, 0xda, 0xce, 0xba, 0x4c,
	0xb3, 0x0b, 0x16, 0xbf, 0x9a, 0xf1, 0x15, 0x9a, 0xca, 0xb5, 0x97, 0x38, 0x0c, 0x13, 0x8f, 0x6d,
	0x70, 0x7f, 0x65, 0x3c, 0x00, 0xee, 0xbe, 0x84, 0x6b, 0xf1, 0x59, 0x1d, 0x93, 0xac, 0x5e, 0x41,
	0xbf, 0x92, 0xfa, 0x94, 0x11, 0x37, 0xf8, 0x36, 0xf3, 0xd7, 0xca, 0x63, 0x62, 0x7c, 0x5e, 0xbc,
	0x83, 0xf4, 0xa5, 0xb5, 0x86, 0xfa, 0x63, 0xcf, 0xa2, 0x1b, 0xc4, 0x5b, 0x6a, 0x35, 0xcd, 0x17,
	0x59, 0x52, 0xad, 0x80, 0x81, 0xf4, 0x80, 0xa8, 0x87, 0x03, 0x3a, 0xf2, 0x59, 0x87, 0xbb, 0x6d,
	0x3c, 0x04, 0xdc, 0x81, 0x0c, 0x97, 0xc5, 0x45, 0x75, 0x99, 0xd7, 0xdd, 0x91, 0x3f, 0x97, 0xc7,
	0xf8, 0x92, 0x48, 0xc2, 0xd0, 0x0f, 0xd1, 0x1d, 0xb5, 0x75, 0x4b, 0x1c, 0x29, 0xcb, 0xbc, 0x54,
	0x58, 0x17, 0x12, 0xc7, 0x16, 0xc7, 0xdd, 0x59, 0x67, 0x28, 0xad, 0x0b, 0xf9, 0x2c, 0x15, 0xe7,
	0xdc, 0x59, 0xe7, 0x49, 0x45, 0x7f, 0x80, 0xad, 0xe6, 0xac, 0x45, 0x60, 0xa2, 0x2d, 0x65, 0xe0,
	0x0c, 0xe2, 0xc4, 0x9c, 0xa5, 0x3c, 0xab, 0x9b, 0x93, 0x16, 0xa1, 0xa9, 0x8f, 0xa4, 0x3c, 0xcf,
	0x77, 0xfd, 0x90, 0xd0, 0xc0, 0xb1, 0x19, 0x45, 0x1f, 0x29, 0x13, 0xda, 0xcb, 0xdd, 0xbc, 0xe4,
	0xb4, 0x98, 0xa5, 0x75, 0xa5, 0x7f, 0x29, 0x2d, 0xb7, 0x9d, 0xc2, 0xdc, 0x55, 0x0a, 0xa1, 0x10,
	0x97, 0x44, 0xa8, 0x8b, 0xdc, 0x96, 0xf5, 0x23, 0x88, 0xcb, 0xb6, 0xd2, 0x96, 0xf5, 0x23, 0x88,
	0x8b, 0x9c, 0xe4, 0x10, 0x3c, 0x81, 0x54, 0x6b, 0xb6, 0x85, 0x76, 0x94, 0x24, 0x87, 0xc7, 0x57,
	0x90, 0x5f, 0xcd, 0xb6, 0xf4, 0x97, 0x8a, 0xa6, 0xe6, 0xe4, 0x04, 0xa4, 0xd8, 0x63, 0xa6, 0x83,
	0x29, 0x6d, 0x72, 0x97, 0xa2, 0x5d, 0xe3, 0x31, 0x10, 0x1f, 0x4a, 0x6d, 0x3b, 0x3f, 0xe7, 0x25,
	0xf8, 0xdd, 0x9c, 0xc5, 0x55, 0xd5, 0x24, 0x6e, 0xa5, 0x88, 0x0d, 0x38, 0x49, 0x08, 0x0f, 0x7c,
	0x16, 0x00, 0x17, 0x7e, 0xb3, 0xd0, 0x9e, 0x32, 0x36, 0xe0, 0x34, 0x21, 0x3c, 0xe2, 0xb7, 0x01,
	0x30, 0xa1, 0x39, 0x4c, 0x14, 0xb6, 0x02, 0x91, 0x06, 0xf0, 0x80, 0x99, 0x1f, 0x76, 0xd4, 0x7d,
	0xa5, 0xad, 0x40, 0xa5, 0x05, 0x3c, 0xc4, 0xd0, 0x81, 0x76, 0x64, 0x47, 0xdb, 0x93, 0x90, 0x47,
	0x0e, 0x9b, 0x40, 0x94, 0x80, 0x7e, 0x35, 0x67, 0x73, 0xa0, 0x14, 0x81, 0x66, 0x21, 0xaf, 0x40,
	0xb6, 0x7a, 0x4e, 0xc7, 0xf3, 0x3b, 0x33, 0x29, 0xba, 0xd7, 0x13, 0x73, 0x9d, 0x79, 0x15, 0x64,
	0xc3, 0xbe, 0xc2, 0xae, 0xb1, 0x87, 0x27, 0xd8, 0x76, 0xc4, 0xc9, 0xa0, 0xfb, 0x4a, 0xc5, 0x6b,
	0x56, 0x66, 0xf1, 0x55, 0x9c, 0xce, 0xc4, 0xa1, 0x28, 0x62, 0x58, 0xac, 0xa2, 0xcc, 0x0f, 0xd0,
	0x03, 0x65, 0x0c, 0x37, 0x4b, 0xea, 0xbc, 0xd0, 0x3f, 0xd3, 0xee, 0x48, 0x08, 0xa7, 0xa2, 0x1d,
	0xb3, 0xd0, 0x43, 0x65, 0xdb, 0x7b, 0x2a, 0x3a, 0xb1, 0x44, 0xf1, 0x7d, 0x68, 0x26, 0x44, 0x7e,
	0x5b, 0xe8, 0x50, 0xf9, 0x7d, 0xe8, 0x2d, 0x71, 0x92, 0x28, 0x7d, 0x2b, 0xbe, 0xdf, 0x51, 0x8e,
	0x94, 0xbe, 0x15, 0x56, 0x74, 0x24, 0x4b, 0xbb, 0xab, 0xf0, 0x85, 0x65, 0x87, 0xf0, 0x88, 0x1e,
	0x29, 0xfb, 0x90, 0x66, 0x51, 0x5a, 0xc2, 0xa3, 0xfe, 0xb5, 0x76, 0x24, 0xa1, 0x04, 0x63, 0x88,
	0x9b, 0xa6, 0x36, 0x47, 0xa6, 0x63, 0x43, 0x1e, 0x91, 0x97, 0x2c, 0xc4, 0xe8, 0xb1, 0xf1, 0x14,
	0x88, 0x47, 0x12, 0x62, 0xb0, 0x00, 0x9d, 0x69, 0xca, 0x73, 0xd4, 0x2a, 0x18, 0xf9, 0xa6, 0x2e,
	0x63, 0x7d, 0x22, 0xad, 0x50, 0xef, 0xd0, 0x5d, 0x9b, 0xba, 0x98, 0x99, 0x23, 0xf4, 0xc4, 0x78,
	0x04, 0xe8, 0x07, 0xfd, 0x68, 0x37, 0xad, 0xe6, 0x71, 0x3d, 0xbd, 0xd4, 0xbf, 0x90, 0xce, 0x54,
	0x62, 0x08, 0x32, 0x5d, 0x0b, 0x3d, 0x35, 0xb6, 0x81, 0x65, 0x28, 0xe6, 0x20, 0x73, 0x9e, 0x40,
	0x5c, 0xca, 0x75, 0xff, 0x79, 0xa3, 0x62, 0x62, 0x10, 0x1a, 0x62, 0x97, 0xa0, 0x63, 0x65, 0xc6,
	0x84, 0xfc, 0xef, 0x8d, 0x92, 0x0d, 0xde, 0x8a, 0xdf, 0xa0, 0x9b, 0x93, 0xed, 0xd3, 0x25, 0x94,
	0xe2, 0x21, 0x89, 0x02, 0x1c, 0x52, 0x12, 0xb5, 0x95, 0xe9, 0x99, 0x71, 0x08, 0xc8, 0x7b, 0x12,
	0xe4, 0xb2, 0x62, 0x06, 0x71, 0x59, 0xf1, 0xb6, 0x3a, 0xbd, 0x94, 0x76, 0xd8, 0xcb, 0x36, 0xb8,
	0xa3, 0xb7, 0xdc, 0x8f, 0x95, 0xfe, 0xb3, 0xb3, 0xab, 0x78, 0x96, 0x26, 0x4b, 0x7c, 0x4b, 0x96,
	0xdb, 0x0b, 0x4a, 0xb4, 0xec, 0xc8, 0xa2, 0x4e, 0x88, 0x3e, 0x51, 0xda, 0x2b, 0x84, 0xa8, 0x39,
	0x96, 0x4e, 0x85, 0x3e, 0xd1, 0x9e, 0x4a, 0xb5, 0x58, 0x44, 0x63, 0xb4, 0x76, 0xb2, 0xf8, 0x04,
	0x54, 0x42, 0x18, 0x5b, 0x7e, 0x03, 0x7d, 0x91, 0xac, 0xdf, 0x58, 0xfe, 0x0b, 0x4b, 0xa2, 0x13,
	0x10, 0x09, 0x08, 0x17, 0xf4, 0x5b, 0xa5, 0x1d, 0xeb, 0x69, 0xf0, 0x04, 0x64, 0x02, 0x82, 0x05,
	0xfa, 0xb8, 0x23, 0xa9, 0x7e, 0x45, 0x01, 0x81, 0xad, 0x0d, 0x43, 0x7f, 0x1c, 0x44, 0x23, 0xec,
	0x59, 0x0e, 0x34, 0xb1, 0xe8, 0x77, 0x4a, 0xc7, 0x79, 0x79, 0xc0, 0x79, 0x39, 0x2c, 0xf3, 0x45,
	0x31, 0x8a, 0xb3, 0x64, 0x06, 0x9d, 0xac, 0x7e, 0x20, 0x4d, 0x5f, 0xd8, 0x60, 0x67, 0x31, 0xfa,
	0xbd, 0xaa, 0xba, 0xf9, 0x7e, 0xd0, 0x4d, 0xf5, 0x7f, 0x50, 0x57, 0xb7, 0x3c, 0x2f, 0xe8, 0xe5,
	0xa2, 0x4e, 0xf2, 0x37, 0x99, 0x42, 0xd7, 0x05, 0xc9, 0x22, 0x18, 0x3e, 0x37, 0x11, 0x23, 0xcc,
	0x1f, 0x95, 0x51, 0x2a, 0x58, 0x16, 0x8f, 0xc1, 0x39, 0x57, 0x71, 0xcd, 0x15, 0x76, 0x8d, 0x60,
	0x68, 0x8f, 0x88, 0x67, 0x35, 0x11, 0xff, 0x27, 0xa5, 0x5d, 0x23, 0x18, 0xd9, 0x23, 0x92, 0x25,
	0x4d, 0xbc, 0x7f, 0x2d, 0xed, 0x94, 0x85, 0x5d, 0x50, 0xc4, 0x27, 0x30, 0x44, 0xfa, 0x10, 0x4a,
	0x9d, 0x81, 0x7f, 0x36, 0x9e, 0x00, 0xf4, 0x50, 0x61, 0xa0, 0xc3, 0xaf, 0xf8, 0xcc, 0xc9, 0xe3,
	0x04, 0xaf, 0xec, 0x9c, 0x48, 0xeb, 0xa4, 0x09, 0xe3, 0x24, 0x23, 0xab, 0x00, 0x15, 0x71, 0x01,
	0x5a, 0xf9, 0xa9, 0xfa, 0x92, 0xa0, 0xe4, 0xc0, 0x6a, 0x23, 0x54, 0x04, 0x46, 0xa3, 0x98, 0x32,
	0xd5, 0x20, 0x2f, 0x6d, 0x06, 0x5d, 0x07, 0xfa, 0x8b, 0xb2, 0x67, 0x24, 0xdf, 0xa4, 0x35, 0x34,
	0x1c, 0x24, 0xbb, 0x48, 0x33, 0x91, 0xeb, 0x47, 0xd2, 0x10, 0xf8, 0x6a, 0x0c, 0xa5, 0xb6, 0xf1,
	0x66, 0x33, 0x59, 0x46, 0x36, 0xc4, 0x17, 0xfa, 0xac, 0x47, 0x41, 0xfe, 0xb1, 0x80, 0x92, 0x1b,
	0x8d, 0xa8, 0x0d, 0xa1, 0x05, 0x1e, 0x3d, 0xbe, 0x16, 0xb5, 0xad, 0xe9, 0xd0, 0x80, 0xa0, 0xcf,
	0x95, 0x19, 0xd1, 0xa1, 0x9b, 0xb2, 0x0e, 0x1d, 0x08, 0x4c, 0x0a, 0xb2, 0xf3, 0x6a, 0xa5, 0x5d,
	0x64, 0x1b, 0xb5, 0x61, 0xb0, 0xf3, 0xcc, 0xb3, 0xc6, 0xad, 0xe8, 0x8b, 0xcd, 0xc2, 0xd1, 0x0a,
	0x3a, 0xa4, 0x59, 0x95, 0x56, 0x35, 0xcf, 0xa6, 0x6f, 0x85, 0x3b, 0x01, 0xf9, 0x58, 0x8d, 0x84,
	0xea, 0x0e, 0x85, 0x3d, 0x62, 0x3e, 0x9c, 0x9a, 0x19, 0xb9, 0x38, 0x40, 0x7f, 0x35, 0x76, 0x81,
	0xf9, 0xd1, 0x26, 0x13, 0x4a, 0x3b, 0x54, 0x75, 0x96, 0xc3, 0x22, 0x37, 0x2e, 0xf4, 0xcf, 0xa5,
	0x9d, 0xfd, 0x12, 0x09, 0xe9, 0x2b, 0x38, 0x5f, 0x1a, 0xb7, 0x80, 0xf3, 0xe1, 0x26, 0x07, 0xfa,
	0x5c, 0x78, 0xfd, 0x54, 0x1e, 0x36, 0xed, 0xeb, 0x5d, 0x47, 0xd8, 0x50, 0xf0, 0x66, 0x85, 0x58,
	0x36, 0xdd, 0xab, 0x2e, 0x50, 0xc0, 0x5c, 0xf9, 0xed, 0x46, 0x1b, 0x7c, 0xe2, 0x58, 0x40, 0xd5,
	0x68, 0x3b, 0x0e, 0xa3, 0xc1, 0x66, 0xd9, 0x6f, 0x63, 0x4e, 0x1c, 0x02, 0x94, 0xd8, 0x66, 0x0c,
	0xae, 0x7a, 0xa7, 0x69, 0x8b, 0x78, 0x67, 0xa2, 0x10, 0x9a, 0xc8, 0xec, 0x9f, 0xa6, 0x2d, 0x9e,
	0xbd, 0x15, 0x35, 0x70, 0xaa, 0x47, 0x3d, 0x57, 0x12, 0x0d, 0x4e, 0x48, 0x39, 0xf6, 0x98, 0x2d,
	0xee, 0x11, 0x30, 0x43, 0x96, 0xb2, 0x70, 0xaf, 0xc1, 0xa0, 0xe7, 0x38, 0xab, 0x53, 0xf3, 0x12,
	0xf2, 0x46, 0x65, 0x6f, 0xbb, 0xfd, 0x4e, 0xc5, 0x88, 0xda, 0xde, 0xc6, 0x07, 0x9d, 0x8e, 0xbd,
	0x92, 0xdb, 0x2b, 0x1a, 0xfd, 0x00, 0x8b, 0x9b, 0x30, 0x40, 0x7b, 0xd6, 0x7a, 0xba, 0x3c, 0x31,
	0x3e, 0x06, 0xf0, 0x13, 0x19, 0x58, 0x74, 0xfd, 0x45, 0x2c, 0xee, 0xc5, 0xe0, 0x23, 0x59, 0x12,
	0x75, 0x93, 0x97, 0xfc, 0x86, 0x21, 0x24, 0x01, 0x24, 0x0b, 0x14, 0x4d, 0xd3, 0x77, 0x03, 0xa8,
	0x3a, 0xa2, 0xdf, 0x1c, 0x2a, 0x45, 0x23, 0xe4, 0x05, 0xa4, 0x8b, 0x9d, 0x4d, 0xf3, 0x79, 0x11,
	0xd7, 0xa9, 0x68, 0x39, 0xff, 0x26, 0x8f, 0xf3, 0x6f, 0x25, 0x68, 0x37, 0x46, 0x8d, 0xae, 0x71,
	0x6b, 0xd9, 0x0c, 0x54, 0x39, 0x0c, 0x54, 0x7d, 0x97, 0x22, 0x9b, 0x26, 0xdb, 0xdf, 0x79, 0x29,
	0xb2, 0x36, 0xf9, 0x53, 0x22, 0xbb, 0xd5, 0x8e, 0xea, 0xfc, 0x35, 0xcf, 0xf4, 0xbd, 0xe3, 0xf6,
	0x42, 0xfb, 0x78, 0x75, 0xa1, 0x7d, 0x4c, 0xb2, 0xc5, 0x7c, 0x12, 0xcf, 0x16, 0xdc, 0x2f, 0xc4,
	0xba, 0xea, 0xf6, 0x7f, 0xff, 0x75, 0x73, 0xf7, 0xc6, 0xe1, 0x07, 0x83, 0x9b, 0xff, 0xbc, 0xf1,
	0xde, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x0d, 0xf2, 0x2b, 0x20, 0x17, 0x00, 0x00,
}
