package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type DeleteApplicationAccount struct {
}

func (m DeleteApplicationAccount) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.DeleteApplicationAccountResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.DeleteApplicationAccount).GetId()

	// Delete Account
	if err := ds.DeleteApplicationAccount(tx, id); err != nil {
		respPbMsg.Rt = SqlDeleteFailureCode.Int32()
		return respPbMsg, err
	}

	respPbMsg.Id = proto.Int64(id)
	return respPbMsg, nil
}
