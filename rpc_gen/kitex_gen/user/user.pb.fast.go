// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Phone, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterRequest[number], err)
}

func (x *RegisterRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PasswordConfirm, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Phone, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterResponse[number], err)
}

func (x *RegisterResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *RegisterResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *LoginRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginRequest[number], err)
}

func (x *LoginRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResponse[number], err)
}

func (x *LoginResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
	return offset, nil
}

func (x *GetUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserRequest[number], err)
}

func (x *GetUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserResponse[number], err)
}

func (x *GetUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *GetUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CommonResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommonResponse = &v
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

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPassword())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetEmail())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if x.Phone == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPhone())
	return offset
}

func (x *RegisterRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *RegisterRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *RegisterRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *RegisterRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PasswordConfirm == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPasswordConfirm())
	return offset
}

func (x *RegisterRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetEmail())
	return offset
}

func (x *RegisterRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Phone == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPhone())
	return offset
}

func (x *RegisterResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RegisterResponse) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *RegisterResponse) fastWriteField2(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetCommonResponse())
	return offset
}

func (x *LoginRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *LoginRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *LoginRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *LoginResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *LoginResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *LoginResponse) fastWriteField2(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetCommonResponse())
	return offset
}

func (x *GetUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *GetUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.CommonResponse == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetCommonResponse())
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

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPassword())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetEmail())
	return n
}

func (x *User) sizeField5() (n int) {
	if x.Phone == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPhone())
	return n
}

func (x *RegisterRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *RegisterRequest) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *RegisterRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *RegisterRequest) sizeField3() (n int) {
	if x.PasswordConfirm == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPasswordConfirm())
	return n
}

func (x *RegisterRequest) sizeField4() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetEmail())
	return n
}

func (x *RegisterRequest) sizeField5() (n int) {
	if x.Phone == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPhone())
	return n
}

func (x *RegisterResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RegisterResponse) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *RegisterResponse) sizeField2() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetCommonResponse())
	return n
}

func (x *LoginRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *LoginRequest) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *LoginRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *LoginResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *LoginResponse) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *LoginResponse) sizeField2() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetCommonResponse())
	return n
}

func (x *GetUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserRequest) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *GetUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetUserResponse) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *GetUserResponse) sizeField2() (n int) {
	if x.CommonResponse == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetCommonResponse())
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

var fieldIDToName_User = map[int32]string{
	1: "Id",
	2: "Username",
	3: "Password",
	4: "Email",
	5: "Phone",
}

var fieldIDToName_RegisterRequest = map[int32]string{
	1: "Username",
	2: "Password",
	3: "PasswordConfirm",
	4: "Email",
	5: "Phone",
}

var fieldIDToName_RegisterResponse = map[int32]string{
	1: "User",
	2: "CommonResponse",
}

var fieldIDToName_LoginRequest = map[int32]string{
	1: "Username",
	2: "Password",
}

var fieldIDToName_LoginResponse = map[int32]string{
	1: "Token",
	2: "CommonResponse",
}

var fieldIDToName_GetUserRequest = map[int32]string{
	1: "Id",
}

var fieldIDToName_GetUserResponse = map[int32]string{
	1: "User",
	2: "CommonResponse",
}

var fieldIDToName_CommonResponse = map[int32]string{
	1: "Code",
	2: "Message",
}
