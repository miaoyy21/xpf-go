package md

import "google.golang.org/protobuf/proto"

var (
	SuccessCode        Code = 1
	ArgumentsErrorCode Code = 2 // 参数错误
	ReadFileErrorCode  Code = 3 // 文件读取失败
	WriteFileErrorCode Code = 4 // 文件写入失败
	ScoreNotEnough     Code = 5 // 积分不足
	FileTooLargeCode   Code = 6 // 文件体积超过最大限制
	InternalErrorCode  Code = 9 // 内部错误

	SqlSelectFailureCode Code = 11 // 执行检索失败
	SqlCreateFailureCode Code = 12 // 执行新增失败
	SqlSaveFailureCode   Code = 13 // 执行更新失败
	SqlDeleteFailureCode Code = 14 // 执行删除失败

	MarshalJsonFailureCode   Code = 21 // JSON序列化失败
	UnmarshalJsonFailureCode Code = 22 // JSON反序列化失败

	PurchaseReceiptFailureCode Code = 91 // 内购验证请求失败
)

type Code int32

func (code Code) Int32() *int32 {
	return proto.Int32(int32(code))
}
