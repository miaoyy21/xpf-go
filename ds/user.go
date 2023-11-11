package ds

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
	"xpf/pb"
)

type User struct {
	BaseModel

	IsGesture       bool   `gorm:"not null"`
	Gesture         []byte `gorm:"not null"`
	ValidityGesture int32  `gorm:"not null"`

	ApplicationCategoriesSeq string `gorm:"not null;size:4096"` // 分类顺序号

	MaxFileSizeM            int             `gorm:"not null"`
	ImageCompressSizeM      int             `gorm:"not null"`
	VideoCompressSizeM      int             `gorm:"not null"`
	AllowCompression        bool            `gorm:"not null"`
	ImageCompressPercentage int             `gorm:"not null"`
	ImageCompressQuality    int             `gorm:"not null"`
	VideoCompressQuality    pb.VideoQuality `gorm:"not null"`

	Score   int64     `gorm:"not null;type:bigint"`
	LoginAt time.Time `json:"not null"`
}

func (user *User) GetApplicationCategoriesSeq() ([]int64, error) {
	seqs := make([]int64, 0)

	if err := json.Unmarshal([]byte(user.ApplicationCategoriesSeq), &seqs); err != nil {
		return nil, err
	}

	return seqs, nil
}

func (user *User) SetApplicationCategoriesSeq(seqs []int64) error {
	bs, err := json.Marshal(seqs)
	if err != nil {
		return err
	}

	user.ApplicationCategoriesSeq = string(bs)
	return nil
}

func CreateUser(tx *gorm.DB, userId int64, score int64, seqs []int64) (*User, error) {
	u := &User{
		IsGesture:       false,
		Gesture:         []byte{},
		ValidityGesture: 0,

		MaxFileSizeM:            128,
		ImageCompressSizeM:      2,
		VideoCompressSizeM:      20,
		AllowCompression:        true,
		ImageCompressPercentage: 70,
		ImageCompressQuality:    10,
		VideoCompressQuality:    pb.VideoQuality_DefaultQuality,

		Score:   score,
		LoginAt: time.Now(),
	}
	u.ID = userId

	// Categories Seq
	if err := u.SetApplicationCategoriesSeq(seqs); err != nil {
		return nil, err
	}

	// Create
	db := tx.Create(u)
	if err := db.Error; err != nil {
		return nil, err
	}

	return u, nil
}

func FindUserById(tx *gorm.DB, id int64) (*User, error) {
	var user User

	db := tx.First(&user, id)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUsers(tx *gorm.DB) ([]*User, error) {
	var us []*User

	db := tx.Find(&us)
	if err := db.Error; err != nil {
		return nil, err
	}

	return us, nil
}
