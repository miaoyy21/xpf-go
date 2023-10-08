package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type SaveApplicationAccount struct {
}

func (m SaveApplicationAccount) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.SaveApplicationAccountResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.SaveApplicationAccount).GetId()
	accountId := pbMsg.(*pb.SaveApplicationAccount).GetAccountId()
	categoryId := pbMsg.(*pb.SaveApplicationAccount).GetCategoryId()
	appId := pbMsg.(*pb.SaveApplicationAccount).GetAppId()
	appName := pbMsg.(*pb.SaveApplicationAccount).GetAppName()
	fields := pbMsg.(*pb.SaveApplicationAccount).GetFields()
	fileIds := pbMsg.(*pb.SaveApplicationAccount).GetFileIds()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Account ID
	if len(accountId) == 0 || categoryId == 0 || len(appId) == 0 || len(appName) == 0 {
		respPbMsg.Rt = ArgumentsErrorCode.Int32()
		return respPbMsg, errors.New("arguments has empty element")
	}

	// Account
	account, err := ds.FindApplicationAccountById(tx, id)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Delete Account Fields
	if err := ds.XDeleteApplicationAccountFieldByAccountId(tx, id); err != nil {
		respPbMsg.Rt = SqlDeleteFailureCode.Int32()
		return respPbMsg, err
	}

	// Save Account Fields
	for i, field := range fields {
		_, err := ds.CreateApplicationAccountField(tx, id, int32(i+1), field.GetName(), field.GetIsPrimary(), field.GetType(), field.GetBytes(), field.GetMinLines(), field.GetMaxLines())
		if err != nil {
			respPbMsg.Rt = SqlCreateFailureCode.Int32()
			return respPbMsg, err
		}
	}

	// Get Store Files by ID
	files, err := ds.FindFilesByAccountId(tx, id)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Map Files
	mapFiles := make(map[int64]*ds.File)
	for _, file := range files {
		mapFiles[file.ID] = file
	}

	// Set IsValid
	var size int32 = 0
	for _, id := range fileIds {
		file, ok := mapFiles[id]
		if ok {
			size += file.Size

			delete(mapFiles, id)
		}
	}

	// Delete File
	for _, file := range mapFiles {
		if err := tx.Delete(file).Error; err != nil {
			respPbMsg.Rt = SqlDeleteFailureCode.Int32()
			return respPbMsg, err
		}
	}

	// Save Account
	account.AccountId = accountId
	account.CategoryId = categoryId
	account.AppId = appId
	account.Size = size

	// Set Account Application Name
	if err := account.SetAppName(appName); err != nil {
		respPbMsg.Rt = MarshalJsonFailureCode.Int32()
		return respPbMsg, err
	}

	// Save Account
	if err := tx.Save(account).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	// Protobuf Snapshot
	pbSnapshot := &pb.AccountSnapshot{
		Id:         proto.Int64(account.ID),
		CategoryId: proto.Int64(account.CategoryId),
		AppId:      proto.String(account.AppId),
		AppName:    appName,
		AccountId:  proto.String(account.AccountId),
		CreateAt:   proto.Int32(int32(account.CreatedAt.Unix())),
		Size:       proto.Int32(account.Size),
	}
	respPbMsg.Snapshot = pbSnapshot

	return respPbMsg, nil
}
