package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/order/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/order/biz/dal/redis"
	"github.com/NJUPTzza/zmall/app/order/biz/model"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
)

type CreateOrderService struct {
	ctx context.Context
	productClient productservice.Client
} 

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context, productClient productservice.Client) *CreateOrderService {
	return &CreateOrderService{
		ctx: ctx,
		productClient: productClient,
	}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	// 1. 查询商品信息
	pResp, err := s.productClient.GetProduct(s.ctx, &product.GetProductRequest{
		Id: req.ProductId,
	}) 
	if err != nil {
		return nil, errors.New("商品服务调用失败: " + err.Error())
	}
	if pResp.Product == nil {
		return nil, errors.New("商品不存在")
	}

	// 2. 检查库存是否足够
	stock, err := redis.GetProductStock(s.ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	if stock < req.Quantity {
		return &order.CreateOrderResponse{
			CommonResponse: &order.CommonResponse{Code: 400, Message: "库存不足"},
		}, nil
	}

	// 3.预扣减库存
	ok, err := redis.DecrProductStock(s.ctx, req.ProductId, req.Quantity)
	if err!= nil {
		return nil, err
	}
	if !ok {
		return &order.CreateOrderResponse{
			CommonResponse: &order.CommonResponse{Code: 400, Message: "库存不足"},
		}, nil
	}

	// 4. 创建订单逻辑...
	newOrder := &model.Order{
		UserId: req.UserId,
		ProductId: req.ProductId,
		Quantity: req.Quantity,
		TotalPrice: float32(req.Quantity) * pResp.Product.Price,
		Status: model.OrderStatusPending,
	}
	err = mysql.Create(mysql.DB, newOrder)
	if err != nil {
		// 订单创建失败，回滚 Redis 预扣库存
		redis.IncrProductStock(s.ctx, req.ProductId, req.Quantity)
		return nil, errors.New("订单创建失败: " + err.Error())
	}

	return &order.CreateOrderResponse{
		CommonResponse: &order.CommonResponse{Code: 200, Message: "订单创建成功"},
	}, nil
}
