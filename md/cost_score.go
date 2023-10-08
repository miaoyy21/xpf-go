package md

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psw/ds"
	"psw/pb"
	"time"
)

func doCosts(db *gorm.DB) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("doCosts() PANIC with RECOVER %s", err)
		}
	}()

	// 查找积分消费任务
	task, err := ds.FindTaskByType(db, ds.CostTask)
	if err != nil {
		logrus.Errorf("doCosts() ds.FindTaskByType() %s", err.Error())
		return
	}

	// 每天 00:00 ~ 00:05 定时执行
	logrus.Debugf("Latest Run at %s , Now at %s", task.LatestAt, time.Now())
	if ds.IsSameDay(task.LatestAt, time.Now()) || task.LatestAt.After(time.Now()) {
		return
	}

	logrus.Info("Running Cost Score Task ...")

	// 检索所有用户
	users, err := ds.FindUsers(db)
	if err != nil {
		logrus.Errorf("doCosts() ds.FindUsers() %s", err.Error())
		return
	}

	tx := db.Begin()
	for _, user := range users {

		// 用户的所有应用
		accounts, err := ds.FindApplicationAccountsByUserId(tx, user.ID)
		if err != nil {
			tx.Rollback()
			logrus.Errorf("doCosts() ds.FindApplicationAccountsByUserId(%d) %s", user.ID, err.Error())
			return
		}

		// 根据应用上传的文件占用的存储空间
		var size int64 = 0
		var accountN int64 = 0
		for _, account := range accounts {
			if len(account.AccountId) == 0 {
				continue
			}

			accountN++
			size += int64(account.Size)
		}

		// 用户的所有文件占用的存储空间
		files, err := ds.FindFilesByUserId(tx, user.ID)
		if err != nil {
			tx.Rollback()
			logrus.Errorf("doCosts() ds.FindFilesByUserId(%d) %s", user.ID, err.Error())
			return
		}

		// 根据存储空间扣费
		for _, file := range files {
			if file.Type != pb.FileType_FileCabinet {
				continue
			}

			size += int64(file.Size)
		}
		cost := accountN*ds.CostAccountScore + size*ds.CostFile10MScore/(10<<20)

		logrus.Debugf("COST is %d", cost)

		// 扣分记录
		if cost != 0 {
			score := user.Score + cost
			if err := ds.CreateScore(tx, user.ID, ds.CostScoreAction, cost, score); err != nil {
				tx.Rollback()
				logrus.Errorf("doCosts() ds.CreateScore(%d) %s", user.ID, err.Error())
				return
			}

			// 保存用户积分
			user.Score = score
			if err := tx.Save(user).Error; err != nil {
				tx.Rollback()
				logrus.Errorf("doCosts() tx.Save.User(%d) %s", user.ID, err.Error())
				return
			}
		}
	}

	task.LatestAt = time.Now()
	if err := tx.Save(task).Error; err != nil {
		tx.Rollback()
		logrus.Errorf("doCosts() tx.Save.Task(%d) %s", task.Type, err.Error())
		return
	}

	tx.Commit()
}
