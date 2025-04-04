// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	"errors"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProcessPayment": kitex.NewMethodInfo(
		processPaymentHandler,
		newProcessPaymentArgs,
		newProcessPaymentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetPaymentStatus": kitex.NewMethodInfo(
		getPaymentStatusHandler,
		newGetPaymentStatusArgs,
		newGetPaymentStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdatePaymentStatus": kitex.NewMethodInfo(
		updatePaymentStatusHandler,
		newUpdatePaymentStatusArgs,
		newUpdatePaymentStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	paymentServiceServiceInfo                = NewServiceInfo()
	paymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	paymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func processPaymentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.ProcessPaymentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).ProcessPayment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProcessPaymentArgs:
		success, err := handler.(payment.PaymentService).ProcessPayment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProcessPaymentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProcessPaymentArgs() interface{} {
	return &ProcessPaymentArgs{}
}

func newProcessPaymentResult() interface{} {
	return &ProcessPaymentResult{}
}

type ProcessPaymentArgs struct {
	Req *payment.ProcessPaymentRequest
}

func (p *ProcessPaymentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.ProcessPaymentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProcessPaymentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProcessPaymentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProcessPaymentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProcessPaymentArgs) Unmarshal(in []byte) error {
	msg := new(payment.ProcessPaymentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProcessPaymentArgs_Req_DEFAULT *payment.ProcessPaymentRequest

func (p *ProcessPaymentArgs) GetReq() *payment.ProcessPaymentRequest {
	if !p.IsSetReq() {
		return ProcessPaymentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProcessPaymentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProcessPaymentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProcessPaymentResult struct {
	Success *payment.ProcessPaymentResponse
}

var ProcessPaymentResult_Success_DEFAULT *payment.ProcessPaymentResponse

func (p *ProcessPaymentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.ProcessPaymentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProcessPaymentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProcessPaymentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProcessPaymentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProcessPaymentResult) Unmarshal(in []byte) error {
	msg := new(payment.ProcessPaymentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProcessPaymentResult) GetSuccess() *payment.ProcessPaymentResponse {
	if !p.IsSetSuccess() {
		return ProcessPaymentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProcessPaymentResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.ProcessPaymentResponse)
}

func (p *ProcessPaymentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProcessPaymentResult) GetResult() interface{} {
	return p.Success
}

func getPaymentStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.GetPaymentStatusRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).GetPaymentStatus(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetPaymentStatusArgs:
		success, err := handler.(payment.PaymentService).GetPaymentStatus(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPaymentStatusResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetPaymentStatusArgs() interface{} {
	return &GetPaymentStatusArgs{}
}

func newGetPaymentStatusResult() interface{} {
	return &GetPaymentStatusResult{}
}

type GetPaymentStatusArgs struct {
	Req *payment.GetPaymentStatusRequest
}

func (p *GetPaymentStatusArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.GetPaymentStatusRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPaymentStatusArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPaymentStatusArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPaymentStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPaymentStatusArgs) Unmarshal(in []byte) error {
	msg := new(payment.GetPaymentStatusRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPaymentStatusArgs_Req_DEFAULT *payment.GetPaymentStatusRequest

func (p *GetPaymentStatusArgs) GetReq() *payment.GetPaymentStatusRequest {
	if !p.IsSetReq() {
		return GetPaymentStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPaymentStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPaymentStatusArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPaymentStatusResult struct {
	Success *payment.GetPaymentStatusResponse
}

var GetPaymentStatusResult_Success_DEFAULT *payment.GetPaymentStatusResponse

func (p *GetPaymentStatusResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.GetPaymentStatusResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPaymentStatusResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPaymentStatusResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPaymentStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPaymentStatusResult) Unmarshal(in []byte) error {
	msg := new(payment.GetPaymentStatusResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPaymentStatusResult) GetSuccess() *payment.GetPaymentStatusResponse {
	if !p.IsSetSuccess() {
		return GetPaymentStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPaymentStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.GetPaymentStatusResponse)
}

func (p *GetPaymentStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPaymentStatusResult) GetResult() interface{} {
	return p.Success
}

func updatePaymentStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.UpdatePaymentStatusRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).UpdatePaymentStatus(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdatePaymentStatusArgs:
		success, err := handler.(payment.PaymentService).UpdatePaymentStatus(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdatePaymentStatusResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdatePaymentStatusArgs() interface{} {
	return &UpdatePaymentStatusArgs{}
}

func newUpdatePaymentStatusResult() interface{} {
	return &UpdatePaymentStatusResult{}
}

type UpdatePaymentStatusArgs struct {
	Req *payment.UpdatePaymentStatusRequest
}

func (p *UpdatePaymentStatusArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.UpdatePaymentStatusRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdatePaymentStatusArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdatePaymentStatusArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdatePaymentStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdatePaymentStatusArgs) Unmarshal(in []byte) error {
	msg := new(payment.UpdatePaymentStatusRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdatePaymentStatusArgs_Req_DEFAULT *payment.UpdatePaymentStatusRequest

func (p *UpdatePaymentStatusArgs) GetReq() *payment.UpdatePaymentStatusRequest {
	if !p.IsSetReq() {
		return UpdatePaymentStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePaymentStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdatePaymentStatusArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdatePaymentStatusResult struct {
	Success *payment.UpdatePaymentStatusResponse
}

var UpdatePaymentStatusResult_Success_DEFAULT *payment.UpdatePaymentStatusResponse

func (p *UpdatePaymentStatusResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.UpdatePaymentStatusResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdatePaymentStatusResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdatePaymentStatusResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdatePaymentStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdatePaymentStatusResult) Unmarshal(in []byte) error {
	msg := new(payment.UpdatePaymentStatusResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePaymentStatusResult) GetSuccess() *payment.UpdatePaymentStatusResponse {
	if !p.IsSetSuccess() {
		return UpdatePaymentStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePaymentStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.UpdatePaymentStatusResponse)
}

func (p *UpdatePaymentStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePaymentStatusResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ProcessPayment(ctx context.Context, Req *payment.ProcessPaymentRequest) (r *payment.ProcessPaymentResponse, err error) {
	var _args ProcessPaymentArgs
	_args.Req = Req
	var _result ProcessPaymentResult
	if err = p.c.Call(ctx, "ProcessPayment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPaymentStatus(ctx context.Context, Req *payment.GetPaymentStatusRequest) (r *payment.GetPaymentStatusResponse, err error) {
	var _args GetPaymentStatusArgs
	_args.Req = Req
	var _result GetPaymentStatusResult
	if err = p.c.Call(ctx, "GetPaymentStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePaymentStatus(ctx context.Context, Req *payment.UpdatePaymentStatusRequest) (r *payment.UpdatePaymentStatusResponse, err error) {
	var _args UpdatePaymentStatusArgs
	_args.Req = Req
	var _result UpdatePaymentStatusResult
	if err = p.c.Call(ctx, "UpdatePaymentStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
