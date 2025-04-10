package agent

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent/agentservice"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendMessage(ctx context.Context, req *agent.ChatRequest, callOptions ...callopt.Option) (resp *agent.ChatResponse, err error) {
	resp, err = defaultClient.SendMessage(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendMessage call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func StreamMessage(ctx context.Context, Req *agent.ChatRequest, callOptions ...callopt.Option) (stream agentservice.AgentService_StreamMessageClient, err error) {
	stream, err = defaultClient.StreamMessage(ctx, Req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "StreamMessage call failed,err =%+v", err)
		return nil, err
	}
	return stream, nil
}

func GetHistory(ctx context.Context, req *agent.GetHistoryRequest, callOptions ...callopt.Option) (resp *agent.GetHistoryResponse, err error) {
	resp, err = defaultClient.GetHistory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetHistory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteHistory(ctx context.Context, req *agent.DeleteHistoryRequest, callOptions ...callopt.Option) (resp *agent.DeleteHistoryResponse, err error) {
	resp, err = defaultClient.DeleteHistory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteHistory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
