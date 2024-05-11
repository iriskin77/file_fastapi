// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: file.proto

package api_gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	CreateFile(ctx context.Context, opts ...grpc.CallOption) (File_CreateFileClient, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) CreateFile(ctx context.Context, opts ...grpc.CallOption) (File_CreateFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &File_ServiceDesc.Streams[0], "/file.File/CreateFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileCreateFileClient{stream}
	return x, nil
}

type File_CreateFileClient interface {
	Send(*CreateFileRequest) error
	CloseAndRecv() (*CreateFileResponse, error)
	grpc.ClientStream
}

type fileCreateFileClient struct {
	grpc.ClientStream
}

func (x *fileCreateFileClient) Send(m *CreateFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileCreateFileClient) CloseAndRecv() (*CreateFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/file.File/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
	CreateFile(File_CreateFileServer) error
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) CreateFile(File_CreateFileServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFileServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_CreateFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServer).CreateFile(&fileCreateFileServer{stream})
}

type File_CreateFileServer interface {
	SendAndClose(*CreateFileResponse) error
	Recv() (*CreateFileRequest, error)
	grpc.ServerStream
}

type fileCreateFileServer struct {
	grpc.ServerStream
}

func (x *fileCreateFileServer) SendAndClose(m *CreateFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileCreateFileServer) Recv() (*CreateFileRequest, error) {
	m := new(CreateFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _File_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _File_GetFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateFile",
			Handler:       _File_CreateFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file.proto",
}