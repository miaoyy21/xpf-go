package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/cache"
	"xpf/ds"
	"xpf/pb"
)

type CreateApplicationAccount struct {
}

func (m CreateApplicationAccount) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.CreateApplicationAccountResp{
		Rt: SuccessCode.Int32(),
	}

	appId := pbMsg.(*pb.CreateApplicationAccount).GetAppId()
	appName := pbMsg.(*pb.CreateApplicationAccount).GetAppName()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Create Account
	account, err := ds.CreateApplicationAccount(tx, user.ID, appId, appName)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// Category ProtoId
	categoryProtoId, ok := cache.Cache.MapApplicationCategory[appId]
	if !ok {
		categoryProtoId = 0
	}

	// Category
	category, err := ds.FindApplicationCategoryByUserIdProtoId(tx, user.ID, categoryProtoId)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	pbAccount := &pb.ApplicationAccount{
		Id:         proto.Int64(account.ID),
		AccountId:  proto.String(account.AccountId),
		CategoryId: proto.Int64(category.ID),
		AppId:      proto.String(account.AppId),
		AppName:    appName,
		CreateAt:   proto.Int32(int32(account.CreatedAt.Unix())),
		Fields:     make([]*pb.AccountField, 0),
	}

	// Category Fields
	fields := make([]*cache.CategoryField, 0)
	mapCategory, ok := cache.Cache.MapCategories[category.ProtoId]
	if !ok {
		fields = cache.Cache.DefaultCategoryFields
	} else {
		fields = mapCategory.Fields
	}

	// Protobuf Fields
	for i, field := range fields {
		pbField := &pb.AccountField{
			Index:     proto.Int32(int32(i + 1)),
			Name:      field.Name,
			IsPrimary: proto.Bool(field.IsPrimary),
			Type:      field.Type.Enum(),
			Bytes:     []byte{},
			MinLines:  proto.Int32(field.MinLines),
			MaxLines:  proto.Int32(field.MaxLines),
		}

		pbAccount.Fields = append(pbAccount.Fields, pbField)
	}

	respPbMsg.Account = pbAccount

	return respPbMsg, nil
}
