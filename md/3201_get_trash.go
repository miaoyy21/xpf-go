package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type GetTrash struct {
}

func (m GetTrash) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.GetTrashResp{
		Rt:      SuccessCode.Int32(),
		Trashes: make([]*pb.Trash, 0),
	}

	// Get Deleted Accounts
	as, err := ds.FindDeletedApplicationAccountsByUserId(tx, user.ID)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	for _, a := range as {
		if len(a.AccountId) == 0 {
			continue
		}

		appName, err := a.GetAppName()
		if err != nil {
			respPbMsg.Rt = UnmarshalJsonFailureCode.Int32()
			return respPbMsg, err
		}

		pbSnapshot := &pb.AccountSnapshot{
			Id:         proto.Int64(a.ID),
			CategoryId: proto.Int64(a.CategoryId),
			AppId:      proto.String(a.AppId),
			AppName:    appName,
			AccountId:  proto.String(a.AccountId),
			CreateAt:   proto.Int32(int32(a.CreatedAt.Unix())),
			Size:       proto.Int32(a.Size),
		}

		pbTrash := &pb.Trash{Snapshot: pbSnapshot, DeleteAt: proto.Int32(int32(a.DeletedAt.Unix()))}
		respPbMsg.Trashes = append(respPbMsg.Trashes, pbTrash)
	}

	// Get Deleted Files
	fs, err := ds.FindDeletedFilesByUserId(tx, user.ID)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	for _, f := range fs {
		if f.Type != pb.FileType_FileCabinet {
			continue
		}

		pbFile := &pb.File{
			Id:           proto.Int64(f.ID),
			Name:         proto.String(f.Name),
			Mime:         proto.String(f.Mime),
			Ext:          proto.String(f.Ext),
			HasThumbnail: proto.Bool(f.HasThumbnail),
			Size:         proto.Int32(f.Size),
			IsShared:     proto.Bool(f.IsShared),
			UploadAt:     proto.Int32(int32(f.CreatedAt.Unix())),
		}

		pbTrash := &pb.Trash{File: pbFile, DeleteAt: proto.Int32(int32(f.DeletedAt.Unix()))}
		respPbMsg.Trashes = append(respPbMsg.Trashes, pbTrash)
	}

	return respPbMsg, nil
}
