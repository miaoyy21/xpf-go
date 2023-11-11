package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type GetUserData struct {
}

func (m GetUserData) DealWith(tx *gorm.DB, user *ds.User, _ proto.Message) (proto.Message, error) {
	respPbMsg := &pb.GetUserDataResp{
		Rt: SuccessCode.Int32(),
	}

	// User Data
	userData, err := NewUserData(user).Message(tx)
	if err != nil {
		respPbMsg.Rt = InternalErrorCode.Int32()
		return respPbMsg, err
	}
	respPbMsg.UserData = userData

	return respPbMsg, nil
}
