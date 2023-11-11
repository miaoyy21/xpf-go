package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/cache"
	"xpf/ds"
	"xpf/pb"
)

type LoadAssets struct {
}

func (m LoadAssets) DealWith(tx *gorm.DB, user *ds.User, _ proto.Message) (proto.Message, error) {
	respPbMsg := &pb.LoadAssetsResp{
		Rt:           SuccessCode.Int32(),
		Applications: make([]*pb.Application, 0),
	}

	respPbMsg.Applications = cache.Cache.Applications
	return respPbMsg, nil
}
