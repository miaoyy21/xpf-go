package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type GetApplicationAccount struct {
}

func (m GetApplicationAccount) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.GetApplicationAccountResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.GetApplicationAccount).GetId()

	// Application Account
	account, err := ds.FindApplicationAccountById(tx, id)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Application Name
	appName, err := account.GetAppName()
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Account Fields
	fields, err := ds.FindApplicationAccountFieldsByAccountId(tx, account.ID)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Protobuf Account
	pbAccount := &pb.ApplicationAccount{
		Id:         proto.Int64(account.ID),
		AccountId:  proto.String(account.AccountId),
		CategoryId: proto.Int64(account.CategoryId),
		AppId:      proto.String(account.AppId),
		AppName:    appName,
		CreateAt:   proto.Int32(int32(account.CreatedAt.Unix())),

		Fields: make([]*pb.AccountField, 0),
	}

	// Account Fields
	for _, field := range fields {
		fieldName, err := field.GetName()
		if err != nil {
			respPbMsg.Rt = SqlSelectFailureCode.Int32()
			return respPbMsg, err
		}

		// Protobuf Account Field
		pbField := &pb.AccountField{
			Index:     proto.Int32(field.Index),
			Name:      fieldName,
			IsPrimary: proto.Bool(field.IsPrimary),
			Type:      field.Type.Enum(),
			Bytes:     field.Bytes,
			MinLines:  proto.Int32(field.MinLines),
			MaxLines:  proto.Int32(field.MaxLines),
		}

		pbAccount.Fields = append(pbAccount.Fields, pbField)
	}

	respPbMsg.Account = pbAccount

	return respPbMsg, nil
}
