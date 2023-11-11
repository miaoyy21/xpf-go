package ds

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"xpf/pb"
)

type Operate struct {
	gorm.Model

	UserId int64 `gorm:"not null;type:bigint"`

	MsgNo pb.MsgNo `gorm:"not null"`
	Msg   string   `gorm:"not null;size:2048"`
	Rt    int32    `gorm:"not null"`
	Err   string   `gorm:"not null;size:4096"`

	IP        string `gorm:"not null;size:256"`
	OS        string `gorm:"not null;size:256"`
	Device    string `gorm:"not null;size:4096"`
	UserAgent string `gorm:"not null;size:512"`
}

func (m *Operate) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_operate_on_user_id", "user_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (m *Operate) GetMsg() ([]string, error) {
	msg := make([]string, 0)

	if err := json.Unmarshal([]byte(m.Msg), &msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Operate) SetMsg(msg []string) error {
	bs, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	m.Msg = string(bs)
	return nil
}

func CreateOperate(db *gorm.DB, userId int64, msgNo pb.MsgNo, msg []string, rt int32, err, ip, os, device string, agent string) error {
	op := &Operate{
		UserId: userId,
		MsgNo:  msgNo,
		Rt:     rt,
		Err:    err,

		IP:        ip,
		OS:        os,
		Device:    device,
		UserAgent: agent,
	}

	// Set Msg
	if err := op.SetMsg(msg); err != nil {
		return err
	}

	// Create
	if err := db.Create(op).Error; err != nil {
		return err
	}

	return nil
}

func FindOperatesByUserIdAt(tx *gorm.DB, userId int64, limit, offset int32) ([]*Operate, error) {
	var ops []*Operate

	db := tx.Where(&Operate{UserId: userId}).Limit(limit).Offset(offset).Order("id desc").Find(&ops)
	if err := db.Error; err != nil {
		return nil, err
	}

	return ops, nil
}
