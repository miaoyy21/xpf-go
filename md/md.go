package md

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"psw/ds"
	"time"
)

type Message interface {
	DealWith(*gorm.DB, *ds.User, proto.Message) (proto.Message, error)
}

func Init(db *gorm.DB) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Task() PANIC with RECOVER %s", err)
		}
	}()
	doCosts(db)

	// 初始化扣费定时任务
	ch30 := time.NewTicker(5 * time.Minute).C
	for {
		select {
		case <-ch30:
			// 执行积分消费任务
			doCosts(db)
		}
	}
}
