// Code generated by protoc-gen-go.
// source: timeformat.proto
// DO NOT EDIT!

package model

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TimeStruct struct {
	Year             *int32  `protobuf:"varint,1,req,name=year" json:"year,omitempty"`
	Month            *int32  `protobuf:"varint,2,req,name=month" json:"month,omitempty"`
	Day              *int32  `protobuf:"varint,3,req,name=day" json:"day,omitempty"`
	Hour             *int32  `protobuf:"varint,4,req,name=hour" json:"hour,omitempty"`
	Minitue          *int32  `protobuf:"varint,5,req,name=minitue" json:"minitue,omitempty"`
	Second           *int32  `protobuf:"varint,6,req,name=second" json:"second,omitempty"`
	Nsec             *int32  `protobuf:"varint,7,opt,name=nsec" json:"nsec,omitempty"`
	Format           *string `protobuf:"bytes,8,opt,name=format" json:"format,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TimeStruct) Reset()         { *m = TimeStruct{} }
func (m *TimeStruct) String() string { return proto.CompactTextString(m) }
func (*TimeStruct) ProtoMessage()    {}

func (m *TimeStruct) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *TimeStruct) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *TimeStruct) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

func (m *TimeStruct) GetHour() int32 {
	if m != nil && m.Hour != nil {
		return *m.Hour
	}
	return 0
}

func (m *TimeStruct) GetMinitue() int32 {
	if m != nil && m.Minitue != nil {
		return *m.Minitue
	}
	return 0
}

func (m *TimeStruct) GetSecond() int32 {
	if m != nil && m.Second != nil {
		return *m.Second
	}
	return 0
}

func (m *TimeStruct) GetNsec() int32 {
	if m != nil && m.Nsec != nil {
		return *m.Nsec
	}
	return 0
}

func (m *TimeStruct) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

func init() {
}