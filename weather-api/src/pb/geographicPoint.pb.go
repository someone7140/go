// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: src/proto/geographicPoint.proto

package pb

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddGeographicPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lat          float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon          float64 `protobuf:"fixed64,3,opt,name=lon,proto3" json:"lon,omitempty"`
	DisplayOrder int32   `protobuf:"varint,4,opt,name=displayOrder,proto3" json:"displayOrder,omitempty"`
}

func (x *AddGeographicPointRequest) Reset() {
	*x = AddGeographicPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGeographicPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGeographicPointRequest) ProtoMessage() {}

func (x *AddGeographicPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGeographicPointRequest.ProtoReflect.Descriptor instead.
func (*AddGeographicPointRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{0}
}

func (x *AddGeographicPointRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddGeographicPointRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *AddGeographicPointRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *AddGeographicPointRequest) GetDisplayOrder() int32 {
	if x != nil {
		return x.DisplayOrder
	}
	return 0
}

type UpdateGeographicPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lat          float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon          float64 `protobuf:"fixed64,4,opt,name=lon,proto3" json:"lon,omitempty"`
	DisplayOrder int32   `protobuf:"varint,5,opt,name=displayOrder,proto3" json:"displayOrder,omitempty"`
}

func (x *UpdateGeographicPointRequest) Reset() {
	*x = UpdateGeographicPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGeographicPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGeographicPointRequest) ProtoMessage() {}

func (x *UpdateGeographicPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGeographicPointRequest.ProtoReflect.Descriptor instead.
func (*UpdateGeographicPointRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGeographicPointRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGeographicPointRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGeographicPointRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *UpdateGeographicPointRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *UpdateGeographicPointRequest) GetDisplayOrder() int32 {
	if x != nil {
		return x.DisplayOrder
	}
	return 0
}

type DeleteGeographicPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGeographicPointRequest) Reset() {
	*x = DeleteGeographicPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGeographicPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGeographicPointRequest) ProtoMessage() {}

func (x *DeleteGeographicPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGeographicPointRequest.ProtoReflect.Descriptor instead.
func (*DeleteGeographicPointRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGeographicPointRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RegsiterGeographicPointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegsiterGeographicPointResponse) Reset() {
	*x = RegsiterGeographicPointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegsiterGeographicPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegsiterGeographicPointResponse) ProtoMessage() {}

func (x *RegsiterGeographicPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegsiterGeographicPointResponse.ProtoReflect.Descriptor instead.
func (*RegsiterGeographicPointResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{3}
}

type GetWeatherListByGeographicPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWeatherListByGeographicPointRequest) Reset() {
	*x = GetWeatherListByGeographicPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherListByGeographicPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherListByGeographicPointRequest) ProtoMessage() {}

func (x *GetWeatherListByGeographicPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherListByGeographicPointRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherListByGeographicPointRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{4}
}

type GetWeatherListByGeographicPointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeatherByGeographicPoint []*WeatherByGeographicPoint `protobuf:"bytes,1,rep,name=weatherByGeographicPoint,proto3" json:"weatherByGeographicPoint,omitempty"`
}

func (x *GetWeatherListByGeographicPointResponse) Reset() {
	*x = GetWeatherListByGeographicPointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherListByGeographicPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherListByGeographicPointResponse) ProtoMessage() {}

func (x *GetWeatherListByGeographicPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherListByGeographicPointResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherListByGeographicPointResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{5}
}

func (x *GetWeatherListByGeographicPointResponse) GetWeatherByGeographicPoint() []*WeatherByGeographicPoint {
	if x != nil {
		return x.WeatherByGeographicPoint
	}
	return nil
}

type WeatherByGeographicPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnixTime      int64   `protobuf:"varint,1,opt,name=unixTime,proto3" json:"unixTime,omitempty"`
	PointId       string  `protobuf:"bytes,2,opt,name=pointId,proto3" json:"pointId,omitempty"`
	PointName     string  `protobuf:"bytes,3,opt,name=pointName,proto3" json:"pointName,omitempty"`
	Lat           float64 `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon           float64 `protobuf:"fixed64,5,opt,name=lon,proto3" json:"lon,omitempty"`
	DisplayOrder  int32   `protobuf:"varint,6,opt,name=displayOrder,proto3" json:"displayOrder,omitempty"`
	WeatherIcon   string  `protobuf:"bytes,7,opt,name=weatherIcon,proto3" json:"weatherIcon,omitempty"`
	TempFeelsLike float64 `protobuf:"fixed64,8,opt,name=tempFeelsLike,proto3" json:"tempFeelsLike,omitempty"`
	TempMin       float64 `protobuf:"fixed64,9,opt,name=tempMin,proto3" json:"tempMin,omitempty"`
	TempMax       float64 `protobuf:"fixed64,10,opt,name=tempMax,proto3" json:"tempMax,omitempty"`
	Clouds        float64 `protobuf:"fixed64,11,opt,name=clouds,proto3" json:"clouds,omitempty"`
	RainFall      float64 `protobuf:"fixed64,12,opt,name=rainFall,proto3" json:"rainFall,omitempty"`
	Humidity      float64 `protobuf:"fixed64,13,opt,name=humidity,proto3" json:"humidity,omitempty"`
	WindSpeed     float64 `protobuf:"fixed64,14,opt,name=windSpeed,proto3" json:"windSpeed,omitempty"`
	Pressure      float64 `protobuf:"fixed64,15,opt,name=pressure,proto3" json:"pressure,omitempty"`
}

func (x *WeatherByGeographicPoint) Reset() {
	*x = WeatherByGeographicPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_geographicPoint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherByGeographicPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherByGeographicPoint) ProtoMessage() {}

func (x *WeatherByGeographicPoint) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_geographicPoint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherByGeographicPoint.ProtoReflect.Descriptor instead.
func (*WeatherByGeographicPoint) Descriptor() ([]byte, []int) {
	return file_src_proto_geographicPoint_proto_rawDescGZIP(), []int{6}
}

func (x *WeatherByGeographicPoint) GetUnixTime() int64 {
	if x != nil {
		return x.UnixTime
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetPointId() string {
	if x != nil {
		return x.PointId
	}
	return ""
}

func (x *WeatherByGeographicPoint) GetPointName() string {
	if x != nil {
		return x.PointName
	}
	return ""
}

func (x *WeatherByGeographicPoint) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetDisplayOrder() int32 {
	if x != nil {
		return x.DisplayOrder
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetWeatherIcon() string {
	if x != nil {
		return x.WeatherIcon
	}
	return ""
}

func (x *WeatherByGeographicPoint) GetTempFeelsLike() float64 {
	if x != nil {
		return x.TempFeelsLike
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetTempMin() float64 {
	if x != nil {
		return x.TempMin
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetTempMax() float64 {
	if x != nil {
		return x.TempMax
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetClouds() float64 {
	if x != nil {
		return x.Clouds
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetRainFall() float64 {
	if x != nil {
		return x.RainFall
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetHumidity() float64 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetWindSpeed() float64 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *WeatherByGeographicPoint) GetPressure() float64 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

var File_src_proto_geographicPoint_proto protoreflect.FileDescriptor

var file_src_proto_geographicPoint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9,
	0x01, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x6c,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xc5, 0x01, 0x0a, 0x1c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e,
	0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x01, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x37, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x65, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x1f, 0x52,
	0x65, 0x67, 0x73, 0x69, 0x74, 0x65, 0x72, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28,
	0x0a, 0x26, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x27, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x47, 0x65, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x18, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x42,
	0x79, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x42, 0x79, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x18, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x42, 0x79, 0x47,
	0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xbc,
	0x03, 0x0a, 0x18, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x42, 0x79, 0x47, 0x65, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75,
	0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x6d,
	0x70, 0x46, 0x65, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x46, 0x65, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x6d,
	0x70, 0x4d, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x65, 0x6d, 0x70,
	0x4d, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x61, 0x69, 0x6e, 0x46, 0x61, 0x6c, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72,
	0x61, 0x69, 0x6e, 0x46, 0x61, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x32, 0xb6, 0x03,
	0x0a, 0x16, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x47,
	0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x73, 0x69, 0x74, 0x65, 0x72, 0x47, 0x65, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x73, 0x69, 0x74, 0x65, 0x72, 0x47, 0x65, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x20, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x65, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x73, 0x69, 0x74, 0x65, 0x72, 0x47,
	0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x47, 0x65, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x47, 0x65, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_geographicPoint_proto_rawDescOnce sync.Once
	file_src_proto_geographicPoint_proto_rawDescData = file_src_proto_geographicPoint_proto_rawDesc
)

func file_src_proto_geographicPoint_proto_rawDescGZIP() []byte {
	file_src_proto_geographicPoint_proto_rawDescOnce.Do(func() {
		file_src_proto_geographicPoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_geographicPoint_proto_rawDescData)
	})
	return file_src_proto_geographicPoint_proto_rawDescData
}

var file_src_proto_geographicPoint_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_src_proto_geographicPoint_proto_goTypes = []interface{}{
	(*AddGeographicPointRequest)(nil),               // 0: pb.AddGeographicPointRequest
	(*UpdateGeographicPointRequest)(nil),            // 1: pb.UpdateGeographicPointRequest
	(*DeleteGeographicPointRequest)(nil),            // 2: pb.DeleteGeographicPointRequest
	(*RegsiterGeographicPointResponse)(nil),         // 3: pb.RegsiterGeographicPointResponse
	(*GetWeatherListByGeographicPointRequest)(nil),  // 4: pb.GetWeatherListByGeographicPointRequest
	(*GetWeatherListByGeographicPointResponse)(nil), // 5: pb.GetWeatherListByGeographicPointResponse
	(*WeatherByGeographicPoint)(nil),                // 6: pb.WeatherByGeographicPoint
}
var file_src_proto_geographicPoint_proto_depIdxs = []int32{
	6, // 0: pb.GetWeatherListByGeographicPointResponse.weatherByGeographicPoint:type_name -> pb.WeatherByGeographicPoint
	0, // 1: pb.GeographicPointService.AddGeographicPoint:input_type -> pb.AddGeographicPointRequest
	1, // 2: pb.GeographicPointService.UpdateGeographicPoint:input_type -> pb.UpdateGeographicPointRequest
	2, // 3: pb.GeographicPointService.DeleteGeographicPoint:input_type -> pb.DeleteGeographicPointRequest
	4, // 4: pb.GeographicPointService.GetWeatherListByGeographicPoint:input_type -> pb.GetWeatherListByGeographicPointRequest
	3, // 5: pb.GeographicPointService.AddGeographicPoint:output_type -> pb.RegsiterGeographicPointResponse
	3, // 6: pb.GeographicPointService.UpdateGeographicPoint:output_type -> pb.RegsiterGeographicPointResponse
	3, // 7: pb.GeographicPointService.DeleteGeographicPoint:output_type -> pb.RegsiterGeographicPointResponse
	5, // 8: pb.GeographicPointService.GetWeatherListByGeographicPoint:output_type -> pb.GetWeatherListByGeographicPointResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_proto_geographicPoint_proto_init() }
func file_src_proto_geographicPoint_proto_init() {
	if File_src_proto_geographicPoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_geographicPoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGeographicPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGeographicPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGeographicPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegsiterGeographicPointResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherListByGeographicPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherListByGeographicPointResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_proto_geographicPoint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherByGeographicPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_proto_geographicPoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_geographicPoint_proto_goTypes,
		DependencyIndexes: file_src_proto_geographicPoint_proto_depIdxs,
		MessageInfos:      file_src_proto_geographicPoint_proto_msgTypes,
	}.Build()
	File_src_proto_geographicPoint_proto = out.File
	file_src_proto_geographicPoint_proto_rawDesc = nil
	file_src_proto_geographicPoint_proto_goTypes = nil
	file_src_proto_geographicPoint_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeographicPointServiceClient is the client API for GeographicPointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeographicPointServiceClient interface {
	// 地点の追加
	AddGeographicPoint(ctx context.Context, in *AddGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error)
	// 地点の更新
	UpdateGeographicPoint(ctx context.Context, in *UpdateGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error)
	// 地点の削除
	DeleteGeographicPoint(ctx context.Context, in *DeleteGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error)
	// 地点毎の天気一覧
	GetWeatherListByGeographicPoint(ctx context.Context, in *GetWeatherListByGeographicPointRequest, opts ...grpc.CallOption) (*GetWeatherListByGeographicPointResponse, error)
}

type geographicPointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographicPointServiceClient(cc grpc.ClientConnInterface) GeographicPointServiceClient {
	return &geographicPointServiceClient{cc}
}

func (c *geographicPointServiceClient) AddGeographicPoint(ctx context.Context, in *AddGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error) {
	out := new(RegsiterGeographicPointResponse)
	err := c.cc.Invoke(ctx, "/pb.GeographicPointService/AddGeographicPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographicPointServiceClient) UpdateGeographicPoint(ctx context.Context, in *UpdateGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error) {
	out := new(RegsiterGeographicPointResponse)
	err := c.cc.Invoke(ctx, "/pb.GeographicPointService/UpdateGeographicPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographicPointServiceClient) DeleteGeographicPoint(ctx context.Context, in *DeleteGeographicPointRequest, opts ...grpc.CallOption) (*RegsiterGeographicPointResponse, error) {
	out := new(RegsiterGeographicPointResponse)
	err := c.cc.Invoke(ctx, "/pb.GeographicPointService/DeleteGeographicPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geographicPointServiceClient) GetWeatherListByGeographicPoint(ctx context.Context, in *GetWeatherListByGeographicPointRequest, opts ...grpc.CallOption) (*GetWeatherListByGeographicPointResponse, error) {
	out := new(GetWeatherListByGeographicPointResponse)
	err := c.cc.Invoke(ctx, "/pb.GeographicPointService/GetWeatherListByGeographicPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographicPointServiceServer is the server API for GeographicPointService service.
type GeographicPointServiceServer interface {
	// 地点の追加
	AddGeographicPoint(context.Context, *AddGeographicPointRequest) (*RegsiterGeographicPointResponse, error)
	// 地点の更新
	UpdateGeographicPoint(context.Context, *UpdateGeographicPointRequest) (*RegsiterGeographicPointResponse, error)
	// 地点の削除
	DeleteGeographicPoint(context.Context, *DeleteGeographicPointRequest) (*RegsiterGeographicPointResponse, error)
	// 地点毎の天気一覧
	GetWeatherListByGeographicPoint(context.Context, *GetWeatherListByGeographicPointRequest) (*GetWeatherListByGeographicPointResponse, error)
}

// UnimplementedGeographicPointServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeographicPointServiceServer struct {
}

func (*UnimplementedGeographicPointServiceServer) AddGeographicPoint(context.Context, *AddGeographicPointRequest) (*RegsiterGeographicPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGeographicPoint not implemented")
}
func (*UnimplementedGeographicPointServiceServer) UpdateGeographicPoint(context.Context, *UpdateGeographicPointRequest) (*RegsiterGeographicPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGeographicPoint not implemented")
}
func (*UnimplementedGeographicPointServiceServer) DeleteGeographicPoint(context.Context, *DeleteGeographicPointRequest) (*RegsiterGeographicPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGeographicPoint not implemented")
}
func (*UnimplementedGeographicPointServiceServer) GetWeatherListByGeographicPoint(context.Context, *GetWeatherListByGeographicPointRequest) (*GetWeatherListByGeographicPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherListByGeographicPoint not implemented")
}

func RegisterGeographicPointServiceServer(s *grpc.Server, srv GeographicPointServiceServer) {
	s.RegisterService(&_GeographicPointService_serviceDesc, srv)
}

func _GeographicPointService_AddGeographicPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGeographicPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicPointServiceServer).AddGeographicPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GeographicPointService/AddGeographicPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicPointServiceServer).AddGeographicPoint(ctx, req.(*AddGeographicPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographicPointService_UpdateGeographicPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGeographicPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicPointServiceServer).UpdateGeographicPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GeographicPointService/UpdateGeographicPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicPointServiceServer).UpdateGeographicPoint(ctx, req.(*UpdateGeographicPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographicPointService_DeleteGeographicPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGeographicPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicPointServiceServer).DeleteGeographicPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GeographicPointService/DeleteGeographicPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicPointServiceServer).DeleteGeographicPoint(ctx, req.(*DeleteGeographicPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeographicPointService_GetWeatherListByGeographicPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherListByGeographicPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicPointServiceServer).GetWeatherListByGeographicPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GeographicPointService/GetWeatherListByGeographicPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicPointServiceServer).GetWeatherListByGeographicPoint(ctx, req.(*GetWeatherListByGeographicPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeographicPointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GeographicPointService",
	HandlerType: (*GeographicPointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGeographicPoint",
			Handler:    _GeographicPointService_AddGeographicPoint_Handler,
		},
		{
			MethodName: "UpdateGeographicPoint",
			Handler:    _GeographicPointService_UpdateGeographicPoint_Handler,
		},
		{
			MethodName: "DeleteGeographicPoint",
			Handler:    _GeographicPointService_DeleteGeographicPoint_Handler,
		},
		{
			MethodName: "GetWeatherListByGeographicPoint",
			Handler:    _GeographicPointService_GetWeatherListByGeographicPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/geographicPoint.proto",
}
