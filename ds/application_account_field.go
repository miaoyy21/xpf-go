package ds

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"xpf/pb"
)

type ApplicationAccountField struct {
	gorm.Model

	AccountId int64 `gorm:"not null;type:bigint"`

	Index     int32        `gorm:"not null"`
	Name      string       `gorm:"not null;size:512"` // Map[string]string
	IsPrimary bool         `gorm:"not null"`
	Type      pb.FieldType `gorm:"not null"`
	Bytes     []byte       `gorm:"not null;size:4096"`
	MinLines  int32        `gorm:"not null"`
	MaxLines  int32        `gorm:"not null"`
}

func (m *ApplicationAccountField) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_application_account_field_on_account_id", "account_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (m *ApplicationAccountField) GetName() (map[string]string, error) {
	name := make(map[string]string)

	if err := json.Unmarshal([]byte(m.Name), &name); err != nil {
		return nil, err
	}

	return name, nil
}

func (m *ApplicationAccountField) SetName(name map[string]string) error {
	bs, err := json.Marshal(name)
	if err != nil {
		return err
	}

	m.Name = string(bs)
	return nil
}

func CreateApplicationAccountField(tx *gorm.DB, accountId int64, index int32, name map[string]string, isParmary bool, fieldType pb.FieldType, bytes []byte, minLines int32, maxLines int32) (*ApplicationAccountField, error) {
	a := &ApplicationAccountField{
		AccountId: accountId,

		Index:     index,
		IsPrimary: isParmary,
		Type:      fieldType,
		Bytes:     bytes,
		MinLines:  minLines,
		MaxLines:  maxLines,
	}

	// Set Field Name
	if err := a.SetName(name); err != nil {
		return nil, err
	}

	// Create
	if err := tx.Create(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func FindApplicationAccountFieldsByAccountId(tx *gorm.DB, accountId int64) ([]*ApplicationAccountField, error) {
	var fs []*ApplicationAccountField

	db := tx.Where(&ApplicationAccountField{AccountId: accountId}).Find(&fs)
	if err := db.Error; err != nil {
		return nil, err
	}

	return fs, nil
}

func XDeleteApplicationAccountFieldByAccountId(tx *gorm.DB, accountId int64) error {
	db := tx.Unscoped().Where(&ApplicationAccountField{AccountId: accountId}).Delete(&ApplicationAccountField{})
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
