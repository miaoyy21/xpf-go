package xmd

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type IssueResult struct {
	issue  int
	result int
	money  int
	member int
}

type HGold struct {
	Time string
	Gold int
}

type Cache struct {
	dir string
	md5 []byte

	user UserBase
	secs float64

	issue  int // 最新期数
	result int // 最新开奖结果
	money  int // 最新投注金额
	member int // 最新参与人数

	histories []IssueResult // 每期存在数据库的开奖记录
	hGolds    []HGold

	Prizes map[string]string // 已兑奖 Map<奖品ID>奖品名称
}

func NewCache(dir string) (*Cache, error) {
	bs, err := os.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		return nil, err
	}

	// MD5
	h := md5.New()
	if _, err := h.Write(bs); err != nil {
		return nil, err
	}

	var conf Config
	if err := json.Unmarshal(bs, &conf); err != nil {
		return nil, err
	}

	user := NewUserBase(
		conf.Seed, conf.BetMode, conf.Cookies, conf.Agent,
		conf.Unix, conf.KeyCode, conf.DeviceId, conf.UserId, conf.Token,
	)

	secs := conf.Secs
	if conf.BetMode == BetModeModeAll || conf.BetMode == BetModeModeOnly {
		secs = 40
		log.Printf("由于设定的投注为Mode模式，将投注时间强制改为40秒～45秒\n")
	}

	cache := &Cache{
		dir: dir,
		md5: h.Sum(nil),

		user: user,
		secs: secs,

		issue:  -1,
		result: -1,
		money:  -1,
		member: -1,

		histories: make([]IssueResult, 0),
		hGolds:    make([]HGold, 0),

		Prizes: make(map[string]string),
	}

	return cache, nil
}
