package model

import "gorm.io/gorm"

type OrderStatus int

const (
	OrderStatusUnknown       OrderStatus = 0 // 未知状态
	OrderStatusPending       OrderStatus = 1 // 待支付
	OrderStatusProcessing    OrderStatus = 2 // 处理中，如支付中，发货中等
	OrderStatusCompleted     OrderStatus = 3 // 订单完成
	OrderStatusFailed        OrderStatus = 4 // 失败或已取消
	OrderStatusShipped       OrderStatus = 5 // 已发货
	OrderStatusDelivered     OrderStatus = 6 // 已签收
	OrderStatusRefunded      OrderStatus = 7 // 已退款
)


type Order struct {
	gorm.Model
	UserId int64 `gorm:"type:int;not null"`
	ProductId int64 `gorm:"type:int;not null"`
	Quantity int32 `gorm:"type:int;not null;default:1"`
	TotalPrice float32 `gorm:"type:decimal(10,2);not null"`
	Status OrderStatus `gorm:"type:tinyint;not null;default:1"`
}

func (Order) TableName() string {
	return "order"
}