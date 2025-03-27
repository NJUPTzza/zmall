package errno

import "fmt"

type BizError struct {
	Code int
	Message string 
	Data map[string]interface{}
}

func (e *BizError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s, data=%v", e.Code, e.Message, e.Data)
}

func NewStockNotEnough(stock int) *BizError {
	return &BizError{
		Code: 1001,
		Message: "库存不足",
		Data: map[string]interface{}{
			"stock": stock,
		},
	}
}