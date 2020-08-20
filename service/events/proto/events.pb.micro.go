// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/events/proto/events.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Stream service

func NewStreamEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Stream service

type StreamService interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Stream_SubscribeService, error)
}

type streamService struct {
	c    client.Client
	name string
}

func NewStreamService(name string, c client.Client) StreamService {
	return &streamService{
		c:    c,
		name: name,
	}
}

func (c *streamService) Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error) {
	req := c.c.NewRequest(c.name, "Stream.Publish", in)
	out := new(PublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamService) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Stream_SubscribeService, error) {
	req := c.c.NewRequest(c.name, "Stream.Subscribe", &SubscribeRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &streamServiceSubscribe{stream}, nil
}

type Stream_SubscribeService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Event, error)
}

type streamServiceSubscribe struct {
	stream client.Stream
}

func (x *streamServiceSubscribe) Close() error {
	return x.stream.Close()
}

func (x *streamServiceSubscribe) Context() context.Context {
	return x.stream.Context()
}

func (x *streamServiceSubscribe) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceSubscribe) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceSubscribe) Recv() (*Event, error) {
	m := new(Event)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Stream service

type StreamHandler interface {
	Publish(context.Context, *PublishRequest, *PublishResponse) error
	Subscribe(context.Context, *SubscribeRequest, Stream_SubscribeStream) error
}

func RegisterStreamHandler(s server.Server, hdlr StreamHandler, opts ...server.HandlerOption) error {
	type stream interface {
		Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error
		Subscribe(ctx context.Context, stream server.Stream) error
	}
	type Stream struct {
		stream
	}
	h := &streamHandler{hdlr}
	return s.Handle(s.NewHandler(&Stream{h}, opts...))
}

type streamHandler struct {
	StreamHandler
}

func (h *streamHandler) Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error {
	return h.StreamHandler.Publish(ctx, in, out)
}

func (h *streamHandler) Subscribe(ctx context.Context, stream server.Stream) error {
	m := new(SubscribeRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StreamHandler.Subscribe(ctx, m, &streamSubscribeStream{stream})
}

type Stream_SubscribeStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Event) error
}

type streamSubscribeStream struct {
	stream server.Stream
}

func (x *streamSubscribeStream) Close() error {
	return x.stream.Close()
}

func (x *streamSubscribeStream) Context() context.Context {
	return x.stream.Context()
}

func (x *streamSubscribeStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamSubscribeStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamSubscribeStream) Send(m *Event) error {
	return x.stream.Send(m)
}

// Api Endpoints for Store service

func NewStoreEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Store service

type StoreService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Store.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error) {
	req := c.c.NewRequest(c.name, "Store.Write", in)
	out := new(WriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Write(context.Context, *WriteRequest, *WriteResponse) error
}

func RegisterStoreHandler(s server.Server, hdlr StoreHandler, opts ...server.HandlerOption) error {
	type store interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error
	}
	type Store struct {
		store
	}
	h := &storeHandler{hdlr}
	return s.Handle(s.NewHandler(&Store{h}, opts...))
}

type storeHandler struct {
	StoreHandler
}

func (h *storeHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.StoreHandler.Read(ctx, in, out)
}

func (h *storeHandler) Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error {
	return h.StoreHandler.Write(ctx, in, out)
}
