package md

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/proto"
	"xpf/ds"
	"xpf/pb"
)

type DeleteApplicationCategory struct {
}

func (m DeleteApplicationCategory) DealWith(tx *gorm.DB, user *ds.User, pbMsg proto.Message) (proto.Message, error) {
	respPbMsg := &pb.DeleteApplicationCategoryResp{
		Rt: SuccessCode.Int32(),
	}

	categoryId := pbMsg.(*pb.DeleteApplicationCategory).GetCategoryId()

	// Delete Category
	if err := ds.DeleteApplicationCategoryById(tx, categoryId); err != nil {
		respPbMsg.Rt = SqlDeleteFailureCode.Int32()
		return respPbMsg, err
	}

	// Categories Seq
	seqs, err := user.GetApplicationCategoriesSeq()
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	index := -1
	for i, seq := range seqs {
		if seq == categoryId {
			index = i
			break
		}
	}
	if index >= 0 {
		// Remove Categories Seq
		seqs = append(seqs[:index], seqs[index+1:]...)
		if err := user.SetApplicationCategoriesSeq(seqs); err != nil {
			respPbMsg.Rt = MarshalJsonFailureCode.Int32()
			return respPbMsg, err
		}

		// Save Categories Seq
		if err := tx.Save(user).Error; err != nil {
			respPbMsg.Rt = SqlSaveFailureCode.Int32()
			return respPbMsg, err
		}
	}

	// Accounts
	accounts, err := ds.FindApplicationAccountsByUserIdCategoryId(tx, user.ID, categoryId)
	if err != nil {
		respPbMsg.Rt = SqlSelectFailureCode.Int32()
		return respPbMsg, err
	}

	// Update Accounts Category Id
	for _, account := range accounts {
		category, err := ds.FindApplicationCategoryByUserIdProtoId(tx, user.ID, 0)
		if err != nil {
			respPbMsg.Rt = SqlSelectFailureCode.Int32()
			return respPbMsg, err
		}

		account.CategoryId = category.ID
		if err := tx.Save(account).Error; err != nil {
			respPbMsg.Rt = SqlSaveFailureCode.Int32()
			return respPbMsg, err
		}
	}

	return respPbMsg, nil
}
