package md

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type UploadFile struct {
}

func (m UploadFile) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.UploadFileResp{
		Rt: SuccessCode.Int32(),
	}

	fileType := pbMsg.(*pb.UploadFile).GetType()
	accountId := pbMsg.(*pb.UploadFile).GetAccountId()
	name := pbMsg.(*pb.UploadFile).GetName()
	mime := pbMsg.(*pb.UploadFile).GetMime()
	ext := pbMsg.(*pb.UploadFile).GetExt()
	bytes := pbMsg.(*pb.UploadFile).GetBytes()
	thumbnail := pbMsg.(*pb.UploadFile).GetThumbnail()
	size := pbMsg.(*pb.UploadFile).GetSize()

	// Exceed Size
	if len(bytes) > user.MaxFileSizeM<<20 {
		respPbMsg.Rt = FileTooLargeCode.Int32()
		return respPbMsg, fmt.Errorf("file Size too large %d (MaxSize is %d)", len(bytes), 100<<20)
	}

	if fileType == pb.FileType_FileImage && accountId <= 0 {
		respPbMsg.Rt = ArgumentsErrorCode.Int32()
		return respPbMsg, errors.New("required field 'AccountId' for File of Image")
	}

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	id := ds.GenerateId()

	// Write File to Store
	if err := TargetDirectoryFile.writeFile(user.ID, id, bytes); err != nil {
		respPbMsg.Rt = WriteFileErrorCode.Int32()
		return respPbMsg, err
	}

	// Write File to Thumbnail
	if len(thumbnail) > 0 {
		if err := TargetDirectoryThumbnail.writeFile(user.ID, id, thumbnail); err != nil {
			respPbMsg.Rt = WriteFileErrorCode.Int32()
			return respPbMsg, err
		}
	}

	// Create File
	f, err := ds.CreateFile(tx, user.ID, id, fileType, accountId, name, mime, ext, len(thumbnail) > 0, size)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	respPbMsg.Id = proto.Int64(f.ID)
	respPbMsg.UploadAt = proto.Int32(int32(f.CreatedAt.Unix()))

	return respPbMsg, nil
}
