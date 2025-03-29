package service

import (
	"context"

	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
)

type UpdateStockService struct {
	ctx context.Context
} // NewUpdateStockService new UpdateStockService
func NewUpdateStockService(ctx context.Context) *UpdateStockService {
	return &UpdateStockService{ctx: ctx}
}

// Run create note info
func (s *UpdateStockService) Run(req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	resp = &product.UpdateStockResponse{CommonResponse: &product.CommonResponse{}}

	// 校验请求参数
	if req.ProductId <= 0 || req.Change <= 0 {
		resp.CommonResponse.Code = 400
		resp.CommonResponse.Message = "请求参数错误"
		return resp, nil
	}

	// 执行库存更新
	err = mysql.UpdateStock(mysql.DB, req.ProductId, req.Change) 
	if err != nil {
		resp.CommonResponse.Code = 500
		resp.CommonResponse.Message = "库存更新失败"
		return resp, err
	}

	// 成功
	resp.CommonResponse.Code = 200
	resp.CommonResponse.Message = "库存更新成功"
	return resp, nil
}
