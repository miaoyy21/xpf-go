package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"psw/pb"
)

type CreateApplicationCategory struct {
}

func (m CreateApplicationCategory) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.CreateApplicationCategoryResp{
		Rt: SuccessCode.Int32(),
	}

	categoryName := pbMsg.(*pb.CreateApplicationCategory).GetCategoryName()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Category
	category, err := ds.CreateApplicationCategory(tx, user.ID, -1, categoryName)
	if err != nil {
		respPbMsg.Rt = SqlCreateFailureCode.Int32()
		return respPbMsg, err
	}

	// Category Name
	name, err := category.GetName()
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Protobuf Category
	pbCategory := &pb.ApplicationCategory{
		Id: proto.Int64(category.ID),

		ProtoId: proto.Int32(category.ProtoId),
		Name:    name,
	}

	// Categories Seq
	seqs, err := user.GetApplicationCategoriesSeq()
	if err != nil {
		respPbMsg.Rt = UnmarshalJsonFailureCode.Int32()
		return respPbMsg, err
	}

	// Append Categories Seq
	seqs = append(seqs, category.ID)
	if err := user.SetApplicationCategoriesSeq(seqs); err != nil {
		respPbMsg.Rt = MarshalJsonFailureCode.Int32()
		return respPbMsg, err
	}

	// Save Categories Seq
	if err := tx.Save(user).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	respPbMsg.Category = pbCategory

	return respPbMsg, nil
}
