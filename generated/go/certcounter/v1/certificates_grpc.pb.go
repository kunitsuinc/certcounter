// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: certcounter/v1/certificates.proto

package v1

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

// CertificatesServiceClient is the client API for CertificatesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificatesServiceClient interface {
	Issue(ctx context.Context, in *CertificatesServiceIssueRequest, opts ...grpc.CallOption) (*CertificatesServiceIssueResponse, error)
}

type certificatesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificatesServiceClient(cc grpc.ClientConnInterface) CertificatesServiceClient {
	return &certificatesServiceClient{cc}
}

func (c *certificatesServiceClient) Issue(ctx context.Context, in *CertificatesServiceIssueRequest, opts ...grpc.CallOption) (*CertificatesServiceIssueResponse, error) {
	out := new(CertificatesServiceIssueResponse)
	err := c.cc.Invoke(ctx, "/certcounter.v1.CertificatesService/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificatesServiceServer is the server API for CertificatesService service.
// All implementations must embed UnimplementedCertificatesServiceServer
// for forward compatibility
type CertificatesServiceServer interface {
	Issue(context.Context, *CertificatesServiceIssueRequest) (*CertificatesServiceIssueResponse, error)
	mustEmbedUnimplementedCertificatesServiceServer()
}

// UnimplementedCertificatesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificatesServiceServer struct {
}

func (UnimplementedCertificatesServiceServer) Issue(context.Context, *CertificatesServiceIssueRequest) (*CertificatesServiceIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedCertificatesServiceServer) mustEmbedUnimplementedCertificatesServiceServer() {}

// UnsafeCertificatesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificatesServiceServer will
// result in compilation errors.
type UnsafeCertificatesServiceServer interface {
	mustEmbedUnimplementedCertificatesServiceServer()
}

func RegisterCertificatesServiceServer(s grpc.ServiceRegistrar, srv CertificatesServiceServer) {
	s.RegisterService(&CertificatesService_ServiceDesc, srv)
}

func _CertificatesService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificatesServiceIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certcounter.v1.CertificatesService/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).Issue(ctx, req.(*CertificatesServiceIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificatesService_ServiceDesc is the grpc.ServiceDesc for CertificatesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificatesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "certcounter.v1.CertificatesService",
	HandlerType: (*CertificatesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _CertificatesService_Issue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certcounter/v1/certificates.proto",
}