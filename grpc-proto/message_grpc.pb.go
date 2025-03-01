// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: grpc-proto/message.proto

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
	MessageService_PingApiMessage_FullMethodName                  = "/message.MessageService/PingApiMessage"
	MessageService_CustomerMessage_FullMethodName                 = "/message.MessageService/CustomerMessage"
	MessageService_ChatbotMessage_FullMethodName                  = "/message.MessageService/ChatbotMessage"
	MessageService_PingTasksQueue_FullMethodName                  = "/message.MessageService/PingTasksQueue"
	MessageService_CreateTaskMessage_FullMethodName               = "/message.MessageService/CreateTaskMessage"
	MessageService_CompleteTaskMessage_FullMethodName             = "/message.MessageService/CompleteTaskMessage"
	MessageService_UserSendMessage_FullMethodName                 = "/message.MessageService/UserSendMessage"
	MessageService_EditMessage_FullMethodName                     = "/message.MessageService/EditMessage"
	MessageService_FetchMessage_FullMethodName                    = "/message.MessageService/FetchMessage"
	MessageService_DeleteMessage_FullMethodName                   = "/message.MessageService/DeleteMessage"
	MessageService_ReadMessage_FullMethodName                     = "/message.MessageService/ReadMessage"
	MessageService_FetchConversationMessages_FullMethodName       = "/message.MessageService/FetchConversationMessages"
	MessageService_DeleteConversation_FullMethodName              = "/message.MessageService/DeleteConversation"
	MessageService_FetchAllUserConversations_FullMethodName       = "/message.MessageService/FetchAllUserConversations"
	MessageService_FetchConversationMessageCounter_FullMethodName = "/message.MessageService/FetchConversationMessageCounter"
	MessageService_StoreConversationMessageCounter_FullMethodName = "/message.MessageService/StoreConversationMessageCounter"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	PingApiMessage(ctx context.Context, in *PingApiMessageRequest, opts ...grpc.CallOption) (*PingApiMessageResponse, error)
	CustomerMessage(ctx context.Context, in *CustomerMessageRequest, opts ...grpc.CallOption) (*CustomerMessageResponse, error)
	ChatbotMessage(ctx context.Context, in *ChatbotMessageRequest, opts ...grpc.CallOption) (*ChatbotMessageResponse, error)
	PingTasksQueue(ctx context.Context, in *PingTasksQueueRequest, opts ...grpc.CallOption) (*PingTasksQueueResponse, error)
	CreateTaskMessage(ctx context.Context, in *CreateTaskMessageRequest, opts ...grpc.CallOption) (*CreateTaskMessageResponse, error)
	CompleteTaskMessage(ctx context.Context, in *CompleteTaskMessageRequest, opts ...grpc.CallOption) (*CompleteTaskMessageResponse, error)
	UserSendMessage(ctx context.Context, in *UserSendMessageRequest, opts ...grpc.CallOption) (*UserSendMessageResponse, error)
	EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error)
	FetchMessage(ctx context.Context, in *FetchMessageRequest, opts ...grpc.CallOption) (*FetchMessageResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (*ReadMessageResponse, error)
	FetchConversationMessages(ctx context.Context, in *FetchConversationMessagesRequest, opts ...grpc.CallOption) (*FetchConversationMessagesResponse, error)
	DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error)
	FetchAllUserConversations(ctx context.Context, in *FetchAllUserConversationsRequest, opts ...grpc.CallOption) (*FetchAllUserConversationsResponse, error)
	FetchConversationMessageCounter(ctx context.Context, in *FetchConversationMessageCounterRequest, opts ...grpc.CallOption) (*FetchConversationMessageCounterResponse, error)
	StoreConversationMessageCounter(ctx context.Context, in *StoreConversationCounterRequest, opts ...grpc.CallOption) (*StoreConversationCounterResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) PingApiMessage(ctx context.Context, in *PingApiMessageRequest, opts ...grpc.CallOption) (*PingApiMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingApiMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_PingApiMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) CustomerMessage(ctx context.Context, in *CustomerMessageRequest, opts ...grpc.CallOption) (*CustomerMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CustomerMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_CustomerMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ChatbotMessage(ctx context.Context, in *ChatbotMessageRequest, opts ...grpc.CallOption) (*ChatbotMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatbotMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_ChatbotMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) PingTasksQueue(ctx context.Context, in *PingTasksQueueRequest, opts ...grpc.CallOption) (*PingTasksQueueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingTasksQueueResponse)
	err := c.cc.Invoke(ctx, MessageService_PingTasksQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) CreateTaskMessage(ctx context.Context, in *CreateTaskMessageRequest, opts ...grpc.CallOption) (*CreateTaskMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_CreateTaskMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) CompleteTaskMessage(ctx context.Context, in *CompleteTaskMessageRequest, opts ...grpc.CallOption) (*CompleteTaskMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteTaskMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_CompleteTaskMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UserSendMessage(ctx context.Context, in *UserSendMessageRequest, opts ...grpc.CallOption) (*UserSendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSendMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_UserSendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_EditMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) FetchMessage(ctx context.Context, in *FetchMessageRequest, opts ...grpc.CallOption) (*FetchMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_FetchMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (*ReadMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadMessageResponse)
	err := c.cc.Invoke(ctx, MessageService_ReadMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) FetchConversationMessages(ctx context.Context, in *FetchConversationMessagesRequest, opts ...grpc.CallOption) (*FetchConversationMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchConversationMessagesResponse)
	err := c.cc.Invoke(ctx, MessageService_FetchConversationMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConversationResponse)
	err := c.cc.Invoke(ctx, MessageService_DeleteConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) FetchAllUserConversations(ctx context.Context, in *FetchAllUserConversationsRequest, opts ...grpc.CallOption) (*FetchAllUserConversationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchAllUserConversationsResponse)
	err := c.cc.Invoke(ctx, MessageService_FetchAllUserConversations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) FetchConversationMessageCounter(ctx context.Context, in *FetchConversationMessageCounterRequest, opts ...grpc.CallOption) (*FetchConversationMessageCounterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchConversationMessageCounterResponse)
	err := c.cc.Invoke(ctx, MessageService_FetchConversationMessageCounter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) StoreConversationMessageCounter(ctx context.Context, in *StoreConversationCounterRequest, opts ...grpc.CallOption) (*StoreConversationCounterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreConversationCounterResponse)
	err := c.cc.Invoke(ctx, MessageService_StoreConversationMessageCounter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility.
type MessageServiceServer interface {
	PingApiMessage(context.Context, *PingApiMessageRequest) (*PingApiMessageResponse, error)
	CustomerMessage(context.Context, *CustomerMessageRequest) (*CustomerMessageResponse, error)
	ChatbotMessage(context.Context, *ChatbotMessageRequest) (*ChatbotMessageResponse, error)
	PingTasksQueue(context.Context, *PingTasksQueueRequest) (*PingTasksQueueResponse, error)
	CreateTaskMessage(context.Context, *CreateTaskMessageRequest) (*CreateTaskMessageResponse, error)
	CompleteTaskMessage(context.Context, *CompleteTaskMessageRequest) (*CompleteTaskMessageResponse, error)
	UserSendMessage(context.Context, *UserSendMessageRequest) (*UserSendMessageResponse, error)
	EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error)
	FetchMessage(context.Context, *FetchMessageRequest) (*FetchMessageResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	ReadMessage(context.Context, *ReadMessageRequest) (*ReadMessageResponse, error)
	FetchConversationMessages(context.Context, *FetchConversationMessagesRequest) (*FetchConversationMessagesResponse, error)
	DeleteConversation(context.Context, *DeleteConversationRequest) (*DeleteConversationResponse, error)
	FetchAllUserConversations(context.Context, *FetchAllUserConversationsRequest) (*FetchAllUserConversationsResponse, error)
	FetchConversationMessageCounter(context.Context, *FetchConversationMessageCounterRequest) (*FetchConversationMessageCounterResponse, error)
	StoreConversationMessageCounter(context.Context, *StoreConversationCounterRequest) (*StoreConversationCounterResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageServiceServer struct{}

func (UnimplementedMessageServiceServer) PingApiMessage(context.Context, *PingApiMessageRequest) (*PingApiMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingApiMessage not implemented")
}
func (UnimplementedMessageServiceServer) CustomerMessage(context.Context, *CustomerMessageRequest) (*CustomerMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerMessage not implemented")
}
func (UnimplementedMessageServiceServer) ChatbotMessage(context.Context, *ChatbotMessageRequest) (*ChatbotMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatbotMessage not implemented")
}
func (UnimplementedMessageServiceServer) PingTasksQueue(context.Context, *PingTasksQueueRequest) (*PingTasksQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingTasksQueue not implemented")
}
func (UnimplementedMessageServiceServer) CreateTaskMessage(context.Context, *CreateTaskMessageRequest) (*CreateTaskMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskMessage not implemented")
}
func (UnimplementedMessageServiceServer) CompleteTaskMessage(context.Context, *CompleteTaskMessageRequest) (*CompleteTaskMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTaskMessage not implemented")
}
func (UnimplementedMessageServiceServer) UserSendMessage(context.Context, *UserSendMessageRequest) (*UserSendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSendMessage not implemented")
}
func (UnimplementedMessageServiceServer) EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedMessageServiceServer) FetchMessage(context.Context, *FetchMessageRequest) (*FetchMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMessage not implemented")
}
func (UnimplementedMessageServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessageServiceServer) ReadMessage(context.Context, *ReadMessageRequest) (*ReadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMessage not implemented")
}
func (UnimplementedMessageServiceServer) FetchConversationMessages(context.Context, *FetchConversationMessagesRequest) (*FetchConversationMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchConversationMessages not implemented")
}
func (UnimplementedMessageServiceServer) DeleteConversation(context.Context, *DeleteConversationRequest) (*DeleteConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConversation not implemented")
}
func (UnimplementedMessageServiceServer) FetchAllUserConversations(context.Context, *FetchAllUserConversationsRequest) (*FetchAllUserConversationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAllUserConversations not implemented")
}
func (UnimplementedMessageServiceServer) FetchConversationMessageCounter(context.Context, *FetchConversationMessageCounterRequest) (*FetchConversationMessageCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchConversationMessageCounter not implemented")
}
func (UnimplementedMessageServiceServer) StoreConversationMessageCounter(context.Context, *StoreConversationCounterRequest) (*StoreConversationCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreConversationMessageCounter not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}
func (UnimplementedMessageServiceServer) testEmbeddedByValue()                        {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_PingApiMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingApiMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).PingApiMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_PingApiMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).PingApiMessage(ctx, req.(*PingApiMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_CustomerMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CustomerMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_CustomerMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CustomerMessage(ctx, req.(*CustomerMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ChatbotMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatbotMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ChatbotMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_ChatbotMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ChatbotMessage(ctx, req.(*ChatbotMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_PingTasksQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingTasksQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).PingTasksQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_PingTasksQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).PingTasksQueue(ctx, req.(*PingTasksQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_CreateTaskMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateTaskMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_CreateTaskMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateTaskMessage(ctx, req.(*CreateTaskMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_CompleteTaskMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CompleteTaskMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_CompleteTaskMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CompleteTaskMessage(ctx, req.(*CompleteTaskMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UserSendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UserSendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_UserSendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UserSendMessage(ctx, req.(*UserSendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_EditMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EditMessage(ctx, req.(*EditMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_FetchMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).FetchMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_FetchMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).FetchMessage(ctx, req.(*FetchMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ReadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ReadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_ReadMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ReadMessage(ctx, req.(*ReadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_FetchConversationMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchConversationMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).FetchConversationMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_FetchConversationMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).FetchConversationMessages(ctx, req.(*FetchConversationMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_DeleteConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteConversation(ctx, req.(*DeleteConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_FetchAllUserConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAllUserConversationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).FetchAllUserConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_FetchAllUserConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).FetchAllUserConversations(ctx, req.(*FetchAllUserConversationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_FetchConversationMessageCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchConversationMessageCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).FetchConversationMessageCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_FetchConversationMessageCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).FetchConversationMessageCounter(ctx, req.(*FetchConversationMessageCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_StoreConversationMessageCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreConversationCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).StoreConversationMessageCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_StoreConversationMessageCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).StoreConversationMessageCounter(ctx, req.(*StoreConversationCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingApiMessage",
			Handler:    _MessageService_PingApiMessage_Handler,
		},
		{
			MethodName: "CustomerMessage",
			Handler:    _MessageService_CustomerMessage_Handler,
		},
		{
			MethodName: "ChatbotMessage",
			Handler:    _MessageService_ChatbotMessage_Handler,
		},
		{
			MethodName: "PingTasksQueue",
			Handler:    _MessageService_PingTasksQueue_Handler,
		},
		{
			MethodName: "CreateTaskMessage",
			Handler:    _MessageService_CreateTaskMessage_Handler,
		},
		{
			MethodName: "CompleteTaskMessage",
			Handler:    _MessageService_CompleteTaskMessage_Handler,
		},
		{
			MethodName: "UserSendMessage",
			Handler:    _MessageService_UserSendMessage_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _MessageService_EditMessage_Handler,
		},
		{
			MethodName: "FetchMessage",
			Handler:    _MessageService_FetchMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessageService_DeleteMessage_Handler,
		},
		{
			MethodName: "ReadMessage",
			Handler:    _MessageService_ReadMessage_Handler,
		},
		{
			MethodName: "FetchConversationMessages",
			Handler:    _MessageService_FetchConversationMessages_Handler,
		},
		{
			MethodName: "DeleteConversation",
			Handler:    _MessageService_DeleteConversation_Handler,
		},
		{
			MethodName: "FetchAllUserConversations",
			Handler:    _MessageService_FetchAllUserConversations_Handler,
		},
		{
			MethodName: "FetchConversationMessageCounter",
			Handler:    _MessageService_FetchConversationMessageCounter_Handler,
		},
		{
			MethodName: "StoreConversationMessageCounter",
			Handler:    _MessageService_StoreConversationMessageCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-proto/message.proto",
}
