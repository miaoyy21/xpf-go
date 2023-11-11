package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type SaveApplicationCategoriesSeq struct {
}

func (m SaveApplicationCategoriesSeq) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.SaveApplicationCategoriesSeqResp{
		Rt: SuccessCode.Int32(),
	}

	seqs := pbMsg.(*pb.SaveApplicationCategoriesSeq).GetSeqs()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	if err := user.SetApplicationCategoriesSeq(seqs); err != nil {
		respPbMsg.Rt = MarshalJsonFailureCode.Int32()
		return respPbMsg, err
	}

	// Save Categories Seq
	if err := tx.Save(user).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	return respPbMsg, nil
}
