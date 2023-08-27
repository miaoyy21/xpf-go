package xmd

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func (o *Cache) Reload() error {
	// 加载配置文件
	if err := o.reloadConfig(); err != nil {
		return err
	}

	// 加载兑奖
	if err := o.reloadPrizes(); err != nil {
		return err
	}

	return nil
}

func (o *Cache) reloadConfig() error {

	bs, err := os.ReadFile(filepath.Join(o.dir, "config.json"))
	if err != nil {
		return err
	}

	// MD5
	h := md5.New()
	if _, err := h.Write(bs); err != nil {
		return err
	}

	var conf Config
	if err := json.Unmarshal(bs, &conf); err != nil {
		return err
	}

	if bytes.Equal(h.Sum(nil), o.md5) {
		return nil
	}

	if o.user.Seed != conf.Seed {
		rand.Seed(o.user.Seed)
	}

	user := NewUserBase(
		conf.Seed, conf.BetMode, conf.Custom, conf.Cookies, conf.Agent,
		conf.Unix, conf.KeyCode, conf.DeviceId, conf.UserId, conf.Token,
	)

	o.md5 = h.Sum(nil)
	o.user = user

	log.Println("配置文件变化，重新加载配置文件完成 ...")
	log.Printf("当前初始化随机种子【%d】 ... \n", o.user.Seed)
	log.Printf("当前是否启用设定投注模式【%s】 ... \n", o.user.BetMode)
	return nil
}

func (o *Cache) reloadPrizes() error {
	dFileName := filepath.Join(o.dir, "data.json")
	if _, err := os.Stat(dFileName); err != nil {
		if os.IsNotExist(err) {
			o.Prizes = make(map[string]string)

			// 创建文件
			dFile, err := os.Create(dFileName)
			if err != nil {
				log.Panicf("unexpected :: os.CreateFile(%s) failure : %s \n", dFileName, err.Error())
			}
			defer dFile.Close()

			js := json.NewEncoder(dFile)
			js.SetIndent("", "\t")
			if err := js.Encode(o.Prizes); err != nil {
				return err
			}

			return nil
		}

		log.Panicf("unexpected :: os.Stat(%s) failure : %s \n", dFileName, err.Error())
	}

	dFile, err := os.Open(dFileName)
	if err != nil {
		return err
	}
	defer dFile.Close()

	if err := json.NewDecoder(dFile).Decode(&o.Prizes); err != nil {
		return err
	}

	return nil
}
