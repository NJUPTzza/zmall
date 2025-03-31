package model

import (
	"gorm.io/gorm"
)

type NotificationType int

const (
	NotificationTypePaymentSuccess NotificationType = 1 // 支付成功通知
	NotificationTypeOrderCreated   NotificationType = 2   // 订单创建通知
	NotificationTypeOrderShipped   NotificationType = 3  // 订单发货通知
	NotificationTypeOrderDelivered NotificationType = 4 // 订单送达通知
	NotificationTypeRefundSuccess  NotificationType = 5  // 退款成功通知
)

type Notification struct {
    gorm.Model
    UserID    int64  `gorm:"type:int;not null"`
    Type      NotificationType `gorm:"type:tinyint;not null"`
    Content   string `gorm:"type:text;not null"`
	Timestamp int64           `gorm:"not null"`
    Read      bool   `gorm:"not null;default:false"`
}

func (Notification) TableName() string {
    return "notification"
}