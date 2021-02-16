// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/logs/v1/logs_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "ingest/internal/opentelemetry-proto-gen/logs/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExportLogsServiceRequest struct {
	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceLogs         []*v1.ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportLogsServiceRequest) Reset()         { *m = ExportLogsServiceRequest{} }
func (m *ExportLogsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceRequest) ProtoMessage()    {}
func (*ExportLogsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{0}
}
func (m *ExportLogsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceRequest.Unmarshal(m, b)
}
func (m *ExportLogsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceRequest.Merge(m, src)
}
func (m *ExportLogsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceRequest.Size(m)
}
func (m *ExportLogsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceRequest proto.InternalMessageInfo

func (m *ExportLogsServiceRequest) GetResourceLogs() []*v1.ResourceLogs {
	if m != nil {
		return m.ResourceLogs
	}
	return nil
}

type ExportLogsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportLogsServiceResponse) Reset()         { *m = ExportLogsServiceResponse{} }
func (m *ExportLogsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceResponse) ProtoMessage()    {}
func (*ExportLogsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{1}
}
func (m *ExportLogsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceResponse.Unmarshal(m, b)
}
func (m *ExportLogsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceResponse.Merge(m, src)
}
func (m *ExportLogsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceResponse.Size(m)
}
func (m *ExportLogsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportLogsServiceRequest)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest")
	proto.RegisterType((*ExportLogsServiceResponse)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/logs/v1/logs_service.proto", fileDescriptor_8e3bf87aaa43acd4)
}

var fileDescriptor_8e3bf87aaa43acd4 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc8, 0x2f, 0x48, 0xcd,
	0x2b, 0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x4f, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0xcf, 0xc9, 0x4f, 0x2f, 0xd6, 0x2f,
	0x33, 0x04, 0xd3, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x60, 0x45, 0x42, 0xaa,
	0x28, 0x3a, 0x21, 0x82, 0x7a, 0x70, 0x9d, 0x7a, 0x20, 0x1d, 0x7a, 0x65, 0x86, 0x52, 0x6a, 0xd8,
	0x2c, 0x40, 0x36, 0x16, 0xa2, 0x53, 0x29, 0x8b, 0x4b, 0xc2, 0xb5, 0xa2, 0x20, 0xbf, 0xa8, 0xc4,
	0x27, 0x3f, 0xbd, 0x38, 0x18, 0x62, 0x53, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x1f,
	0x17, 0x6f, 0x51, 0x6a, 0x71, 0x7e, 0x69, 0x51, 0x72, 0x6a, 0x3c, 0x48, 0x8b, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x91, 0xa6, 0x1e, 0x36, 0x27, 0x40, 0x2d, 0xd6, 0x0b, 0x82, 0xea, 0x00, 0x99,
	0x17, 0xc4, 0x53, 0x84, 0xc4, 0x53, 0x92, 0xe6, 0x92, 0xc4, 0x62, 0x57, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0xd1, 0x5c, 0x46, 0x2e, 0x6e, 0x24, 0x71, 0xa1, 0x5e, 0x46, 0x2e, 0x36, 0x88, 0x6a,
	0x21, 0x7b, 0x3d, 0xa2, 0xfc, 0xac, 0x87, 0xcb, 0x23, 0x52, 0x0e, 0xe4, 0x1b, 0x00, 0x71, 0x9d,
	0x12, 0x83, 0x53, 0x25, 0x97, 0x46, 0x66, 0x3e, 0x71, 0xe6, 0x38, 0x09, 0x20, 0x19, 0x11, 0x00,
	0x52, 0x13, 0xc0, 0x18, 0x65, 0x99, 0x99, 0x97, 0x9e, 0x5a, 0x5c, 0xa2, 0x9f, 0x99, 0x57, 0x92,
	0x5a, 0x94, 0x97, 0x98, 0xa3, 0x8f, 0x62, 0x94, 0x2e, 0xd8, 0x28, 0xdd, 0xf4, 0xd4, 0x3c, 0xcc,
	0x54, 0x90, 0xc4, 0x06, 0x96, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x5a, 0x14, 0xc2,
	0x35, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error)
}

type logsServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogsServiceClient(cc *grpc.ClientConn) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error) {
	out := new(ExportLogsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.logs.v1.LogsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServiceServer is the server API for LogsService service.
type LogsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error)
}

// UnimplementedLogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (*UnimplementedLogsServiceServer) Export(ctx context.Context, req *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterLogsServiceServer(s *grpc.Server, srv LogsServiceServer) {
	s.RegisterService(&_LogsService_serviceDesc, srv)
}

func _LogsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportLogsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.logs.v1.LogsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Export(ctx, req.(*ExportLogsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.logs.v1.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _LogsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/logs/v1/logs_service.proto",
}