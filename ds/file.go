package ds

import (
	"github.com/jinzhu/gorm"
	"psw/pb"
)

var MaxFileSizeM = 128

type File struct {
	BaseModel

	UserId    int64       `gorm:"not null;type:bigint"`
	Type      pb.FileType `gorm:"not null"`
	AccountId int64       `gorm:"not null;type:bigint"`

	Name         string `gorm:"not null;size:2048"`
	Mime         string `gorm:"not null;size:256"`
	Ext          string `gorm:"not null;size:256"`
	HasThumbnail bool   `gorm:"not null"`
	Size         int32  `gorm:"not null"`

	IsShared bool `gorm:"not null"`
}

func (m *File) AfterCreateTable(db *gorm.DB) error {
	tx1 := db.Model(m).AddIndex("idx_file_on_user_id", "user_id")
	if err := tx1.Error; err != nil {
		return err
	}

	tx2 := db.Model(m).AddIndex("idx_file_on_account_id", "account_id")
	if err := tx2.Error; err != nil {
		return err
	}

	return nil
}

func CreateFile(tx *gorm.DB, userId int64, id int64, fileType pb.FileType, accountId int64, name string, mime string, ext string, hasThumbnail bool, size int32) (*File, error) {
	fls := &File{
		UserId:    userId,
		Type:      fileType,
		AccountId: accountId,

		Name:         name,
		Mime:         mime,
		Ext:          ext,
		HasThumbnail: hasThumbnail,
		Size:         size,

		IsShared: false,
	}
	fls.ID = id

	if err := tx.Create(fls).Error; err != nil {
		return nil, err
	}

	return fls, nil
}

func FindFileById(tx *gorm.DB, id int64) (*File, error) {
	var f File

	db := tx.First(&f, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func FindFilesByAccountId(tx *gorm.DB, accountId int64) ([]*File, error) {
	var fs []*File

	db := tx.Where(&File{AccountId: accountId}).Find(&fs)
	if err := db.Error; err != nil {
		return nil, err
	}

	return fs, nil
}

func FindFilesByUserId(tx *gorm.DB, userId int64) ([]*File, error) {
	var fs []*File

	db := tx.Where(&File{UserId: userId}).Find(&fs)
	if err := db.Error; err != nil {
		return nil, err
	}

	return fs, nil
}

func FindDeletedFilesByUserId(tx *gorm.DB, userId int64) ([]*File, error) {
	var as []*File

	db := tx.Unscoped().Where(&File{UserId: userId}).Where("deleted_at IS NOT NULL").Find(&as)
	if err := db.Error; err != nil {
		return nil, err
	}

	return as, nil
}

func FindDeletedFileById(tx *gorm.DB, id int64) (*File, error) {
	var f File

	db := tx.Unscoped().First(&f, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func DeleteFileById(tx *gorm.DB, id int64) error {
	db := tx.Delete(&File{}, id)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}
