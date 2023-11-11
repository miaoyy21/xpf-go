package md

import (
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type SaveApplicationCategory struct {
}

func (m SaveApplicationCategory) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.SaveApplicationCategoryResp{
		Rt: SuccessCode.Int32(),
	}

	categoryId := pbMsg.(*pb.SaveApplicationCategory).GetCategoryId()
	categoryName := pbMsg.(*pb.SaveApplicationCategory).GetCategoryName()

	// Check Score
	if user.Score < 0 {
		respPbMsg.Rt = ScoreNotEnough.Int32()
		return respPbMsg, errors.New("user's Score is NOT Enough")
	}

	// Category
	category, err := ds.FindApplicationCategoryById(tx, categoryId)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Set Category Name
	if err := category.SetName(categoryName); err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Save
	if err := tx.Save(category).Error; err != nil {
		respPbMsg.Rt = SqlSaveFailureCode.Int32()
		return respPbMsg, err
	}

	return respPbMsg, nil
}
