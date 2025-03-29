package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/cart/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/cart/biz/errno"
	"github.com/NJUPTzza/zmall/app/cart/biz/model"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
)

type AddToCartService struct {
	ctx context.Context
	productClient productservice.Client
} 

// NewAddToCartService new AddToCartService
func NewAddToCartService(ctx context.Context, productClient productservice.Client) *AddToCartService {
	return &AddToCartService{
		ctx: ctx,
		productClient: productClient,
	}
}

// Run create note info
func (s *AddToCartService) Run(req *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {

	// -----------------------
	// 1. 校验参数是否合法
	// -----------------------
	if req.UserId == 0 || req.ProductId == 0 || req.Quantity == 0 {
		return nil, errors.New("非法参数")
	}

	// -----------------------
	// 2. 调用商品服务查询商品详情
	// -----------------------
	pReq := &product.GetProductRequest {
		Id: req.ProductId,
	}
	pResp, err := s.productClient.GetProduct(s.ctx, pReq)
	if err != nil {
		return nil, errors.New("商品服务调用失败: " + err.Error())
	}
	if pResp.Product == nil {
		return nil, errors.New("商品不存在")
	}
	if pResp.Product.Stock < req.Quantity {
		return nil, errno.NewStockNotEnough(int(pResp.Product.Stock))
	}

	// -----------------------
	// 3. 插入数据库
	// -----------------------
	newCartItem := &model.Cart{
		UserId: req.UserId,
		ProductId: req.ProductId,
		Quantity: req.Quantity,
	}
	err = mysql.Create(mysql.DB, newCartItem)
	if err != nil {
		return nil, err
	}
	
	// -----------------------
	// 4. 返回信息
	// -----------------------
	return &cart.AddToCartResponse{
		CommonResponse: &cart.CommonResponse{
			Code: 200,
			Message: "添加到购物车成功",
		},
		Item: &cart.CartItem{
			ProductId: req.ProductId,
			Quantity: req.Quantity,
		},
	}, nil
}
