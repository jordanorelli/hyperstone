// Code generated by protoc-gen-go.
// source: connectionless_netmessages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type C2S_CONNECT_Message struct {
	HostVersion       *uint32                       `protobuf:"varint,1,opt,name=host_version" json:"host_version,omitempty"`
	AuthProtocol      *uint32                       `protobuf:"varint,2,opt,name=auth_protocol" json:"auth_protocol,omitempty"`
	ChallengeNumber   *uint32                       `protobuf:"varint,3,opt,name=challenge_number" json:"challenge_number,omitempty"`
	ReservationCookie *uint64                       `protobuf:"fixed64,4,opt,name=reservation_cookie" json:"reservation_cookie,omitempty"`
	LowViolence       *bool                         `protobuf:"varint,5,opt,name=low_violence" json:"low_violence,omitempty"`
	EncryptedPassword []byte                        `protobuf:"bytes,6,opt,name=encrypted_password" json:"encrypted_password,omitempty"`
	Splitplayers      []*CCLCMsg_SplitPlayerConnect `protobuf:"bytes,7,rep,name=splitplayers" json:"splitplayers,omitempty"`
	AuthSteam         []byte                        `protobuf:"bytes,8,opt,name=auth_steam" json:"auth_steam,omitempty"`
	ChallengeContext  *string                       `protobuf:"bytes,9,opt,name=challenge_context" json:"challenge_context,omitempty"`
	XXX_unrecognized  []byte                        `json:"-"`
}

func (m *C2S_CONNECT_Message) Reset()                    { *m = C2S_CONNECT_Message{} }
func (m *C2S_CONNECT_Message) String() string            { return proto.CompactTextString(m) }
func (*C2S_CONNECT_Message) ProtoMessage()               {}
func (*C2S_CONNECT_Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *C2S_CONNECT_Message) GetHostVersion() uint32 {
	if m != nil && m.HostVersion != nil {
		return *m.HostVersion
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetAuthProtocol() uint32 {
	if m != nil && m.AuthProtocol != nil {
		return *m.AuthProtocol
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetChallengeNumber() uint32 {
	if m != nil && m.ChallengeNumber != nil {
		return *m.ChallengeNumber
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetReservationCookie() uint64 {
	if m != nil && m.ReservationCookie != nil {
		return *m.ReservationCookie
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetLowViolence() bool {
	if m != nil && m.LowViolence != nil {
		return *m.LowViolence
	}
	return false
}

func (m *C2S_CONNECT_Message) GetEncryptedPassword() []byte {
	if m != nil {
		return m.EncryptedPassword
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetSplitplayers() []*CCLCMsg_SplitPlayerConnect {
	if m != nil {
		return m.Splitplayers
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetAuthSteam() []byte {
	if m != nil {
		return m.AuthSteam
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetChallengeContext() string {
	if m != nil && m.ChallengeContext != nil {
		return *m.ChallengeContext
	}
	return ""
}

type C2S_CONNECTION_Message struct {
	AddonName        *string `protobuf:"bytes,1,opt,name=addon_name" json:"addon_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2S_CONNECTION_Message) Reset()                    { *m = C2S_CONNECTION_Message{} }
func (m *C2S_CONNECTION_Message) String() string            { return proto.CompactTextString(m) }
func (*C2S_CONNECTION_Message) ProtoMessage()               {}
func (*C2S_CONNECTION_Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *C2S_CONNECTION_Message) GetAddonName() string {
	if m != nil && m.AddonName != nil {
		return *m.AddonName
	}
	return ""
}

func init() {
	proto.RegisterType((*C2S_CONNECT_Message)(nil), "dota.C2S_CONNECT_Message")
	proto.RegisterType((*C2S_CONNECTION_Message)(nil), "dota.C2S_CONNECTION_Message")
}

func init() { proto.RegisterFile("connectionless_netmessages.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0xe5, 0x47, 0x84, 0x0a, 0x89, 0xd4, 0x9f, 0x54, 0x56, 0x13, 0x56, 0x2c, 0x0c, 0x0b,
	0x16, 0x3e, 0x80, 0x8d, 0x0b, 0x13, 0x01, 0x13, 0xdc, 0x37, 0xb5, 0x73, 0xc3, 0x4c, 0xec, 0xf4,
	0x4e, 0xda, 0x32, 0xc8, 0xce, 0x57, 0xf0, 0x8d, 0x6d, 0x4b, 0xa2, 0x26, 0xae, 0x26, 0x73, 0xbf,
	0x7b, 0x4e, 0xcf, 0x3d, 0x24, 0x53, 0x68, 0x0c, 0x28, 0x5f, 0xa2, 0xd1, 0xe0, 0x9c, 0x30, 0xe0,
	0xab, 0xf0, 0x95, 0x5b, 0x70, 0xf3, 0xda, 0xa2, 0x47, 0xda, 0xcd, 0xd1, 0xcb, 0xc9, 0xf8, 0x1f,
	0x98, 0x7e, 0xb5, 0xc9, 0x25, 0x5f, 0x6c, 0x04, 0x5f, 0xaf, 0x56, 0x8f, 0xfc, 0x55, 0x2c, 0x8f,
	0x98, 0x5e, 0x91, 0x61, 0x81, 0xce, 0x8b, 0x06, 0xac, 0x0b, 0xb6, 0xac, 0x95, 0xb5, 0x66, 0x23,
	0x7a, 0x4d, 0x46, 0x72, 0xe7, 0x0b, 0x91, 0xb4, 0x0a, 0x35, 0x6b, 0xa7, 0x31, 0x23, 0x17, 0xaa,
	0x90, 0x5a, 0x83, 0xd9, 0x82, 0x30, 0xbb, 0xea, 0x0d, 0x2c, 0xeb, 0x24, 0x32, 0x21, 0xd4, 0x82,
	0x03, 0xdb, 0xc8, 0x18, 0x4e, 0x28, 0xc4, 0xf7, 0x12, 0x58, 0x37, 0xb0, 0x5e, 0x7c, 0x42, 0xe3,
	0x5e, 0x34, 0x25, 0x06, 0xa1, 0x02, 0x76, 0x1a, 0xa6, 0xfd, 0xa8, 0x08, 0x7f, 0xf6, 0x50, 0x7b,
	0xc8, 0x45, 0x2d, 0x9d, 0xdb, 0xa3, 0xcd, 0x59, 0x2f, 0xb0, 0x21, 0xbd, 0x27, 0x43, 0x57, 0xeb,
	0xd2, 0xd7, 0x5a, 0x1e, 0x42, 0x2e, 0x76, 0x96, 0x75, 0x66, 0xe7, 0x8b, 0x6c, 0x1e, 0x8f, 0x9b,
	0x73, 0xfe, 0xcc, 0x97, 0x6e, 0x2b, 0x36, 0x71, 0xe3, 0x25, 0x6d, 0xf0, 0x63, 0x2d, 0x94, 0x12,
	0x92, 0x62, 0x3b, 0x0f, 0xb2, 0x62, 0xfd, 0xe4, 0x75, 0x4b, 0xc6, 0xbf, 0x99, 0x43, 0x7f, 0x1e,
	0x3e, 0x3c, 0x1b, 0x04, 0x34, 0x98, 0xde, 0x91, 0x9b, 0x3f, 0x95, 0x3c, 0xad, 0x57, 0x3f, 0xad,
	0x44, 0xa3, 0x3c, 0x0f, 0x87, 0x18, 0x59, 0x41, 0xea, 0x64, 0xf0, 0xd0, 0xf9, 0x6c, 0x9d, 0x7c,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x22, 0xbd, 0xbd, 0x23, 0x82, 0x01, 0x00, 0x00,
}
