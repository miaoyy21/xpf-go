package ds

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type ApplicationCategory struct {
	BaseModel

	UserId int64 `gorm:"not null;type:bigint"`

	ProtoId int32  `gorm:"not null"`
	Name    string `gorm:"not null;size:512"`
}

func (m *ApplicationCategory) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_application_category_on_user_id", "user_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}


func (m *ApplicationCategory) GetName() (map[string]string, error) {
	name := make(map[string]string)

	if err := json.Unmarshal([]byte(m.Name), &name); err != nil {
		return nil, err
	}

	return name, nil
}

func (m *ApplicationCategory) SetName(name map[string]string) error {
	bs, err := json.Marshal(name)
	if err != nil {
		return err
	}

	m.Name = string(bs)
	return nil
}

func CreateApplicationCategory(tx *gorm.DB, userId int64, protoId int32, name map[string]string) (*ApplicationCategory, error) {
	a := &ApplicationCategory{
		UserId: userId,

		ProtoId: protoId,
	}
	a.ID = GenerateId()

	if err := a.SetName(name); err != nil {
		return nil, err
	}

	if err := tx.Create(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func FindApplicationCategoryById(tx *gorm.DB, id int64) (*ApplicationCategory, error) {
	var a ApplicationCategory

	db := tx.First(&a, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func FindApplicationCategoriesByUserId(tx *gorm.DB, userId int64) ([]*ApplicationCategory, error) {
	var cs []*ApplicationCategory

	db := tx.Where(&ApplicationCategory{UserId: userId}).Find(&cs)
	if err := db.Error; err != nil {
		return nil, err
	}

	return cs, nil
}

func FindApplicationCategoryByUserIdProtoId(tx *gorm.DB, userId int64, protoId int32) (*ApplicationCategory, error) {
	var a ApplicationCategory

	db := tx.Where(&ApplicationCategory{UserId: userId, ProtoId: protoId}).First(&a)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func DeleteApplicationCategoryById(tx *gorm.DB, id int64) error {
	db := tx.Delete(&ApplicationCategory{}, id)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
