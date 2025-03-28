package model

import "gorm.io/gorm"

type OrderStatus int

const (
	OrderStatusUnknown       OrderStatus = 0 // 未知状态
	OrderStatusPending       OrderStatus = 1 // 待支付
	OrderStatusCompleted     OrderStatus = 2 // 已完成
	OrderStatusCancelled     OrderStatus = 3 // 已取消
	OrderStatusPaymentFailed OrderStatus = 4 // 支付失败
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