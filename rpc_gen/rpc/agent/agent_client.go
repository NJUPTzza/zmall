package agent

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent/agentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() agentservice.Client
	Service() string
	SendMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (r *agent.ChatResponse, err error)
	StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (stream agentservice.AgentService_StreamMessageClient, err error)
	GetHistory(ctx context.Context, Req *agent.GetHistoryRequest, callOptions ...callopt.Option) (r *agent.GetHistoryResponse, err error)
	DeleteHistory(ctx context.Context, Req *agent.DeleteHistoryRequest, callOptions ...callopt.Option) (r *agent.DeleteHistoryResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := agentservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient agentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() agentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SendMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (r *agent.ChatResponse, err error) {
	return c.kitexClient.SendMessage(ctx, Req, callOptions...)
}

func (c *clientImpl) StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (stream agentservice.AgentService_StreamMessageClient, err error) {
	return c.kitexClient.StreamMessage(ctx, Req, callOptions...)
}

func (c *clientImpl) GetHistory(ctx context.Context, Req *agent.GetHistoryRequest, callOptions ...callopt.Option) (r *agent.GetHistoryResponse, err error) {
	return c.kitexClient.GetHistory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteHistory(ctx context.Context, Req *agent.DeleteHistoryRequest, callOptions ...callopt.Option) (r *agent.DeleteHistoryResponse, err error) {
	return c.kitexClient.DeleteHistory(ctx, Req, callOptions...)
}
