// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Grpc.Logic.proto

package cim

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("Grpc.Logic.proto", fileDescriptor_3e29f99bd654035c) }

var fileDescriptor_3e29f99bd654035c = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x37, 0xc1, 0x3f, 0x58, 0x13, 0xd0, 0xc6, 0xc4, 0xa4, 0x07, 0x0f, 0x1b, 0xcf, 0x3d,
	0xe8, 0xcd, 0xdb, 0x52, 0x10, 0x36, 0xa1, 0x89, 0x59, 0x3c, 0x79, 0xab, 0x65, 0xac, 0xcd, 0x42,
	0xbb, 0xd2, 0x7a, 0xe0, 0x73, 0xfb, 0x05, 0xcc, 0x96, 0x2e, 0x60, 0x60, 0x13, 0x6f, 0x3b, 0xef,
	0xf7, 0xf6, 0xcd, 0x64, 0xa6, 0xe8, 0x6a, 0xbc, 0xaa, 0x24, 0x9d, 0x5a, 0xa5, 0x25, 0xad, 0x56,
	0xd6, 0x5b, 0xdc, 0x65, 0x39, 0xa7, 0xb5, 0x4a, 0x7a, 0xf5, 0xd7, 0x54, 0x3b, 0xbf, 0x21, 0xa4,
	0x1f, 0x6a, 0xab, 0xb4, 0x89, 0xc2, 0x75, 0x2d, 0x70, 0x70, 0x4e, 0x28, 0x88, 0x52, 0x2f, 0xe4,
	0x0d, 0xe1, 0x63, 0x53, 0x3f, 0xfc, 0x74, 0xd0, 0x69, 0x48, 0xc7, 0x14, 0x75, 0x67, 0x62, 0x3d,
	0x81, 0xc5, 0xc2, 0xe2, 0x3e, 0x6d, 0x9a, 0xd0, 0x20, 0x90, 0x3d, 0x61, 0xb4, 0xac, 0xfc, 0x3a,
	0x4d, 0xf0, 0x13, 0x3a, 0x79, 0xd1, 0x46, 0xe1, 0x5b, 0xba, 0x6b, 0xcb, 0x72, 0x3e, 0x01, 0xb1,
	0xf2, 0x03, 0x10, 0x9e, 0xb4, 0x81, 0x34, 0xc1, 0x23, 0x74, 0x91, 0x7d, 0xfb, 0xcf, 0x57, 0x5b,
	0x82, 0xc1, 0xe4, 0xaf, 0x6f, 0x0b, 0x0a, 0xf8, 0x22, 0xad, 0xcc, 0x55, 0x69, 0x82, 0xe7, 0xe8,
	0xa6, 0x00, 0x09, 0xc6, 0x33, 0x6b, 0xbc, 0x90, 0x7e, 0x06, 0xce, 0x69, 0x6b, 0xf0, 0x3d, 0xdd,
	0x6e, 0x86, 0xe5, 0xfc, 0x98, 0xa5, 0xce, 0xfe, 0x87, 0x2b, 0x74, 0x79, 0x46, 0x68, 0x0c, 0x9e,
	0x3b, 0x55, 0x5b, 0x9b, 0x69, 0xe3, 0x5f, 0x3b, 0xb0, 0x37, 0xed, 0x21, 0x0b, 0x39, 0x43, 0x74,
	0x39, 0x03, 0x33, 0xe7, 0x4e, 0x0d, 0x85, 0x17, 0x71, 0x6f, 0xcd, 0x75, 0x58, 0xce, 0x23, 0x88,
	0x29, 0x87, 0x20, 0x93, 0x65, 0x9a, 0x60, 0x86, 0x7a, 0x05, 0x08, 0x99, 0xc9, 0xb2, 0x09, 0xba,
	0x6b, 0xf1, 0x17, 0x20, 0xe6, 0x99, 0x2c, 0x8f, 0xdc, 0x6e, 0x70, 0x3e, 0xe9, 0xbc, 0x75, 0xa4,
	0x5e, 0xbe, 0x9f, 0x85, 0x57, 0xf0, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x82, 0xb7, 0x9c,
	0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicClient interface {
	// 传递我方信息，双向GRPC
	SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Empty, error)
	// ping
	Ping(ctx context.Context, in *CIMHeartBeat, opts ...grpc.CallOption) (*CIMHeartBeat, error)
	// 验证token
	AuthToken(ctx context.Context, in *CIMAuthTokenReq, opts ...grpc.CallOption) (*CIMAuthTokenRsp, error)
	// 查询会话列表
	RecentContactSession(ctx context.Context, in *CIMRecentContactSessionReq, opts ...grpc.CallOption) (*CIMRecentContactSessionRsp, error)
	// 查询历史消息列表
	GetMsgList(ctx context.Context, in *CIMGetMsgListReq, opts ...grpc.CallOption) (*CIMGetMsgListRsp, error)
	// 发消息
	SendMsgData(ctx context.Context, in *CIMMsgData, opts ...grpc.CallOption) (*CIMMsgDataAck, error)
	// 已读消息回执
	ReacAckMsgData(ctx context.Context, in *CIMMsgDataReadAck, opts ...grpc.CallOption) (*Empty, error)
}

type logicClient struct {
	cc *grpc.ClientConn
}

func NewLogicClient(cc *grpc.ClientConn) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) Ping(ctx context.Context, in *CIMHeartBeat, opts ...grpc.CallOption) (*CIMHeartBeat, error) {
	out := new(CIMHeartBeat)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) AuthToken(ctx context.Context, in *CIMAuthTokenReq, opts ...grpc.CallOption) (*CIMAuthTokenRsp, error) {
	out := new(CIMAuthTokenRsp)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/AuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) RecentContactSession(ctx context.Context, in *CIMRecentContactSessionReq, opts ...grpc.CallOption) (*CIMRecentContactSessionRsp, error) {
	out := new(CIMRecentContactSessionRsp)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/RecentContactSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) GetMsgList(ctx context.Context, in *CIMGetMsgListReq, opts ...grpc.CallOption) (*CIMGetMsgListRsp, error) {
	out := new(CIMGetMsgListRsp)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/GetMsgList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) SendMsgData(ctx context.Context, in *CIMMsgData, opts ...grpc.CallOption) (*CIMMsgDataAck, error) {
	out := new(CIMMsgDataAck)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/SendMsgData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) ReacAckMsgData(ctx context.Context, in *CIMMsgDataReadAck, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/CIM.Grpc.Logic/ReacAckMsgData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
type LogicServer interface {
	// 传递我方信息，双向GRPC
	SayHello(context.Context, *Hello) (*Empty, error)
	// ping
	Ping(context.Context, *CIMHeartBeat) (*CIMHeartBeat, error)
	// 验证token
	AuthToken(context.Context, *CIMAuthTokenReq) (*CIMAuthTokenRsp, error)
	// 查询会话列表
	RecentContactSession(context.Context, *CIMRecentContactSessionReq) (*CIMRecentContactSessionRsp, error)
	// 查询历史消息列表
	GetMsgList(context.Context, *CIMGetMsgListReq) (*CIMGetMsgListRsp, error)
	// 发消息
	SendMsgData(context.Context, *CIMMsgData) (*CIMMsgDataAck, error)
	// 已读消息回执
	ReacAckMsgData(context.Context, *CIMMsgDataReadAck) (*Empty, error)
}

// UnimplementedLogicServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (*UnimplementedLogicServer) SayHello(ctx context.Context, req *Hello) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedLogicServer) Ping(ctx context.Context, req *CIMHeartBeat) (*CIMHeartBeat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedLogicServer) AuthToken(ctx context.Context, req *CIMAuthTokenReq) (*CIMAuthTokenRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthToken not implemented")
}
func (*UnimplementedLogicServer) RecentContactSession(ctx context.Context, req *CIMRecentContactSessionReq) (*CIMRecentContactSessionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecentContactSession not implemented")
}
func (*UnimplementedLogicServer) GetMsgList(ctx context.Context, req *CIMGetMsgListReq) (*CIMGetMsgListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsgList not implemented")
}
func (*UnimplementedLogicServer) SendMsgData(ctx context.Context, req *CIMMsgData) (*CIMMsgDataAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgData not implemented")
}
func (*UnimplementedLogicServer) ReacAckMsgData(ctx context.Context, req *CIMMsgDataReadAck) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReacAckMsgData not implemented")
}

func RegisterLogicServer(s *grpc.Server, srv LogicServer) {
	s.RegisterService(&_Logic_serviceDesc, srv)
}

func _Logic_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SayHello(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMHeartBeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Ping(ctx, req.(*CIMHeartBeat))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_AuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMAuthTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).AuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/AuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).AuthToken(ctx, req.(*CIMAuthTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_RecentContactSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMRecentContactSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).RecentContactSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/RecentContactSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).RecentContactSession(ctx, req.(*CIMRecentContactSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_GetMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMGetMsgListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).GetMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/GetMsgList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).GetMsgList(ctx, req.(*CIMGetMsgListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_SendMsgData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMMsgData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SendMsgData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/SendMsgData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SendMsgData(ctx, req.(*CIMMsgData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_ReacAckMsgData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMMsgDataReadAck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).ReacAckMsgData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CIM.Grpc.Logic/ReacAckMsgData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).ReacAckMsgData(ctx, req.(*CIMMsgDataReadAck))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CIM.Grpc.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Logic_SayHello_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Logic_Ping_Handler,
		},
		{
			MethodName: "AuthToken",
			Handler:    _Logic_AuthToken_Handler,
		},
		{
			MethodName: "RecentContactSession",
			Handler:    _Logic_RecentContactSession_Handler,
		},
		{
			MethodName: "GetMsgList",
			Handler:    _Logic_GetMsgList_Handler,
		},
		{
			MethodName: "SendMsgData",
			Handler:    _Logic_SendMsgData_Handler,
		},
		{
			MethodName: "ReacAckMsgData",
			Handler:    _Logic_ReacAckMsgData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Grpc.Logic.proto",
}
