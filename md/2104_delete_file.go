package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type DeleteFile struct {
}

func (m DeleteFile) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.DeleteFileResp{
		Rt: SuccessCode.Int32(),
	}

	id := pbMsg.(*pb.DeleteFile).GetId()

	// Delete File
	err := ds.DeleteFileById(tx, id)
	if err != nil {
		respPbMsg.Rt = SqlDeleteFailureCode.Int32()
		return respPbMsg, err
	}

	// TODO
	//
	//// Delete File
	//if err := ds.DeleteFileByStoreId(tx, storeId); err != nil {
	//	respPbMsg.Rt = SqlDeleteFailureCode.Int32()
	//	return respPbMsg, err
	//}
	//
	//p1 := filepath.Join("store", "file", strconv.FormatInt(user.ID, 32), storeId)
	//
	//newStoreId := fmt.Sprintf("%s_%s", strconv.FormatInt(user.ID, 32), storeId)
	//p2 := filepath.Join("store", "file", "_", newStoreId)
	//
	//if err := os.Rename(p1, p2); err != nil {
	//	logrus.Warnf("2104(DeleteFile) Rename() Failure %s", err.Error())
	//}

	respPbMsg.Id = proto.Int64(id)
	return respPbMsg, nil
}
