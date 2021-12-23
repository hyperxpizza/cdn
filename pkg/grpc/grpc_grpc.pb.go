// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// CDNGrpcServiceClient is the client API for CDNGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CDNGrpcServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (CDNGrpcService_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (CDNGrpcService_DownloadFileClient, error)
	SearchFiles(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type cDNGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCDNGrpcServiceClient(cc grpc.ClientConnInterface) CDNGrpcServiceClient {
	return &cDNGrpcServiceClient{cc}
}

func (c *cDNGrpcServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (CDNGrpcService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &CDNGrpcService_ServiceDesc.Streams[0], "/CDNGrpcService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDNGrpcServiceUploadFileClient{stream}
	return x, nil
}

type CDNGrpcService_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type cDNGrpcServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *cDNGrpcServiceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cDNGrpcServiceUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDNGrpcServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (CDNGrpcService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &CDNGrpcService_ServiceDesc.Streams[1], "/CDNGrpcService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDNGrpcServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDNGrpcService_DownloadFileClient interface {
	Recv() (*DownlaodFileRespose, error)
	grpc.ClientStream
}

type cDNGrpcServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *cDNGrpcServiceDownloadFileClient) Recv() (*DownlaodFileRespose, error) {
	m := new(DownlaodFileRespose)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDNGrpcServiceClient) SearchFiles(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/CDNGrpcService/SearchFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CDNGrpcServiceServer is the server API for CDNGrpcService service.
// All implementations must embed UnimplementedCDNGrpcServiceServer
// for forward compatibility
type CDNGrpcServiceServer interface {
	UploadFile(CDNGrpcService_UploadFileServer) error
	DownloadFile(*DownloadFileRequest, CDNGrpcService_DownloadFileServer) error
	SearchFiles(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedCDNGrpcServiceServer()
}

// UnimplementedCDNGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCDNGrpcServiceServer struct {
}

func (UnimplementedCDNGrpcServiceServer) UploadFile(CDNGrpcService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedCDNGrpcServiceServer) DownloadFile(*DownloadFileRequest, CDNGrpcService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedCDNGrpcServiceServer) SearchFiles(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFiles not implemented")
}
func (UnimplementedCDNGrpcServiceServer) mustEmbedUnimplementedCDNGrpcServiceServer() {}

// UnsafeCDNGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CDNGrpcServiceServer will
// result in compilation errors.
type UnsafeCDNGrpcServiceServer interface {
	mustEmbedUnimplementedCDNGrpcServiceServer()
}

func RegisterCDNGrpcServiceServer(s grpc.ServiceRegistrar, srv CDNGrpcServiceServer) {
	s.RegisterService(&CDNGrpcService_ServiceDesc, srv)
}

func _CDNGrpcService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CDNGrpcServiceServer).UploadFile(&cDNGrpcServiceUploadFileServer{stream})
}

type CDNGrpcService_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type cDNGrpcServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *cDNGrpcServiceUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cDNGrpcServiceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CDNGrpcService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDNGrpcServiceServer).DownloadFile(m, &cDNGrpcServiceDownloadFileServer{stream})
}

type CDNGrpcService_DownloadFileServer interface {
	Send(*DownlaodFileRespose) error
	grpc.ServerStream
}

type cDNGrpcServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *cDNGrpcServiceDownloadFileServer) Send(m *DownlaodFileRespose) error {
	return x.ServerStream.SendMsg(m)
}

func _CDNGrpcService_SearchFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDNGrpcServiceServer).SearchFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CDNGrpcService/SearchFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDNGrpcServiceServer).SearchFiles(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CDNGrpcService_ServiceDesc is the grpc.ServiceDesc for CDNGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CDNGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CDNGrpcService",
	HandlerType: (*CDNGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchFiles",
			Handler:    _CDNGrpcService_SearchFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _CDNGrpcService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _CDNGrpcService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}