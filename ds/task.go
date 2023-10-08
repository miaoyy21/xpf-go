package ds

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TaskType int

const (
	CostTask TaskType = iota + 1
)

type Task struct {
	gorm.Model

	Type     TaskType `gorm:"not null"`
	LatestAt time.Time
}

func CreateTask(tx *gorm.DB, taskType TaskType) (*Task, error) {
	task := &Task{
		Type:     taskType,
		LatestAt: time.Now(),
	}

	// Create
	db := tx.Create(task)
	if err := db.Error; err != nil {
		return nil, err
	}

	return task, nil
}

func FindTaskByType(tx *gorm.DB, taskType TaskType) (*Task, error) {
	var task Task

	db := tx.Where(&Task{Type: taskType}).First(&task)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &task, nil
}
