// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package notification

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Notification) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Notification[number], err)
}

func (x *Notification) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Notification) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Notification) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = NotificationType(v)
	return offset, nil
}

func (x *Notification) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Notification) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Timestamp, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Notification) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Read, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *SendNotificationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendNotificationRequest[number], err)
}

func (x *SendNotificationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendNotificationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = NotificationType(v)
	return offset, nil
}

func (x *SendNotificationRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendNotificationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendNotificationResponse[number], err)
}

func (x *SendNotificationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *SendNotificationResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Notification
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Notification = &v
	return offset, nil
}

func (x *GetNotificationsRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationsRequest[number], err)
}

func (x *GetNotificationsRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetNotificationsResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationsResponse[number], err)
}

func (x *GetNotificationsResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *GetNotificationsResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Notification
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Notifications = append(x.Notifications, &v)
	return offset, nil
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

func (x *Notification) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Notification) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Notification) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *Notification) fastWriteField3(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetType()))
	return offset
}

func (x *Notification) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *Notification) fastWriteField5(buf []byte) (offset int) {
	if x.Timestamp == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetTimestamp())
	return offset
}

func (x *Notification) fastWriteField6(buf []byte) (offset int) {
	if !x.Read {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 6, x.GetRead())
	return offset
}

func (x *SendNotificationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SendNotificationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *SendNotificationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetType()))
	return offset
}

func (x *SendNotificationRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetContent())
	return offset
}

func (x *SendNotificationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SendNotificationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCommonResponse())
	return offset
}

func (x *SendNotificationResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Notification == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetNotification())
	return offset
}

func (x *GetNotificationsRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetNotificationsRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetNotificationsResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetNotificationsResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCommonResponse())
	return offset
}

func (x *GetNotificationsResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Notifications == nil {
		return offset
	}
	for i := range x.GetNotifications() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetNotifications()[i])
	}
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

func (x *Notification) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Notification) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *Notification) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserId())
	return n
}

func (x *Notification) sizeField3() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetType()))
	return n
}

func (x *Notification) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *Notification) sizeField5() (n int) {
	if x.Timestamp == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetTimestamp())
	return n
}

func (x *Notification) sizeField6() (n int) {
	if !x.Read {
		return n
	}
	n += fastpb.SizeBool(6, x.GetRead())
	return n
}

func (x *SendNotificationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SendNotificationRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *SendNotificationRequest) sizeField2() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetType()))
	return n
}

func (x *SendNotificationRequest) sizeField3() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetContent())
	return n
}

func (x *SendNotificationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SendNotificationResponse) sizeField1() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCommonResponse())
	return n
}

func (x *SendNotificationResponse) sizeField2() (n int) {
	if x.Notification == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetNotification())
	return n
}

func (x *GetNotificationsRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetNotificationsRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *GetNotificationsResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetNotificationsResponse) sizeField1() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCommonResponse())
	return n
}

func (x *GetNotificationsResponse) sizeField2() (n int) {
	if x.Notifications == nil {
		return n
	}
	for i := range x.GetNotifications() {
		n += fastpb.SizeMessage(2, x.GetNotifications()[i])
	}
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

var fieldIDToName_Notification = map[int32]string{
	1: "Id",
	2: "UserId",
	3: "Type",
	4: "Content",
	5: "Timestamp",
	6: "Read",
}

var fieldIDToName_SendNotificationRequest = map[int32]string{
	1: "UserId",
	2: "Type",
	3: "Content",
}

var fieldIDToName_SendNotificationResponse = map[int32]string{
	1: "CommonResponse",
	2: "Notification",
}

var fieldIDToName_GetNotificationsRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetNotificationsResponse = map[int32]string{
	1: "CommonResponse",
	2: "Notifications",
}

var fieldIDToName_CommonResponse = map[int32]string{
	1: "Code",
	2: "Message",
}
