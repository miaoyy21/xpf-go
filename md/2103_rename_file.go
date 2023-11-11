package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type RenameFile struct {
}

func (m RenameFile) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.RenameFileResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.RenameFile).GetId()
	name := pbMsg.(*pb.RenameFile).GetName()

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

	// Save File
	f.Name = name
	if err := tx.Save(f).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	return respPbMsg, nil
}
