package model

import "gorm.io/gorm"

type PaymentStatus int

const (
	PaymentStatusUnknown  PaymentStatus = 0 // 未知
	PaymentStatusPending  PaymentStatus = 1 // 待支付
	PaymentStatusPaid     PaymentStatus = 2 // 已支付
	PaymentStatusFailed   PaymentStatus = 3 // 支付失败
	PaymentStatusRefunded PaymentStatus = 4 // 已退款
)

type Payment struct {
	gorm.Model
	OrderId int64 `gorm:"type:int;not null"`
	UserId int64 `gorm:"type:int;not null"`
	Amount float32 `gorm:"type:decimal(10,2);not null"`      // 支付金额
	Status PaymentStatus `gorm:"type:tinyint;not null;default:1"` 
	PayChannel string `gorm:"type:varchar(20);not null"` // 支付状态
	PayURL string `gorm:"type:text;not null"`
}

func (Payment) TableName() string {
	return "payment"
}