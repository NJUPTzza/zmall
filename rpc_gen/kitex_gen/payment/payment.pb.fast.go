// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package payment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Payment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Payment[number], err)
}

func (x *Payment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Payment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Payment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *Payment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Status = PaymentStatus(v)
	return offset, nil
}

func (x *ProcessPaymentRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProcessPaymentRequest[number], err)
}

func (x *ProcessPaymentRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProcessPaymentRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *ProcessPaymentRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PayChannel, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProcessPaymentResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProcessPaymentResponse[number], err)
}

func (x *ProcessPaymentResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *ProcessPaymentResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Payment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Payment = &v
	return offset, nil
}

func (x *ProcessPaymentResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPaymentStatusRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPaymentStatusRequest[number], err)
}

func (x *GetPaymentStatusRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PaymentId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetPaymentStatusResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPaymentStatusResponse[number], err)
}

func (x *GetPaymentStatusResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *GetPaymentStatusResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Payment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Payment = &v
	return offset, nil
}

func (x *UpdatePaymentStatusRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdatePaymentStatusRequest[number], err)
}

func (x *UpdatePaymentStatusRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PaymentId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdatePaymentStatusRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Event = PaymentEvent(v)
	return offset, nil
}

func (x *UpdatePaymentStatusResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdatePaymentStatusResponse[number], err)
}

func (x *UpdatePaymentStatusResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *UpdatePaymentStatusResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Payment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Payment = &v
	return offset, nil
}

func (x *PaymentSuccessMessage) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PaymentSuccessMessage[number], err)
}

func (x *PaymentSuccessMessage) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PaymentId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PaymentSuccessMessage) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PaymentSuccessMessage) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *PaymentSuccessMessage) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PaymentStatus, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CommonResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CommonResponse[number], err)
}

func (x *CommonResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CommonResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Payment) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Payment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Payment) fastWriteField2(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetOrderId())
	return offset
}

func (x *Payment) fastWriteField3(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 3, x.GetAmount())
	return offset
}

func (x *Payment) fastWriteField4(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, int32(x.GetStatus()))
	return offset
}

func (x *ProcessPaymentRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ProcessPaymentRequest) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *ProcessPaymentRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 2, x.GetAmount())
	return offset
}

func (x *ProcessPaymentRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PayChannel == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPayChannel())
	return offset
}

func (x *ProcessPaymentResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ProcessPaymentResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCommonResponse())
	return offset
}

func (x *ProcessPaymentResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Payment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPayment())
	return offset
}

func (x *ProcessPaymentResponse) fastWriteField3(buf []byte) (offset int) {
	if x.PayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPayUrl())
	return offset
}

func (x *GetPaymentStatusRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPaymentStatusRequest) fastWriteField1(buf []byte) (offset int) {
	if x.PaymentId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetPaymentId())
	return offset
}

func (x *GetPaymentStatusResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetPaymentStatusResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCommonResponse())
	return offset
}

func (x *GetPaymentStatusResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Payment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPayment())
	return offset
}

func (x *UpdatePaymentStatusRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpdatePaymentStatusRequest) fastWriteField1(buf []byte) (offset int) {
	if x.PaymentId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetPaymentId())
	return offset
}

func (x *UpdatePaymentStatusRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Event == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetEvent()))
	return offset
}

func (x *UpdatePaymentStatusResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpdatePaymentStatusResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCommonResponse())
	return offset
}

func (x *UpdatePaymentStatusResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Payment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPayment())
	return offset
}

func (x *PaymentSuccessMessage) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PaymentSuccessMessage) fastWriteField1(buf []byte) (offset int) {
	if x.PaymentId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetPaymentId())
	return offset
}

func (x *PaymentSuccessMessage) fastWriteField2(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetOrderId())
	return offset
}

func (x *PaymentSuccessMessage) fastWriteField3(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 3, x.GetAmount())
	return offset
}

func (x *PaymentSuccessMessage) fastWriteField4(buf []byte) (offset int) {
	if x.PaymentStatus == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPaymentStatus())
	return offset
}

func (x *CommonResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CommonResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *CommonResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *Payment) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Payment) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *Payment) sizeField2() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetOrderId())
	return n
}

func (x *Payment) sizeField3() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeFloat(3, x.GetAmount())
	return n
}

func (x *Payment) sizeField4() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, int32(x.GetStatus()))
	return n
}

func (x *ProcessPaymentRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ProcessPaymentRequest) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetOrderId())
	return n
}

func (x *ProcessPaymentRequest) sizeField2() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeFloat(2, x.GetAmount())
	return n
}

func (x *ProcessPaymentRequest) sizeField3() (n int) {
	if x.PayChannel == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPayChannel())
	return n
}

func (x *ProcessPaymentResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ProcessPaymentResponse) sizeField1() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCommonResponse())
	return n
}

func (x *ProcessPaymentResponse) sizeField2() (n int) {
	if x.Payment == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPayment())
	return n
}

func (x *ProcessPaymentResponse) sizeField3() (n int) {
	if x.PayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPayUrl())
	return n
}

func (x *GetPaymentStatusRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPaymentStatusRequest) sizeField1() (n int) {
	if x.PaymentId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetPaymentId())
	return n
}

func (x *GetPaymentStatusResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetPaymentStatusResponse) sizeField1() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCommonResponse())
	return n
}

func (x *GetPaymentStatusResponse) sizeField2() (n int) {
	if x.Payment == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPayment())
	return n
}

func (x *UpdatePaymentStatusRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpdatePaymentStatusRequest) sizeField1() (n int) {
	if x.PaymentId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetPaymentId())
	return n
}

func (x *UpdatePaymentStatusRequest) sizeField2() (n int) {
	if x.Event == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetEvent()))
	return n
}

func (x *UpdatePaymentStatusResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpdatePaymentStatusResponse) sizeField1() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCommonResponse())
	return n
}

func (x *UpdatePaymentStatusResponse) sizeField2() (n int) {
	if x.Payment == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPayment())
	return n
}

func (x *PaymentSuccessMessage) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PaymentSuccessMessage) sizeField1() (n int) {
	if x.PaymentId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetPaymentId())
	return n
}

func (x *PaymentSuccessMessage) sizeField2() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetOrderId())
	return n
}

func (x *PaymentSuccessMessage) sizeField3() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeFloat(3, x.GetAmount())
	return n
}

func (x *PaymentSuccessMessage) sizeField4() (n int) {
	if x.PaymentStatus == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPaymentStatus())
	return n
}

func (x *CommonResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CommonResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *CommonResponse) sizeField2() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMessage())
	return n
}

var fieldIDToName_Payment = map[int32]string{
	1: "Id",
	2: "OrderId",
	3: "Amount",
	4: "Status",
}

var fieldIDToName_ProcessPaymentRequest = map[int32]string{
	1: "OrderId",
	2: "Amount",
	3: "PayChannel",
}

var fieldIDToName_ProcessPaymentResponse = map[int32]string{
	1: "CommonResponse",
	2: "Payment",
	3: "PayUrl",
}

var fieldIDToName_GetPaymentStatusRequest = map[int32]string{
	1: "PaymentId",
}

var fieldIDToName_GetPaymentStatusResponse = map[int32]string{
	1: "CommonResponse",
	2: "Payment",
}

var fieldIDToName_UpdatePaymentStatusRequest = map[int32]string{
	1: "PaymentId",
	2: "Event",
}

var fieldIDToName_UpdatePaymentStatusResponse = map[int32]string{
	1: "CommonResponse",
	2: "Payment",
}

var fieldIDToName_PaymentSuccessMessage = map[int32]string{
	1: "PaymentId",
	2: "OrderId",
	3: "Amount",
	4: "PaymentStatus",
}

var fieldIDToName_CommonResponse = map[int32]string{
	1: "Code",
	2: "Message",
}
