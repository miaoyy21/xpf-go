package ds

import (
	"github.com/jinzhu/gorm"
)

type ScoreAction int

const (
	RegisterScoreAction ScoreAction = 0
	LoginScoreAction    ScoreAction = 1
	CostScoreAction     ScoreAction = 2
	RestoreScoreAction  ScoreAction = 3
	PurchaseScoreAction ScoreAction = 9
)

const (
	RegisterScore    int64 = 3000
	LoginScore       int64 = 100
	CostAccountScore int64 = -3
	CostFile10MScore int64 = -9
)

type Score struct {
	gorm.Model

	UserId int64 `gorm:"not null;type:bigint"`

	Action ScoreAction `gorm:"not null"`
	Cost   int64       `gorm:"not null;type:bigint"`
	Score  int64       `gorm:"not null;type:bigint"`
}

func (m *Score) AfterCreateTable(db *gorm.DB) error {
	tx := db.Model(m).AddIndex("idx_scores_on_user_id", "user_id")
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func CreateScore(tx *gorm.DB, userId int64, action ScoreAction, cost int64, score int64) error {
	s := &Score{
		UserId: userId,

		Action: action,
		Cost:   cost,
		Score:  score,
	}

	// Create
	db := tx.Create(s)
	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func FindScoresByUserId(tx *gorm.DB, userId int64, limit, offset int32) ([]*Score, error) {
	var ss []*Score

	db := tx.Where(&Score{UserId: userId}).Limit(limit).Offset(offset).Order("id desc").Find(&ss)
	if err := db.Error; err != nil {
		return nil, err
	}

	return ss, nil
}
