// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MediaService_PingApiMedia_FullMethodName             = "/media.MediaService/PingApiMedia"
	MediaService_PostUploadMedia_FullMethodName          = "/media.MediaService/PostUploadMedia"
	MediaService_PostFetchMedia_FullMethodName           = "/media.MediaService/PostFetchMedia"
	MediaService_PostDeleteMedia_FullMethodName          = "/media.MediaService/PostDeleteMedia"
	MediaService_PostsFetchMedia_FullMethodName          = "/media.MediaService/PostsFetchMedia"
	MediaService_AvatarMediaUpload_FullMethodName        = "/media.MediaService/AvatarMediaUpload"
	MediaService_AvatarMediaFetch_FullMethodName         = "/media.MediaService/AvatarMediaFetch"
	MediaService_AvatarMediaFetchMultiple_FullMethodName = "/media.MediaService/AvatarMediaFetchMultiple"
	MediaService_AvatarMediaDelete_FullMethodName        = "/media.MediaService/AvatarMediaDelete"
	MediaService_BackgroundMediaUpload_FullMethodName    = "/media.MediaService/BackgroundMediaUpload"
	MediaService_BackgroundMediaFetch_FullMethodName     = "/media.MediaService/BackgroundMediaFetch"
	MediaService_BackgroundMediaDelete_FullMethodName    = "/media.MediaService/BackgroundMediaDelete"
)

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
	BackgroundMediaUpload(ctx context.Context, in *BackgroundMediaUploadRequest, opts ...grpc.CallOption) (*BackgroundMediaUploadResponse, error)
	BackgroundMediaFetch(ctx context.Context, in *BackgroundMediaFetchRequest, opts ...grpc.CallOption) (*BackgroundMediaFetchResponse, error)
	BackgroundMediaDelete(ctx context.Context, in *BackgroundMediaDeleteRequest, opts ...grpc.CallOption) (*BackgroundMediaDeleteResponse, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) PingApiMedia(ctx context.Context, in *PingApiMediaRequest, opts ...grpc.CallOption) (*PingApiMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingApiMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_PingApiMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostUploadMedia(ctx context.Context, in *PostUploadMediaRequest, opts ...grpc.CallOption) (*PostUploadMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostUploadMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_PostUploadMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostFetchMedia(ctx context.Context, in *PostFetchMediaRequest, opts ...grpc.CallOption) (*PostFetchMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostFetchMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_PostFetchMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostDeleteMedia(ctx context.Context, in *PostDeleteMediaRequest, opts ...grpc.CallOption) (*PostDeleteMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostDeleteMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_PostDeleteMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) PostsFetchMedia(ctx context.Context, in *PostsFetchMediaRequest, opts ...grpc.CallOption) (*PostsFetchMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostsFetchMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_PostsFetchMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaUpload(ctx context.Context, in *AvatarMediaUploadRequest, opts ...grpc.CallOption) (*AvatarMediaUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvatarMediaUploadResponse)
	err := c.cc.Invoke(ctx, MediaService_AvatarMediaUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaFetch(ctx context.Context, in *AvatarMediaFetchRequest, opts ...grpc.CallOption) (*AvatarMediaFetchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvatarMediaFetchResponse)
	err := c.cc.Invoke(ctx, MediaService_AvatarMediaFetch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaFetchMultiple(ctx context.Context, in *AvatarMediaFetchMultipleRequest, opts ...grpc.CallOption) (*AvatarMediaFetchMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvatarMediaFetchMultipleResponse)
	err := c.cc.Invoke(ctx, MediaService_AvatarMediaFetchMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AvatarMediaDelete(ctx context.Context, in *AvatarMediaDeleteRequest, opts ...grpc.CallOption) (*AvatarMediaDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvatarMediaDeleteResponse)
	err := c.cc.Invoke(ctx, MediaService_AvatarMediaDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) BackgroundMediaUpload(ctx context.Context, in *BackgroundMediaUploadRequest, opts ...grpc.CallOption) (*BackgroundMediaUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackgroundMediaUploadResponse)
	err := c.cc.Invoke(ctx, MediaService_BackgroundMediaUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) BackgroundMediaFetch(ctx context.Context, in *BackgroundMediaFetchRequest, opts ...grpc.CallOption) (*BackgroundMediaFetchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackgroundMediaFetchResponse)
	err := c.cc.Invoke(ctx, MediaService_BackgroundMediaFetch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) BackgroundMediaDelete(ctx context.Context, in *BackgroundMediaDeleteRequest, opts ...grpc.CallOption) (*BackgroundMediaDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackgroundMediaDeleteResponse)
	err := c.cc.Invoke(ctx, MediaService_BackgroundMediaDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility.
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
	BackgroundMediaUpload(context.Context, *BackgroundMediaUploadRequest) (*BackgroundMediaUploadResponse, error)
	BackgroundMediaFetch(context.Context, *BackgroundMediaFetchRequest) (*BackgroundMediaFetchResponse, error)
	BackgroundMediaDelete(context.Context, *BackgroundMediaDeleteRequest) (*BackgroundMediaDeleteResponse, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMediaServiceServer struct{}

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
func (UnimplementedMediaServiceServer) BackgroundMediaUpload(context.Context, *BackgroundMediaUploadRequest) (*BackgroundMediaUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackgroundMediaUpload not implemented")
}
func (UnimplementedMediaServiceServer) BackgroundMediaFetch(context.Context, *BackgroundMediaFetchRequest) (*BackgroundMediaFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackgroundMediaFetch not implemented")
}
func (UnimplementedMediaServiceServer) BackgroundMediaDelete(context.Context, *BackgroundMediaDeleteRequest) (*BackgroundMediaDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackgroundMediaDelete not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}
func (UnimplementedMediaServiceServer) testEmbeddedByValue()                      {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	// If the following call pancis, it indicates UnimplementedMediaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: MediaService_PingApiMedia_FullMethodName,
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
		FullMethod: MediaService_PostUploadMedia_FullMethodName,
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
		FullMethod: MediaService_PostFetchMedia_FullMethodName,
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
		FullMethod: MediaService_PostDeleteMedia_FullMethodName,
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
		FullMethod: MediaService_PostsFetchMedia_FullMethodName,
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
		FullMethod: MediaService_AvatarMediaUpload_FullMethodName,
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
		FullMethod: MediaService_AvatarMediaFetch_FullMethodName,
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
		FullMethod: MediaService_AvatarMediaFetchMultiple_FullMethodName,
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
		FullMethod: MediaService_AvatarMediaDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AvatarMediaDelete(ctx, req.(*AvatarMediaDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_BackgroundMediaUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackgroundMediaUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).BackgroundMediaUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_BackgroundMediaUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).BackgroundMediaUpload(ctx, req.(*BackgroundMediaUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_BackgroundMediaFetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackgroundMediaFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).BackgroundMediaFetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_BackgroundMediaFetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).BackgroundMediaFetch(ctx, req.(*BackgroundMediaFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_BackgroundMediaDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackgroundMediaDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).BackgroundMediaDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_BackgroundMediaDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).BackgroundMediaDelete(ctx, req.(*BackgroundMediaDeleteRequest))
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
		{
			MethodName: "BackgroundMediaUpload",
			Handler:    _MediaService_BackgroundMediaUpload_Handler,
		},
		{
			MethodName: "BackgroundMediaFetch",
			Handler:    _MediaService_BackgroundMediaFetch_Handler,
		},
		{
			MethodName: "BackgroundMediaDelete",
			Handler:    _MediaService_BackgroundMediaDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-proto/media.proto",
}
