package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type GetOperates struct {
}

func (m GetOperates) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.GetOperatesResp{
		Rt:       SuccessCode.Int32(),
		Operates: make([]*pb.Operate, 0),
	}

	limit := pbMsg.(*pb.GetOperates).GetLimit()
	offset := pbMsg.(*pb.GetOperates).GetOffset()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Get Operates
	ops, err := ds.FindOperatesByUserIdAt(tx, user.ID, limit, offset)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Protobuf Operate
	for _, op := range ops {
		msg, err := op.GetMsg()
		if err != nil {
			respPbMsg.Rt = SqlSelectFailureCode.Int32()
			return respPbMsg, err
		}

		operate := &pb.Operate{
			MsgNo: op.MsgNo.Enum(),
			Msg:   msg,
			Rt:    proto.Int32(op.Rt),
			Ip:    proto.String(op.IP),
			At:    proto.Int32(int32(op.CreatedAt.Unix())),
		}

		respPbMsg.Operates = append(respPbMsg.Operates, operate)
	}

	return respPbMsg, nil
}
