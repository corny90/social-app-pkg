// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpc-proto/media.proto

package grpc_proto

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

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	PingApiMedia(ctx context.Context, in *PingApiMediaRequest, opts ...grpc.CallOption) (*PingApiMediaResponse, error)
	PostUploadMedia(ctx context.Context, in *PostUploadMediaRequest, opts ...grpc.CallOption) (*PostUploadMediaResponse, error)
	PostFetchMedia(ctx context.Context, in *PostFetchMediaRequest, opts ...grpc.CallOption) (*PostFetchMediaResponse, error)
	PostDeleteMedia(ctx context.Context, in *PostDeleteMediaRequest, opts ...grpc.CallOption) (*PostDeleteMediaResponse, error)
	PostsFetchMedia(ctx context.Context, in *PostsFetchMediaRequest, opts ...grpc.CallOption) (*PostsFetchMediaResponse, error)
	AvatarMediaUpload(ctx context.Context, in *AvatarMediaUploadRequest, opts ...grpc.CallOption) (*AvatarMediaUploadResponse, error)
	AvatarMediaFetch(ctx context.Context, in *AvatarMediaFetchRequest, opts ...grpc.CallOption) (*AvatarMediaFetchResponse, error)
	AvatarMediaFetchMultiple(ctx context.Context, in *AvatarMediaFetchMultipleRequest, opts ...grpc.CallOption) (*AvatarMediaFetchMultipleResponse, error)
	AvatarMediaDelete(ctx context.Context, in *AvatarMediaDeleteRequest, opts ...grpc.CallOption) (*AvatarMediaDeleteResponse, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) PingApiMedia(ctx context.Context, in *PingApiMediaRequest, opts ...grpc.CallOption) (*PingApiMediaResponse, error) {
	out := new(PingApiMediaResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/PingApiMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostUploadMedia(ctx context.Context, in *PostUploadMediaRequest, opts ...grpc.CallOption) (*PostUploadMediaResponse, error) {
	out := new(PostUploadMediaResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/PostUploadMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostFetchMedia(ctx context.Context, in *PostFetchMediaRequest, opts ...grpc.CallOption) (*PostFetchMediaResponse, error) {
	out := new(PostFetchMediaResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/PostFetchMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostDeleteMedia(ctx context.Context, in *PostDeleteMediaRequest, opts ...grpc.CallOption) (*PostDeleteMediaResponse, error) {
	out := new(PostDeleteMediaResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/PostDeleteMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostsFetchMedia(ctx context.Context, in *PostsFetchMediaRequest, opts ...grpc.CallOption) (*PostsFetchMediaResponse, error) {
	out := new(PostsFetchMediaResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/PostsFetchMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaUpload(ctx context.Context, in *AvatarMediaUploadRequest, opts ...grpc.CallOption) (*AvatarMediaUploadResponse, error) {
	out := new(AvatarMediaUploadResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/AvatarMediaUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaFetch(ctx context.Context, in *AvatarMediaFetchRequest, opts ...grpc.CallOption) (*AvatarMediaFetchResponse, error) {
	out := new(AvatarMediaFetchResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/AvatarMediaFetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaFetchMultiple(ctx context.Context, in *AvatarMediaFetchMultipleRequest, opts ...grpc.CallOption) (*AvatarMediaFetchMultipleResponse, error) {
	out := new(AvatarMediaFetchMultipleResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/AvatarMediaFetchMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaDelete(ctx context.Context, in *AvatarMediaDeleteRequest, opts ...grpc.CallOption) (*AvatarMediaDeleteResponse, error) {
	out := new(AvatarMediaDeleteResponse)
	err := c.cc.Invoke(ctx, "/media.MediaService/AvatarMediaDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	PingApiMedia(context.Context, *PingApiMediaRequest) (*PingApiMediaResponse, error)
	PostUploadMedia(context.Context, *PostUploadMediaRequest) (*PostUploadMediaResponse, error)
	PostFetchMedia(context.Context, *PostFetchMediaRequest) (*PostFetchMediaResponse, error)
	PostDeleteMedia(context.Context, *PostDeleteMediaRequest) (*PostDeleteMediaResponse, error)
	PostsFetchMedia(context.Context, *PostsFetchMediaRequest) (*PostsFetchMediaResponse, error)
	AvatarMediaUpload(context.Context, *AvatarMediaUploadRequest) (*AvatarMediaUploadResponse, error)
	AvatarMediaFetch(context.Context, *AvatarMediaFetchRequest) (*AvatarMediaFetchResponse, error)
	AvatarMediaFetchMultiple(context.Context, *AvatarMediaFetchMultipleRequest) (*AvatarMediaFetchMultipleResponse, error)
	AvatarMediaDelete(context.Context, *AvatarMediaDeleteRequest) (*AvatarMediaDeleteResponse, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) PingApiMedia(context.Context, *PingApiMediaRequest) (*PingApiMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingApiMedia not implemented")
}
func (UnimplementedMediaServiceServer) PostUploadMedia(context.Context, *PostUploadMediaRequest) (*PostUploadMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUploadMedia not implemented")
}
func (UnimplementedMediaServiceServer) PostFetchMedia(context.Context, *PostFetchMediaRequest) (*PostFetchMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostFetchMedia not implemented")
}
func (UnimplementedMediaServiceServer) PostDeleteMedia(context.Context, *PostDeleteMediaRequest) (*PostDeleteMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDeleteMedia not implemented")
}
func (UnimplementedMediaServiceServer) PostsFetchMedia(context.Context, *PostsFetchMediaRequest) (*PostsFetchMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostsFetchMedia not implemented")
}
func (UnimplementedMediaServiceServer) AvatarMediaUpload(context.Context, *AvatarMediaUploadRequest) (*AvatarMediaUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvatarMediaUpload not implemented")
}
func (UnimplementedMediaServiceServer) AvatarMediaFetch(context.Context, *AvatarMediaFetchRequest) (*AvatarMediaFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvatarMediaFetch not implemented")
}
func (UnimplementedMediaServiceServer) AvatarMediaFetchMultiple(context.Context, *AvatarMediaFetchMultipleRequest) (*AvatarMediaFetchMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvatarMediaFetchMultiple not implemented")
}
func (UnimplementedMediaServiceServer) AvatarMediaDelete(context.Context, *AvatarMediaDeleteRequest) (*AvatarMediaDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvatarMediaDelete not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_PingApiMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingApiMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).PingApiMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/PingApiMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).PingApiMedia(ctx, req.(*PingApiMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_PostUploadMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUploadMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).PostUploadMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/PostUploadMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).PostUploadMedia(ctx, req.(*PostUploadMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_PostFetchMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostFetchMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).PostFetchMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/PostFetchMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).PostFetchMedia(ctx, req.(*PostFetchMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_PostDeleteMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDeleteMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).PostDeleteMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/PostDeleteMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).PostDeleteMedia(ctx, req.(*PostDeleteMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_PostsFetchMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsFetchMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).PostsFetchMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/PostsFetchMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).PostsFetchMedia(ctx, req.(*PostsFetchMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_AvatarMediaUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarMediaUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AvatarMediaUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/AvatarMediaUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AvatarMediaUpload(ctx, req.(*AvatarMediaUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_AvatarMediaFetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarMediaFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AvatarMediaFetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/AvatarMediaFetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AvatarMediaFetch(ctx, req.(*AvatarMediaFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_AvatarMediaFetchMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarMediaFetchMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AvatarMediaFetchMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/AvatarMediaFetchMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AvatarMediaFetchMultiple(ctx, req.(*AvatarMediaFetchMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_AvatarMediaDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarMediaDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AvatarMediaDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/AvatarMediaDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AvatarMediaDelete(ctx, req.(*AvatarMediaDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "media.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingApiMedia",
			Handler:    _MediaService_PingApiMedia_Handler,
		},
		{
			MethodName: "PostUploadMedia",
			Handler:    _MediaService_PostUploadMedia_Handler,
		},
		{
			MethodName: "PostFetchMedia",
			Handler:    _MediaService_PostFetchMedia_Handler,
		},
		{
			MethodName: "PostDeleteMedia",
			Handler:    _MediaService_PostDeleteMedia_Handler,
		},
		{
			MethodName: "PostsFetchMedia",
			Handler:    _MediaService_PostsFetchMedia_Handler,
		},
		{
			MethodName: "AvatarMediaUpload",
			Handler:    _MediaService_AvatarMediaUpload_Handler,
		},
		{
			MethodName: "AvatarMediaFetch",
			Handler:    _MediaService_AvatarMediaFetch_Handler,
		},
		{
			MethodName: "AvatarMediaFetchMultiple",
			Handler:    _MediaService_AvatarMediaFetchMultiple_Handler,
		},
		{
			MethodName: "AvatarMediaDelete",
			Handler:    _MediaService_AvatarMediaDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-proto/media.proto",
}
