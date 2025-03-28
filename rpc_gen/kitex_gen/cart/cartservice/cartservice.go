// Code generated by Kitex v0.9.1. DO NOT EDIT.

package cartservice

import (
	"context"
	"errors"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AddToCart": kitex.NewMethodInfo(
		addToCartHandler,
		newAddToCartArgs,
		newAddToCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetCart": kitex.NewMethodInfo(
		getCartHandler,
		newGetCartArgs,
		newGetCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RemoveFromCart": kitex.NewMethodInfo(
		removeFromCartHandler,
		newRemoveFromCartArgs,
		newRemoveFromCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateItemQuantity": kitex.NewMethodInfo(
		updateItemQuantityHandler,
		newUpdateItemQuantityArgs,
		newUpdateItemQuantityResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ClearCart": kitex.NewMethodInfo(
		clearCartHandler,
		newClearCartArgs,
		newClearCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	cartServiceServiceInfo                = NewServiceInfo()
	cartServiceServiceInfoForClient       = NewServiceInfoForClient()
	cartServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return cartServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return cartServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return cartServiceServiceInfoForClient
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
	serviceName := "CartService"
	handlerType := (*cart.CartService)(nil)
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
		"PackageName": "cart",
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

func addToCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.AddToCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).AddToCart(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddToCartArgs:
		success, err := handler.(cart.CartService).AddToCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddToCartResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddToCartArgs() interface{} {
	return &AddToCartArgs{}
}

func newAddToCartResult() interface{} {
	return &AddToCartResult{}
}

type AddToCartArgs struct {
	Req *cart.AddToCartRequest
}

func (p *AddToCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.AddToCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddToCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddToCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddToCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddToCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.AddToCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddToCartArgs_Req_DEFAULT *cart.AddToCartRequest

func (p *AddToCartArgs) GetReq() *cart.AddToCartRequest {
	if !p.IsSetReq() {
		return AddToCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddToCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddToCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddToCartResult struct {
	Success *cart.AddToCartResponse
}

var AddToCartResult_Success_DEFAULT *cart.AddToCartResponse

func (p *AddToCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.AddToCartResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddToCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddToCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddToCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddToCartResult) Unmarshal(in []byte) error {
	msg := new(cart.AddToCartResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddToCartResult) GetSuccess() *cart.AddToCartResponse {
	if !p.IsSetSuccess() {
		return AddToCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddToCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.AddToCartResponse)
}

func (p *AddToCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddToCartResult) GetResult() interface{} {
	return p.Success
}

func getCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.GetCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).GetCart(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetCartArgs:
		success, err := handler.(cart.CartService).GetCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCartResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetCartArgs() interface{} {
	return &GetCartArgs{}
}

func newGetCartResult() interface{} {
	return &GetCartResult{}
}

type GetCartArgs struct {
	Req *cart.GetCartRequest
}

func (p *GetCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.GetCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.GetCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCartArgs_Req_DEFAULT *cart.GetCartRequest

func (p *GetCartArgs) GetReq() *cart.GetCartRequest {
	if !p.IsSetReq() {
		return GetCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCartResult struct {
	Success *cart.GetCartResponse
}

var GetCartResult_Success_DEFAULT *cart.GetCartResponse

func (p *GetCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.GetCartResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCartResult) Unmarshal(in []byte) error {
	msg := new(cart.GetCartResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCartResult) GetSuccess() *cart.GetCartResponse {
	if !p.IsSetSuccess() {
		return GetCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.GetCartResponse)
}

func (p *GetCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCartResult) GetResult() interface{} {
	return p.Success
}

func removeFromCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.RemoveFromCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).RemoveFromCart(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RemoveFromCartArgs:
		success, err := handler.(cart.CartService).RemoveFromCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RemoveFromCartResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRemoveFromCartArgs() interface{} {
	return &RemoveFromCartArgs{}
}

func newRemoveFromCartResult() interface{} {
	return &RemoveFromCartResult{}
}

type RemoveFromCartArgs struct {
	Req *cart.RemoveFromCartRequest
}

func (p *RemoveFromCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.RemoveFromCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RemoveFromCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RemoveFromCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RemoveFromCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RemoveFromCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.RemoveFromCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RemoveFromCartArgs_Req_DEFAULT *cart.RemoveFromCartRequest

func (p *RemoveFromCartArgs) GetReq() *cart.RemoveFromCartRequest {
	if !p.IsSetReq() {
		return RemoveFromCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RemoveFromCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RemoveFromCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RemoveFromCartResult struct {
	Success *cart.RemoveFromCartResponse
}

var RemoveFromCartResult_Success_DEFAULT *cart.RemoveFromCartResponse

func (p *RemoveFromCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.RemoveFromCartResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RemoveFromCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RemoveFromCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RemoveFromCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RemoveFromCartResult) Unmarshal(in []byte) error {
	msg := new(cart.RemoveFromCartResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RemoveFromCartResult) GetSuccess() *cart.RemoveFromCartResponse {
	if !p.IsSetSuccess() {
		return RemoveFromCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RemoveFromCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.RemoveFromCartResponse)
}

func (p *RemoveFromCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RemoveFromCartResult) GetResult() interface{} {
	return p.Success
}

func updateItemQuantityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.UpdateQuantityRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).UpdateItemQuantity(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateItemQuantityArgs:
		success, err := handler.(cart.CartService).UpdateItemQuantity(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateItemQuantityResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateItemQuantityArgs() interface{} {
	return &UpdateItemQuantityArgs{}
}

func newUpdateItemQuantityResult() interface{} {
	return &UpdateItemQuantityResult{}
}

type UpdateItemQuantityArgs struct {
	Req *cart.UpdateQuantityRequest
}

func (p *UpdateItemQuantityArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.UpdateQuantityRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateItemQuantityArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateItemQuantityArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateItemQuantityArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateItemQuantityArgs) Unmarshal(in []byte) error {
	msg := new(cart.UpdateQuantityRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateItemQuantityArgs_Req_DEFAULT *cart.UpdateQuantityRequest

func (p *UpdateItemQuantityArgs) GetReq() *cart.UpdateQuantityRequest {
	if !p.IsSetReq() {
		return UpdateItemQuantityArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateItemQuantityArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateItemQuantityArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateItemQuantityResult struct {
	Success *cart.UpdateQuantityResponse
}

var UpdateItemQuantityResult_Success_DEFAULT *cart.UpdateQuantityResponse

func (p *UpdateItemQuantityResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.UpdateQuantityResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateItemQuantityResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateItemQuantityResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateItemQuantityResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateItemQuantityResult) Unmarshal(in []byte) error {
	msg := new(cart.UpdateQuantityResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateItemQuantityResult) GetSuccess() *cart.UpdateQuantityResponse {
	if !p.IsSetSuccess() {
		return UpdateItemQuantityResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateItemQuantityResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.UpdateQuantityResponse)
}

func (p *UpdateItemQuantityResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateItemQuantityResult) GetResult() interface{} {
	return p.Success
}

func clearCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(cart.ClearCartRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(cart.CartService).ClearCart(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ClearCartArgs:
		success, err := handler.(cart.CartService).ClearCart(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ClearCartResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newClearCartArgs() interface{} {
	return &ClearCartArgs{}
}

func newClearCartResult() interface{} {
	return &ClearCartResult{}
}

type ClearCartArgs struct {
	Req *cart.ClearCartRequest
}

func (p *ClearCartArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(cart.ClearCartRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ClearCartArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ClearCartArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ClearCartArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ClearCartArgs) Unmarshal(in []byte) error {
	msg := new(cart.ClearCartRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ClearCartArgs_Req_DEFAULT *cart.ClearCartRequest

func (p *ClearCartArgs) GetReq() *cart.ClearCartRequest {
	if !p.IsSetReq() {
		return ClearCartArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ClearCartArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ClearCartArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ClearCartResult struct {
	Success *cart.ClearCartResponse
}

var ClearCartResult_Success_DEFAULT *cart.ClearCartResponse

func (p *ClearCartResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(cart.ClearCartResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ClearCartResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ClearCartResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ClearCartResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ClearCartResult) Unmarshal(in []byte) error {
	msg := new(cart.ClearCartResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ClearCartResult) GetSuccess() *cart.ClearCartResponse {
	if !p.IsSetSuccess() {
		return ClearCartResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ClearCartResult) SetSuccess(x interface{}) {
	p.Success = x.(*cart.ClearCartResponse)
}

func (p *ClearCartResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ClearCartResult) GetResult() interface{} {
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

func (p *kClient) AddToCart(ctx context.Context, Req *cart.AddToCartRequest) (r *cart.AddToCartResponse, err error) {
	var _args AddToCartArgs
	_args.Req = Req
	var _result AddToCartResult
	if err = p.c.Call(ctx, "AddToCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCart(ctx context.Context, Req *cart.GetCartRequest) (r *cart.GetCartResponse, err error) {
	var _args GetCartArgs
	_args.Req = Req
	var _result GetCartResult
	if err = p.c.Call(ctx, "GetCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveFromCart(ctx context.Context, Req *cart.RemoveFromCartRequest) (r *cart.RemoveFromCartResponse, err error) {
	var _args RemoveFromCartArgs
	_args.Req = Req
	var _result RemoveFromCartResult
	if err = p.c.Call(ctx, "RemoveFromCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateItemQuantity(ctx context.Context, Req *cart.UpdateQuantityRequest) (r *cart.UpdateQuantityResponse, err error) {
	var _args UpdateItemQuantityArgs
	_args.Req = Req
	var _result UpdateItemQuantityResult
	if err = p.c.Call(ctx, "UpdateItemQuantity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ClearCart(ctx context.Context, Req *cart.ClearCartRequest) (r *cart.ClearCartResponse, err error) {
	var _args ClearCartArgs
	_args.Req = Req
	var _result ClearCartResult
	if err = p.c.Call(ctx, "ClearCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
