package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
	"time"
)

type RestoreTrash struct {
}

func (m RestoreTrash) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.RestoreTrashResp{
		Rt: SuccessCode.Int32(),
	}

	accountId := pbMsg.(*pb.RestoreTrash).GetAccountId()
	fileId := pbMsg.(*pb.RestoreTrash).GetFileId()

	var cost int64

	// Account
	if accountId > 0 {
		a, err := ds.FindDeletedApplicationAccountsById(tx, accountId)
		if err != nil {
			respPbMsg.Rt = SqlSelectFailureCode.Int32()
			return respPbMsg, err
		}

		cost = (int64(time.Now().Sub(*a.DeletedAt).Seconds()) / (24 * 3600)) * (ds.CostAccountScore + int64(a.Size)*ds.CostFile10MScore/(10<<20))

		a.DeletedAt = nil
		if err := tx.Unscoped().Save(a).Error; err != nil {
			respPbMsg.Rt = SqlSaveFailureCode.Int32()
			return respPbMsg, err
		}

		appName, err := a.GetAppName()
		if err != nil {
			respPbMsg.Rt = UnmarshalJsonFailureCode.Int32()
			return respPbMsg, err
		}

		respPbMsg.Snapshot = &pb.AccountSnapshot{
			Id:         proto.Int64(a.ID),
			CategoryId: proto.Int64(a.CategoryId),
			AppId:      proto.String(a.AppId),
			AppName:    appName,
			AccountId:  proto.String(a.AccountId),
			CreateAt:   proto.Int32(int32(a.CreatedAt.Unix())),
			Size:       proto.Int32(a.Size),
		}
	}

	// File
	if fileId > 0 {
		f, err := ds.FindDeletedFileById(tx, fileId)
		if err != nil {
			respPbMsg.Rt = SqlSelectFailureCode.Int32()
			return respPbMsg, err
		}

		cost = (int64(time.Now().Sub(*f.DeletedAt).Seconds()) / (24 * 3600)) * (int64(f.Size) * ds.CostFile10MScore / (10 << 20))

		f.DeletedAt = nil
		if err := tx.Unscoped().Save(f).Error; err != nil {
			respPbMsg.Rt = SqlSaveFailureCode.Int32()
			return respPbMsg, err
		}

		respPbMsg.File = &pb.File{
			Id:           proto.Int64(f.ID),
			Name:         proto.String(f.Name),
			Mime:         proto.String(f.Mime),
			Ext:          proto.String(f.Ext),
			HasThumbnail: proto.Bool(f.HasThumbnail),
			Size:         proto.Int32(f.Size),
			IsShared:     proto.Bool(f.IsShared),
			UploadAt:     proto.Int32(int32(f.CreatedAt.Unix())),
		}
	}

	if cost != 0 {
		// User's Score Cost Record
		score := user.Score + cost
		err := ds.CreateScore(tx, user.ID, ds.RestoreScoreAction, cost, score)
		if err != nil {
			respPbMsg.Rt = SqlCreateFailureCode.Int32()
			return respPbMsg, err
		}

		// Save User's Score
		user.Score = score
		if err := tx.Save(user).Error; err != nil {
			respPbMsg.Rt = SqlSaveFailureCode.Int32()
			return respPbMsg, err
		}
	}

	respPbMsg.Score = proto.Int64(user.Score)
	return respPbMsg, nil
}
