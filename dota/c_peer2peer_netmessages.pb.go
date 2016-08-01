// Code generated by protoc-gen-go.
// source: c_peer2peer_netmessages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type P2P_Messages int32

const (
	P2P_Messages_p2p_TextMessage          P2P_Messages = 256
	P2P_Messages_p2p_Voice                P2P_Messages = 257
	P2P_Messages_p2p_Ping                 P2P_Messages = 258
	P2P_Messages_p2p_VRAvatarPosition     P2P_Messages = 259
	P2P_Messages_p2p_WatchSynchronization P2P_Messages = 260
)

var P2P_Messages_name = map[int32]string{
	256: "p2p_TextMessage",
	257: "p2p_Voice",
	258: "p2p_Ping",
	259: "p2p_VRAvatarPosition",
	260: "p2p_WatchSynchronization",
}
var P2P_Messages_value = map[string]int32{
	"p2p_TextMessage":          256,
	"p2p_Voice":                257,
	"p2p_Ping":                 258,
	"p2p_VRAvatarPosition":     259,
	"p2p_WatchSynchronization": 260,
}

func (x P2P_Messages) Enum() *P2P_Messages {
	p := new(P2P_Messages)
	*p = x
	return p
}
func (x P2P_Messages) String() string {
	return proto.EnumName(P2P_Messages_name, int32(x))
}
func (x *P2P_Messages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(P2P_Messages_value, data, "P2P_Messages")
	if err != nil {
		return err
	}
	*x = P2P_Messages(value)
	return nil
}
func (P2P_Messages) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type CP2P_Voice_Handler_Flags int32

const (
	CP2P_Voice_Played_Audio CP2P_Voice_Handler_Flags = 1
)

var CP2P_Voice_Handler_Flags_name = map[int32]string{
	1: "Played_Audio",
}
var CP2P_Voice_Handler_Flags_value = map[string]int32{
	"Played_Audio": 1,
}

func (x CP2P_Voice_Handler_Flags) Enum() *CP2P_Voice_Handler_Flags {
	p := new(CP2P_Voice_Handler_Flags)
	*p = x
	return p
}
func (x CP2P_Voice_Handler_Flags) String() string {
	return proto.EnumName(CP2P_Voice_Handler_Flags_name, int32(x))
}
func (x *CP2P_Voice_Handler_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CP2P_Voice_Handler_Flags_value, data, "CP2P_Voice_Handler_Flags")
	if err != nil {
		return err
	}
	*x = CP2P_Voice_Handler_Flags(value)
	return nil
}
func (CP2P_Voice_Handler_Flags) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

type CP2P_TextMessage struct {
	Text             []byte `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CP2P_TextMessage) Reset()                    { *m = CP2P_TextMessage{} }
func (m *CP2P_TextMessage) String() string            { return proto.CompactTextString(m) }
func (*CP2P_TextMessage) ProtoMessage()               {}
func (*CP2P_TextMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *CP2P_TextMessage) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

type CSteam_Voice_Encoding struct {
	VoiceData        []byte `protobuf:"bytes,1,opt,name=voice_data" json:"voice_data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSteam_Voice_Encoding) Reset()                    { *m = CSteam_Voice_Encoding{} }
func (m *CSteam_Voice_Encoding) String() string            { return proto.CompactTextString(m) }
func (*CSteam_Voice_Encoding) ProtoMessage()               {}
func (*CSteam_Voice_Encoding) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *CSteam_Voice_Encoding) GetVoiceData() []byte {
	if m != nil {
		return m.VoiceData
	}
	return nil
}

type CP2P_Voice struct {
	Audio            *CMsgVoiceAudio `protobuf:"bytes,1,opt,name=audio" json:"audio,omitempty"`
	BroadcastGroup   *uint32         `protobuf:"varint,2,opt,name=broadcast_group" json:"broadcast_group,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *CP2P_Voice) Reset()                    { *m = CP2P_Voice{} }
func (m *CP2P_Voice) String() string            { return proto.CompactTextString(m) }
func (*CP2P_Voice) ProtoMessage()               {}
func (*CP2P_Voice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *CP2P_Voice) GetAudio() *CMsgVoiceAudio {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *CP2P_Voice) GetBroadcastGroup() uint32 {
	if m != nil && m.BroadcastGroup != nil {
		return *m.BroadcastGroup
	}
	return 0
}

type CP2P_Ping struct {
	SendTime         *uint64 `protobuf:"varint,1,req,name=send_time" json:"send_time,omitempty"`
	IsReply          *bool   `protobuf:"varint,2,req,name=is_reply" json:"is_reply,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CP2P_Ping) Reset()                    { *m = CP2P_Ping{} }
func (m *CP2P_Ping) String() string            { return proto.CompactTextString(m) }
func (*CP2P_Ping) ProtoMessage()               {}
func (*CP2P_Ping) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *CP2P_Ping) GetSendTime() uint64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *CP2P_Ping) GetIsReply() bool {
	if m != nil && m.IsReply != nil {
		return *m.IsReply
	}
	return false
}

type CP2P_VRAvatarPosition struct {
	BodyParts        []*CP2P_VRAvatarPosition_COrientation `protobuf:"bytes,1,rep,name=body_parts" json:"body_parts,omitempty"`
	HatId            *int32                                `protobuf:"varint,2,opt,name=hat_id" json:"hat_id,omitempty"`
	SceneId          *int32                                `protobuf:"varint,3,opt,name=scene_id" json:"scene_id,omitempty"`
	WorldScale       *int32                                `protobuf:"varint,4,opt,name=world_scale" json:"world_scale,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *CP2P_VRAvatarPosition) Reset()                    { *m = CP2P_VRAvatarPosition{} }
func (m *CP2P_VRAvatarPosition) String() string            { return proto.CompactTextString(m) }
func (*CP2P_VRAvatarPosition) ProtoMessage()               {}
func (*CP2P_VRAvatarPosition) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *CP2P_VRAvatarPosition) GetBodyParts() []*CP2P_VRAvatarPosition_COrientation {
	if m != nil {
		return m.BodyParts
	}
	return nil
}

func (m *CP2P_VRAvatarPosition) GetHatId() int32 {
	if m != nil && m.HatId != nil {
		return *m.HatId
	}
	return 0
}

func (m *CP2P_VRAvatarPosition) GetSceneId() int32 {
	if m != nil && m.SceneId != nil {
		return *m.SceneId
	}
	return 0
}

func (m *CP2P_VRAvatarPosition) GetWorldScale() int32 {
	if m != nil && m.WorldScale != nil {
		return *m.WorldScale
	}
	return 0
}

type CP2P_VRAvatarPosition_COrientation struct {
	Pos              *CMsgVector `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	Ang              *CMsgQAngle `protobuf:"bytes,2,opt,name=ang" json:"ang,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CP2P_VRAvatarPosition_COrientation) Reset()         { *m = CP2P_VRAvatarPosition_COrientation{} }
func (m *CP2P_VRAvatarPosition_COrientation) String() string { return proto.CompactTextString(m) }
func (*CP2P_VRAvatarPosition_COrientation) ProtoMessage()    {}
func (*CP2P_VRAvatarPosition_COrientation) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{4, 0}
}

func (m *CP2P_VRAvatarPosition_COrientation) GetPos() *CMsgVector {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *CP2P_VRAvatarPosition_COrientation) GetAng() *CMsgQAngle {
	if m != nil {
		return m.Ang
	}
	return nil
}

type CP2P_WatchSynchronization struct {
	DemoTick                         *int32 `protobuf:"varint,1,opt,name=demo_tick" json:"demo_tick,omitempty"`
	Paused                           *bool  `protobuf:"varint,2,opt,name=paused" json:"paused,omitempty"`
	TvListenVoiceIndices             *int32 `protobuf:"varint,3,opt,name=tv_listen_voice_indices" json:"tv_listen_voice_indices,omitempty"`
	DotaSpectatorMode                *int32 `protobuf:"varint,4,opt,name=dota_spectator_mode" json:"dota_spectator_mode,omitempty"`
	DotaSpectatorWatchingBroadcaster *int32 `protobuf:"varint,5,opt,name=dota_spectator_watching_broadcaster" json:"dota_spectator_watching_broadcaster,omitempty"`
	DotaSpectatorHeroIndex           *int32 `protobuf:"varint,6,opt,name=dota_spectator_hero_index" json:"dota_spectator_hero_index,omitempty"`
	DotaSpectatorAutospeed           *int32 `protobuf:"varint,7,opt,name=dota_spectator_autospeed" json:"dota_spectator_autospeed,omitempty"`
	DotaReplaySpeed                  *int32 `protobuf:"varint,8,opt,name=dota_replay_speed" json:"dota_replay_speed,omitempty"`
	XXX_unrecognized                 []byte `json:"-"`
}

func (m *CP2P_WatchSynchronization) Reset()                    { *m = CP2P_WatchSynchronization{} }
func (m *CP2P_WatchSynchronization) String() string            { return proto.CompactTextString(m) }
func (*CP2P_WatchSynchronization) ProtoMessage()               {}
func (*CP2P_WatchSynchronization) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *CP2P_WatchSynchronization) GetDemoTick() int32 {
	if m != nil && m.DemoTick != nil {
		return *m.DemoTick
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetPaused() bool {
	if m != nil && m.Paused != nil {
		return *m.Paused
	}
	return false
}

func (m *CP2P_WatchSynchronization) GetTvListenVoiceIndices() int32 {
	if m != nil && m.TvListenVoiceIndices != nil {
		return *m.TvListenVoiceIndices
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorMode() int32 {
	if m != nil && m.DotaSpectatorMode != nil {
		return *m.DotaSpectatorMode
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorWatchingBroadcaster() int32 {
	if m != nil && m.DotaSpectatorWatchingBroadcaster != nil {
		return *m.DotaSpectatorWatchingBroadcaster
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorHeroIndex() int32 {
	if m != nil && m.DotaSpectatorHeroIndex != nil {
		return *m.DotaSpectatorHeroIndex
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaSpectatorAutospeed() int32 {
	if m != nil && m.DotaSpectatorAutospeed != nil {
		return *m.DotaSpectatorAutospeed
	}
	return 0
}

func (m *CP2P_WatchSynchronization) GetDotaReplaySpeed() int32 {
	if m != nil && m.DotaReplaySpeed != nil {
		return *m.DotaReplaySpeed
	}
	return 0
}

func init() {
	proto.RegisterType((*CP2P_TextMessage)(nil), "dota.CP2P_TextMessage")
	proto.RegisterType((*CSteam_Voice_Encoding)(nil), "dota.CSteam_Voice_Encoding")
	proto.RegisterType((*CP2P_Voice)(nil), "dota.CP2P_Voice")
	proto.RegisterType((*CP2P_Ping)(nil), "dota.CP2P_Ping")
	proto.RegisterType((*CP2P_VRAvatarPosition)(nil), "dota.CP2P_VRAvatarPosition")
	proto.RegisterType((*CP2P_VRAvatarPosition_COrientation)(nil), "dota.CP2P_VRAvatarPosition.COrientation")
	proto.RegisterType((*CP2P_WatchSynchronization)(nil), "dota.CP2P_WatchSynchronization")
	proto.RegisterEnum("dota.P2P_Messages", P2P_Messages_name, P2P_Messages_value)
	proto.RegisterEnum("dota.CP2P_Voice_Handler_Flags", CP2P_Voice_Handler_Flags_name, CP2P_Voice_Handler_Flags_value)
}

func init() { proto.RegisterFile("c_peer2peer_netmessages.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0xfd, 0x65, 0x93, 0xf4, 0x97, 0x4e, 0x53, 0xba, 0x75, 0x0b, 0xdd, 0x16, 0x55, 0xb4, 0xe9,
	0x25, 0xa2, 0x52, 0x84, 0x72, 0xe6, 0x52, 0x22, 0x10, 0x07, 0x2a, 0x02, 0x45, 0x70, 0xb4, 0xdc,
	0xf5, 0x68, 0x63, 0x75, 0x63, 0xaf, 0x6c, 0xa7, 0x6d, 0x90, 0x90, 0xca, 0x9f, 0x0b, 0x5f, 0x91,
	0x4f, 0xc3, 0xd8, 0x59, 0x50, 0x1b, 0x7a, 0xd9, 0x95, 0xe7, 0xbd, 0x99, 0x79, 0x33, 0x6f, 0x60,
	0x3f, 0xe7, 0x15, 0xa2, 0x1d, 0x86, 0x0f, 0xd7, 0xe8, 0xa7, 0xe8, 0x9c, 0x28, 0xd0, 0x0d, 0x2a,
	0x6b, 0xbc, 0x61, 0x2d, 0x69, 0xbc, 0xd8, 0xdb, 0xfc, 0x07, 0xd8, 0x7b, 0x44, 0xa1, 0x2b, 0x63,
	0x2f, 0xce, 0x85, 0x43, 0x3f, 0xaf, 0xfe, 0xc4, 0x7b, 0x07, 0x90, 0x8e, 0xc6, 0xc3, 0x31, 0xff,
	0x80, 0xd7, 0xfe, 0x74, 0x91, 0xc2, 0xba, 0xd0, 0xf2, 0xf4, 0xcc, 0x1a, 0x07, 0x8d, 0x7e, 0xb7,
	0x77, 0x0c, 0x0f, 0x47, 0x67, 0x1e, 0xc5, 0x94, 0x7f, 0x34, 0x2a, 0x47, 0xfe, 0x52, 0xe7, 0x46,
	0x2a, 0x5d, 0x30, 0x06, 0x70, 0x19, 0x23, 0x52, 0x78, 0x51, 0x93, 0xa7, 0x00, 0xb1, 0x5c, 0xa4,
	0xb2, 0x23, 0x68, 0x8b, 0x99, 0x54, 0x26, 0x82, 0x6b, 0xc3, 0xed, 0x41, 0x50, 0x37, 0x18, 0x9d,
	0xba, 0x22, 0xe2, 0x27, 0x01, 0x63, 0x3b, 0xb0, 0x71, 0x6e, 0x8d, 0x90, 0xb9, 0x70, 0x9e, 0x17,
	0xd6, 0xcc, 0xaa, 0x2c, 0x21, 0xfa, 0x7a, 0xef, 0x10, 0xd6, 0x5f, 0x0b, 0x2d, 0x4b, 0x1a, 0xf4,
	0x55, 0x29, 0x0a, 0xc7, 0x52, 0xe8, 0x8e, 0x4b, 0x31, 0x47, 0xc9, 0x63, 0x66, 0xda, 0xe8, 0x3d,
	0x83, 0xd5, 0xd8, 0x6e, 0x1c, 0xf4, 0x6c, 0xc2, 0xaa, 0x43, 0x2d, 0xb9, 0x57, 0x53, 0xa4, 0x8e,
	0x49, 0xbf, 0x45, 0x19, 0x1d, 0xe5, 0xb8, 0xc5, 0xaa, 0x9c, 0x53, 0xd1, 0xa4, 0xdf, 0xe9, 0xfd,
	0x6a, 0xd0, 0x38, 0x51, 0xe1, 0xfb, 0x93, 0x4b, 0xd2, 0x6d, 0xc7, 0xc6, 0x29, 0xaf, 0x8c, 0x66,
	0xcf, 0x01, 0xce, 0x8d, 0x9c, 0xf3, 0x4a, 0x58, 0xef, 0x28, 0xbf, 0x49, 0x8a, 0xfb, 0xb5, 0xe2,
	0xfb, 0x12, 0x06, 0xa3, 0xb7, 0x56, 0xa1, 0xf6, 0x22, 0x66, 0x3f, 0x80, 0x95, 0x89, 0xf0, 0x5c,
	0xc9, 0x28, 0xbe, 0x1d, 0x3a, 0xbb, 0x1c, 0x35, 0x86, 0x48, 0x33, 0x46, 0xb6, 0x60, 0x8d, 0x0c,
	0x28, 0x25, 0x77, 0xb9, 0x28, 0x31, 0x6b, 0x85, 0xe0, 0xde, 0x1b, 0xe8, 0xde, 0x29, 0xb3, 0x0f,
	0xcd, 0xca, 0xb8, 0x7a, 0x5f, 0xe9, 0xad, 0x7d, 0x61, 0xee, 0x8d, 0x0d, 0xb0, 0xd0, 0x45, 0x6c,
	0x71, 0x07, 0x7e, 0x77, 0xa2, 0x8b, 0x12, 0x7b, 0x3f, 0x13, 0xd8, 0x8d, 0x5a, 0x3f, 0x09, 0x9f,
	0x4f, 0xce, 0xe6, 0x3a, 0x9f, 0x58, 0xa3, 0xd5, 0xe7, 0x45, 0x6d, 0xda, 0x8f, 0xc4, 0xa9, 0xa1,
	0xfd, 0xe4, 0x17, 0xb1, 0x43, 0x3b, 0xa8, 0xae, 0xc4, 0xcc, 0xe1, 0x42, 0x75, 0x87, 0x3d, 0x81,
	0x1d, 0x7f, 0xc9, 0x4b, 0xe5, 0x3c, 0x6a, 0xbe, 0x30, 0x57, 0x69, 0x49, 0x3f, 0x57, 0x0f, 0xf1,
	0x18, 0xb6, 0x42, 0x53, 0xee, 0x2a, 0x12, 0x24, 0x48, 0x12, 0x9f, 0x1a, 0x59, 0x0f, 0xc3, 0x8e,
	0xe1, 0x68, 0x09, 0xbc, 0x0a, 0x3a, 0xc8, 0x1b, 0xfe, 0xd7, 0x61, 0xb4, 0x59, 0x3b, 0x92, 0x0f,
	0x61, 0x77, 0x89, 0x3c, 0x41, 0x6b, 0x42, 0x3b, 0xbc, 0xce, 0x56, 0x22, 0xe5, 0x00, 0xb2, 0x25,
	0x8a, 0x98, 0x79, 0x43, 0x2f, 0xd2, 0xfb, 0x7f, 0x64, 0xec, 0xc2, 0x66, 0x64, 0x04, 0x87, 0xc5,
	0x9c, 0x2f, 0xa0, 0x4e, 0x80, 0x9e, 0x7e, 0xa1, 0x63, 0xa1, 0x4d, 0xd4, 0x37, 0xed, 0xd8, 0x36,
	0x6c, 0x54, 0xc3, 0xea, 0xf6, 0x9d, 0xa7, 0x37, 0x09, 0x2d, 0x60, 0x35, 0x44, 0xe3, 0x39, 0xa6,
	0x5f, 0x13, 0xb6, 0x0e, 0x9d, 0xf0, 0x0e, 0xf7, 0x94, 0x7e, 0x4b, 0xa8, 0xfe, 0x76, 0x84, 0x97,
	0xac, 0x4f, 0xbf, 0x27, 0x64, 0x45, 0x16, 0xa0, 0xfb, 0x36, 0x9d, 0xfe, 0x48, 0x5e, 0x34, 0x6f,
	0x1a, 0xff, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x05, 0x02, 0xb3, 0xad, 0x03, 0x00, 0x00,
}
