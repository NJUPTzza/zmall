// Code generated by Kitex v0.9.1. DO NOT EDIT.

package agentservice

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (r *agent.ChatResponse, err error)
	StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (stream AgentService_StreamMessageClient, err error)
	GetHistory(ctx context.Context, Req *agent.GetHistoryRequest, callOptions ...callopt.Option) (r *agent.GetHistoryResponse, err error)
	DeleteHistory(ctx context.Context, Req *agent.DeleteHistoryRequest, callOptions ...callopt.Option) (r *agent.DeleteHistoryResponse, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...streamcall.Option) (stream AgentService_StreamMessageClient, err error)
}

type AgentService_StreamMessageClient interface {
	streaming.Stream
	Recv() (*agent.ChatResponse, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAgentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAgentServiceClient struct {
	*kClient
}

func (p *kAgentServiceClient) SendMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (r *agent.ChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, Req)
}

func (p *kAgentServiceClient) StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (stream AgentService_StreamMessageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StreamMessage(ctx, Req)
}

func (p *kAgentServiceClient) GetHistory(ctx context.Context, Req *agent.GetHistoryRequest, callOptions ...callopt.Option) (r *agent.GetHistoryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetHistory(ctx, Req)
}

func (p *kAgentServiceClient) DeleteHistory(ctx context.Context, Req *agent.DeleteHistoryRequest, callOptions ...callopt.Option) (r *agent.DeleteHistoryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteHistory(ctx, Req)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kAgentServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAgentServiceStreamClient struct {
	*kClient
}

func (p *kAgentServiceStreamClient) StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...streamcall.Option) (stream AgentService_StreamMessageClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.StreamMessage(ctx, Req)
}
