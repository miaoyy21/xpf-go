package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type ShareFile struct {
}

func (m ShareFile) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.ShareFileResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.ShareFile).GetId()
	isShared := pbMsg.(*pb.ShareFile).GetIsShared()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// File
	f, err := ds.FindFileById(tx, id)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Is Valid
	if isShared == f.IsShared {
		respPbMsg.Rt = ArgumentsErrorCode.Int32()
		return respPbMsg, err
	}

	// Share
	if isShared {
		if err := TargetDirectoryFile.copyFile(TargetDirectoryShare, user.ID, id); err != nil {
			respPbMsg.Rt = InternalErrorCode.Int32()
			return respPbMsg, err
		}
	} else {
		if err := TargetDirectoryShare.deleteFile(user.ID, id); err != nil {
			respPbMsg.Rt = InternalErrorCode.Int32()
			return respPbMsg, err
		}
	}

	// Save
	f.IsShared = isShared
	if err := tx.Save(f).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	return respPbMsg, nil
}
