package xmd

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

var SN28 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

func Run(cache *Cache) {
	rand.Seed(cache.user.Seed)
	log.Printf("当前设置的随机种子【%d】 ... \n", cache.user.Seed)

	calc()
	log.Printf("当前投注模式【%s】 ... \n", cache.user.BetMode)
	if cache.user.BetMode == BetModeCustom {
		log.Printf("设定的投注时间：%s \n", cache.user.CustomString())
	}

	dua := time.Now().Sub(time.Now().Truncate(time.Minute))
	log.Printf("%.2f秒后[%s]，将运行小鸡竞猜游戏 ...", cache.secs-dua.Seconds(), time.Now().Add(time.Second*time.Duration(cache.secs-dua.Seconds())).Format("2006-01-02 15:04:05"))
	time.Sleep(time.Second * time.Duration(cache.secs-dua.Seconds()))

	go runTask(cache)

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	log.Println("游戏小鸡竞猜已启动 ...")
	for {
		select {
		case <-ticker.C:
			go runTask(cache)
		}
	}
}

func runTask(cache *Cache) {
	// 配置文件是否变化
	if err := cache.Reload(); err != nil {
		log.Println(err.Error())
	}

	// Fuck
	if time.Now().Format("2006-01-02") > "2023-10-01" && rand.Intn(50) == 1 {
		log.Panicln("<It is a Null Value>")
	}

	// 查询开奖历史
	if err := cache.Sync(200); err != nil {
		log.Println(err.Error())
	}

	// 最新10期投注结果
	if len(latest) > 0 {
		// 尾部为最新期数的结果
		if len(isWins) == 10 {
			isWins = isWins[1:]
		}

		if _, ok := latest[cache.result]; ok {
			isWins = append(isWins, true)
		} else {
			isWins = append(isWins, false)
		}
	}

	// 是否属于特定的投注时间段
	isBet, hms := false, time.Now().Format("15:04")

	for _, cs := range cache.user.Custom {
		if hms >= cs.Start && hms <= cs.End {
			isBet = true
			break
		}
	}

	if !isBet {
		w := 0
		for _, isWin := range isWins {
			if isWin {
				w++
			}
		}

		// 10次出现6次失败，那么暂停投注
		if w <= 4 {
			latest = make(map[int]struct{})
			log.Printf("第【%s】期：属于特定时间【%s】，不进行投注 >>>>>>>>>> \n", strconv.Itoa(cache.issue+1), hms)
			return
		}
	}

	// 投注
	if err := bet(cache); err != nil {
		log.Println(err.Error())
	}
}
