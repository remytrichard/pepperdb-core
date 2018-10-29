// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package dappserverpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	DAppServerConfig
*/
package dappserverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Dappserver *DAppServerConfig `protobuf:"bytes,1,opt,name=dappserver" json:"dappserver"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetDappserver() *DAppServerConfig {
	if m != nil {
		return m.Dappserver
	}
	return nil
}

type DAppServerConfig struct {
	// Enable
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable"`
	// Host
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host"`
	// Port
	Port uint64 `protobuf:"varint,3,opt,name=port,proto3" json:"port"`
	// ReadTimeout in millisecond
	ReadTimeoutMs uint64 `protobuf:"varint,4,opt,name=read_timeout_ms,json=readTimeoutMs,proto3" json:"read_timeout_ms"`
	// WriteTimeout in millisecond
	WriteTimeoutMs uint64 `protobuf:"varint,5,opt,name=write_timeout_ms,json=writeTimeoutMs,proto3" json:"write_timeout_ms"`
	// Log level
	LogLevel string `protobuf:"bytes,6,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	// Log file name
	LogFile string `protobuf:"bytes,7,opt,name=log_file,json=logFile,proto3" json:"log_file"`
	// Log file age, unit is s.
	LogAge uint32 `protobuf:"varint,8,opt,name=log_age,json=logAge,proto3" json:"log_age"`
}

func (m *DAppServerConfig) Reset()                    { *m = DAppServerConfig{} }
func (m *DAppServerConfig) String() string            { return proto.CompactTextString(m) }
func (*DAppServerConfig) ProtoMessage()               {}
func (*DAppServerConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *DAppServerConfig) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *DAppServerConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DAppServerConfig) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DAppServerConfig) GetReadTimeoutMs() uint64 {
	if m != nil {
		return m.ReadTimeoutMs
	}
	return 0
}

func (m *DAppServerConfig) GetWriteTimeoutMs() uint64 {
	if m != nil {
		return m.WriteTimeoutMs
	}
	return 0
}

func (m *DAppServerConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *DAppServerConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *DAppServerConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "dappserverpb.Config")
	proto.RegisterType((*DAppServerConfig)(nil), "dappserverpb.DAppServerConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xd6, 0xae, 0x7b, 0x6e, 0x3a, 0x72, 0xd0, 0x88, 0x20, 0x65, 0x07, 0xe9, 0xa9,
	0x07, 0xbd, 0x0b, 0x43, 0x11, 0x0f, 0x7a, 0xa9, 0xde, 0x43, 0xeb, 0xde, 0x62, 0x20, 0xdb, 0x0b,
	0x69, 0x9c, 0x7f, 0xb9, 0x77, 0xc9, 0x9b, 0xb0, 0xb2, 0xdb, 0xf7, 0xc7, 0xe7, 0x9b, 0xc0, 0x83,
	0xc9, 0x27, 0x6d, 0x56, 0xd6, 0xd4, 0x3e, 0x50, 0x24, 0x39, 0x59, 0xb6, 0xde, 0xf7, 0x18, 0xb6,
	0x18, 0x7c, 0x37, 0x7f, 0x81, 0xfc, 0x91, 0x5b, 0xf9, 0x00, 0xb0, 0x6f, 0x94, 0x28, 0x45, 0x75,
	0x7a, 0x77, 0x53, 0x0f, 0xe1, 0xfa, 0x69, 0xe1, 0xfd, 0x3b, 0x9b, 0xdd, 0xa6, 0x19, 0x2c, 0xe6,
	0xbf, 0x02, 0x66, 0x87, 0x80, 0xbc, 0x80, 0x1c, 0x37, 0x6d, 0xe7, 0x90, 0x1f, 0x2c, 0x9a, 0x7f,
	0x27, 0x25, 0x64, 0x5f, 0xd4, 0x47, 0x75, 0x54, 0x8a, 0x6a, 0xdc, 0xb0, 0x4e, 0x99, 0xa7, 0x10,
	0xd5, 0x71, 0x29, 0xaa, 0xac, 0x61, 0x2d, 0x6f, 0xe1, 0x3c, 0x60, 0xbb, 0xd4, 0xd1, 0xae, 0x91,
	0xbe, 0xa3, 0x5e, 0xf7, 0x2a, 0xe3, 0x7a, 0x9a, 0xe2, 0x8f, 0x5d, 0xfa, 0xd6, 0xcb, 0x0a, 0x66,
	0x3f, 0xc1, 0x46, 0x1c, 0x82, 0x27, 0x0c, 0x9e, 0x71, 0xbe, 0x27, 0xaf, 0x61, 0xec, 0xc8, 0x68,
	0x87, 0x5b, 0x74, 0x2a, 0xe7, 0xef, 0x0b, 0x47, 0xe6, 0x35, 0x79, 0x79, 0x05, 0x49, 0xeb, 0x95,
	0x75, 0xa8, 0x46, 0xdc, 0x8d, 0x1c, 0x99, 0x67, 0xeb, 0x50, 0x5e, 0x42, 0x92, 0xba, 0x35, 0xa8,
	0x8a, 0x52, 0x54, 0xd3, 0x26, 0x77, 0x64, 0x16, 0x06, 0xbb, 0x9c, 0xcf, 0x7a, 0xff, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x01, 0xa0, 0x12, 0x66, 0x01, 0x00, 0x00,
}
