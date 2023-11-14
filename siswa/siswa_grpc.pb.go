// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: siswa/siswa.proto

package siswa

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

// SiswaServiceClient is the client API for SiswaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiswaServiceClient interface {
	GetSiswa(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*Siswa, error)
}

type siswaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSiswaServiceClient(cc grpc.ClientConnInterface) SiswaServiceClient {
	return &siswaServiceClient{cc}
}

func (c *siswaServiceClient) GetSiswa(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*Siswa, error) {
	out := new(Siswa)
	err := c.cc.Invoke(ctx, "/siswa.SiswaService/GetSiswa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiswaServiceServer is the server API for SiswaService service.
// All implementations must embed UnimplementedSiswaServiceServer
// for forward compatibility
type SiswaServiceServer interface {
	GetSiswa(context.Context, *SiswaRequest) (*Siswa, error)
	mustEmbedUnimplementedSiswaServiceServer()
}

// UnimplementedSiswaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSiswaServiceServer struct {
}

func (UnimplementedSiswaServiceServer) GetSiswa(context.Context, *SiswaRequest) (*Siswa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSiswa not implemented")
}
func (UnimplementedSiswaServiceServer) mustEmbedUnimplementedSiswaServiceServer() {}

// UnsafeSiswaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiswaServiceServer will
// result in compilation errors.
type UnsafeSiswaServiceServer interface {
	mustEmbedUnimplementedSiswaServiceServer()
}

func RegisterSiswaServiceServer(s grpc.ServiceRegistrar, srv SiswaServiceServer) {
	s.RegisterService(&SiswaService_ServiceDesc, srv)
}

func _SiswaService_GetSiswa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiswaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiswaServiceServer).GetSiswa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/siswa.SiswaService/GetSiswa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiswaServiceServer).GetSiswa(ctx, req.(*SiswaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SiswaService_ServiceDesc is the grpc.ServiceDesc for SiswaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiswaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "siswa.SiswaService",
	HandlerType: (*SiswaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSiswa",
			Handler:    _SiswaService_GetSiswa_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswa/siswa.proto",
}