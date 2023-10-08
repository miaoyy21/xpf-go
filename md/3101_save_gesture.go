package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type SaveGesture struct {
}

func (m SaveGesture) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.SaveGestureResp{
		Rt: SuccessCode.Int32(),
	}

	isGesture := pbMsg.(*pb.SaveGesture).GetIsGesture()
	gesture := pbMsg.(*pb.SaveGesture).GetGesture()
	validityGesture := pbMsg.(*pb.SaveGesture).GetValidityGesture()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Save User
	user.IsGesture = isGesture
	user.Gesture = gesture
	user.ValidityGesture = validityGesture
	if err := tx.Save(user).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	respPbMsg.IsGesture = proto.Bool(isGesture)
	respPbMsg.Gesture = gesture
	respPbMsg.ValidityGesture = proto.Int32(validityGesture)
	return respPbMsg, nil
}
