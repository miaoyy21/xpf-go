package xmd

import (
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

var isWins = make([]bool, 0, 10)
var latest = make(map[int]struct{})

func bet(cache *Cache) error {
	issue := strconv.Itoa(cache.issue + 1)
	if cache.user.BetMode == BetModeMode {
		ms := rand.Intn(5000)

		log.Printf("第【%d】期：在Mode投注模式下，%9.2f秒再进行投注\n", cache.issue, float64(ms)/1000)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

	// 当前账户可用余额
	surplus, err := hGetGold(cache.user)
	if err != nil {
		return err
	}

	// 显示当前中奖情况
	log.Printf("⭐️⭐️⭐️ 第【%d】期：开奖结果【%d】，余额【%d】，开始执行分析 ...\n", cache.issue, cache.result, surplus)

	var coverage float64

	bets, aBets := make(map[int]float64), make(map[int]float64)
	for _, result := range SN28 {
		r0 := rand.Float64()
		r1 := rand.Float64()

		aBets[result] = r1 / r0

		var rx float64
		if r1/r0 >= 1.0 {
			rx = 1.0
		} else {
			rx = (r1/r0 - 0.99) * 100.0
		}

		if rx <= 0.01 {
			log.Printf("第【%s】期：竞猜数字【   %02d】，幸运值不够 \n", issue, result)
			continue
		}

		if rx >= 1.0 {
			log.Printf("第【%s】期：竞猜数字【 H %02d】，很幸运被选中 \n", issue, result)
		} else {
			log.Printf("第【%s】期：竞猜数字【 L %02d】，幸运值比较弱 \n", issue, result)
		}

		bets[result] = rx
		coverage = coverage + float64(stds[result])*rx
	}

	if rand.Float64() < 0.05 {
		log.Printf("第【%s】期：总体幸运值不够，不进行投注 >>>>>>>>>> \n", issue)
		latest = make(map[int]struct{})
		return nil
	}

	// 倍率
	mrx := 1.0
	if cache.money < 1<<27 {
		mrx = float64(cache.money) / float64(1<<27) // 134,217,728
	}

	// 设置的投注金额
	var m1Gold int
	if surplus < 1<<26 && cache.user.BetMode != BetModeMode {
		m1Gold = surplus / 150
	} else {
		m1Gold, err = hCustomModes(cache.user)
		if err != nil {
			return err
		}
	}

	latest = make(map[int]struct{})
	if m1Gold*2 <= MinBetGold {
		log.Printf("第【%s】期：投注金额【%6d】小于设定的最小金额【%6d】，不进行投注 >>>>>>>>>> \n", issue, m1Gold, MinBetGold)
		return nil
	}

	switch cache.user.BetMode {
	case BetModeAll:
		if err := betSingle(cache, issue, mrx, m1Gold, bets); err != nil {
			return err
		}
	case BetModeWork:
		hms := time.Now().Format("15:04")
		if (hms >= "08:30" && hms <= "11:30") || (hms >= "14:30" && hms <= "17:00") || (hms >= "19:30" && hms <= "21:30") {
			w := 0
			for _, isWin := range isWins {
				if isWin {
					w++
				}
			}

			// 10次出现7次失败，那么暂停投注
			if w <= 3 {
				mrx = 0.00001
				log.Printf("第【%s】期：属于特定时间【%s】，不进行投注 >>>>>>>>>> \n", issue, hms)
			}
		}

		if err := betSingle(cache, issue, mrx, m1Gold, bets); err != nil {
			return err
		}
	case BetModeMode:
		if err := betMode(cache, issue, m1Gold, bets); err != nil {
			return err
		}
	case BetModeHalf:
		if err := betHalf(cache, issue, mrx, m1Gold, aBets); err != nil {
			return err
		}
	}

	return nil
}

// 使用基于投注模式方式投注
func betMode(cache *Cache, issue string, m1Gold int, bets map[int]float64) error {
	rs := make([]int, 0, len(bets))
	for result := range bets {
		rs = append(rs, result)
	}

	// 数字排序
	sort.Ints(rs)
	log.Printf("第【%s】期：预投注数字【%s】 >>>>>>>>>> \n", issue, fmtIntSlice(rs))

	// 确定投注模式ID
	md := 400
	modeId, modeName := modeFn(bets, md)
	if modeId == 0 {
		log.Printf("第【%s】期：所有模式权重均不超过%d，的无法确定投注模式，暂不投注 >>>>>>>>>> \n", issue, md)
		latest = make(map[int]struct{})
		return nil
	}

	log.Printf("第【%s】期：使用投注模式【%s】 >>>>>>>>>> \n", issue, modeName)
	if err := hModesBetting(issue, modeId, cache.user); err != nil {
		return err
	}

	// 投注模式之外的数字
	ams, extras := extraFn(modeId, m1Gold, bets)
	if len(extras) > 0 {
		log.Printf("第【%s】期：额外投注数字【%s】>>>>>>>>>> \n", issue, fmtIntSlice(m2sFn(extras)))
	}

	// 使用单数字投注模式，必须使用其提供的标准投注金额
	stdBets := []int{200000, 50000, 10000, 5000, 2000, 1000, 500}
	betMaps := make(map[int][]int)

	for _, stdBet := range stdBets {
		betSlice, ok := betMaps[stdBet]
		if !ok {
			betSlice = make([]int, 0)
		}

		for result, betGold := range extras {
			qn := betGold / stdBet
			if qn > 0 {
				for i := 0; i < qn; i++ {
					betSlice = append(betSlice, result)
				}

				extras[result] = betGold - qn*stdBet
			}
		}

		sort.Ints(betSlice)
		betMaps[stdBet] = betSlice
	}

	// 单数字投注
	latest = ams
	for _, stdBet := range stdBets {
		if len(betMaps[stdBet]) > 0 {
			log.Printf("第【%s】期：押注金额【%-6d】，押注数字【%s】，投注成功 >>>>>>>>>> \n", issue, stdBet, fmtIntSlice(betMaps[stdBet]))
		}

		for _, result := range betMaps[stdBet] {
			latest[result] = struct{}{}
			if err := hBetting1(issue, stdBet, result, cache.user); err != nil {
				return err
			}
		}
	}

	return nil
}

func betSingle(cache *Cache, issue string, mrx float64, m1Gold int, bets map[int]float64) error {
	for _, result := range SN28 {
		if _, ok := bets[result]; !ok || mrx*bets[result] <= 0.01 {
			continue
		}

		latest[result] = struct{}{}
		betGold := int(mrx * bets[result] * float64(2*m1Gold) * float64(stds[result]) / 1000)
		if err := hBetting1(issue, betGold, result, cache.user); err != nil {
			return err
		}
	}

	return nil
}

func betHalf(cache *Cache, issue string, mrx float64, m1Gold int, aBets map[int]float64) error {
	type BetResult struct {
		Result int
		Rx     float64
	}

	rs := make([]BetResult, 0, len(aBets))
	for result, rx := range aBets {
		rs = append(rs, BetResult{Result: result, Rx: rx})
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i].Rx > rs[j].Rx })

	var coverage float64

	ns := make([]int, 0)
	for _, s := range rs {
		if coverage > 500 {
			break
		}

		betGold := int(mrx * float64(2*m1Gold) * float64(stds[s.Result]) / 1000)
		if err := hBetting1(issue, betGold, s.Result, cache.user); err != nil {
			return err
		}

		latest[s.Result] = struct{}{}
		ns = append(ns, s.Result)
		coverage = coverage + float64(stds[s.Result])
	}
	log.Printf("第【%s】期：已投注数字【%s】，覆盖率【%.2f%%】 >>>>>>>>>> \n", issue, fmtIntSlice(ns), coverage/10)

	return nil
}
