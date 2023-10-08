package ds

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type ApplicationAccount struct {
	BaseModel

	UserId int64 `gorm:"not null;type:bigint"`

	AccountId  string `gorm:"not null;size:128"`
	CategoryId int64  `gorm:"not null;type:bigint"`

	AppId   string `gorm:"not null;size:64"`
	AppName string `gorm:"not null;size:512"` // Map[string]string
	Size    int32  `gorm:"not null"`
}

func (m *ApplicationAccount) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_application_account_on_user_id", "user_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (m *ApplicationAccount) GetAppName() (map[string]string, error) {
	appName := make(map[string]string)

	if err := json.Unmarshal([]byte(m.AppName), &appName); err != nil {
		return nil, err
	}

	return appName, nil
}

func (m *ApplicationAccount) SetAppName(appName map[string]string) error {
	bs, err := json.Marshal(appName)
	if err != nil {
		return err
	}

	m.AppName = string(bs)
	return nil
}

func CreateApplicationAccount(tx *gorm.DB, userId int64, appId string, appName map[string]string) (*ApplicationAccount, error) {
	a := &ApplicationAccount{
		UserId:     userId,
		AccountId:  "",
		CategoryId: 0,
		AppId:      appId,
		Size:       0,
	}
	a.ID = GenerateId()

	// Set Application Id
	if len(appId) == 0 || strings.EqualFold(appId, "0") {
		a.AppId = strings.Join([]string{"psw", strconv.FormatInt(a.ID, 16)}, "-")
	}

	// Set Application Name
	if err := a.SetAppName(appName); err != nil {
		return nil, err
	}

	// Create
	db := tx.Create(a)
	if err := db.Error; err != nil {
		return nil, err
	}

	return a, nil
}

func FindApplicationAccountById(tx *gorm.DB, id int64) (*ApplicationAccount, error) {
	var ac ApplicationAccount

	db := tx.First(&ac, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &ac, nil
}

func FindApplicationAccountsByUserId(tx *gorm.DB, userId int64) ([]*ApplicationAccount, error) {
	var as []*ApplicationAccount

	db := tx.Where(&ApplicationAccount{UserId: userId}).Find(&as)
	if err := db.Error; err != nil {
		return nil, err
	}

	return as, nil
}

func FindApplicationAccountsByUserIdCategoryId(tx *gorm.DB, userId int64, categoryId int64) ([]*ApplicationAccount, error) {
	var as []*ApplicationAccount

	db := tx.Where(&ApplicationAccount{UserId: userId, CategoryId: categoryId}).Find(&as)
	if err := db.Error; err != nil {
		return nil, err
	}

	return as, nil
}

func FindDeletedApplicationAccountsByUserId(tx *gorm.DB, userId int64) ([]*ApplicationAccount, error) {
	var as []*ApplicationAccount

	db := tx.Unscoped().Where(&ApplicationAccount{UserId: userId}).Where("deleted_at IS NOT NULL").Find(&as)
	if err := db.Error; err != nil {
		return nil, err
	}

	return as, nil
}

func FindDeletedApplicationAccountsById(tx *gorm.DB, id int64) (*ApplicationAccount, error) {
	var a ApplicationAccount

	db := tx.Unscoped().First(&a, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func DeleteApplicationAccount(tx *gorm.DB, id int64) error {
	db := tx.Delete(&ApplicationAccount{}, id)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
