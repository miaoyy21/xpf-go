package xmd

import (
	"log"
	"regexp"
	"strings"
)

func prize(cache *Cache) error {
	// 获取HTML
	bs, err := hDoPrizeHTML(cache.user)
	if err != nil {
		return err
	}

	lis := regexp.MustCompile(`<li class="spzs">(.*?)</li>`).FindAllString(string(bs), -1)
	log.Println("Find All String List ::")

	reId := regexp.MustCompile(`onclick="window.open\('` + URLOpenPrize + `\?id=([^']*)'\)"`)
	reName := regexp.MustCompile(`<div class="sp_name">(.*?)</div>`)
	reGold := regexp.MustCompile(`<div class="db_xh">(.*?)</div>`)
	rePercent := regexp.MustCompile(`<div class="jd_wz">(.*?)</div>`)
	for i, li := range lis {
		// ID
		sId := reId.FindStringSubmatch(li)[1]

		// 名称
		sName := strings.TrimSpace(reName.FindStringSubmatch(li)[1])
		sName = strings.TrimSpace(sName)

		// 数量
		sGold := strings.TrimSpace(reGold.FindStringSubmatch(li)[1])
		sGold = strings.ReplaceAll(strings.ReplaceAll(sGold, "需要：", ""), "<span>", "")
		sGold = strings.ReplaceAll(sGold, ",", "")
		sGold = strings.ReplaceAll(strings.ReplaceAll(sGold, "</span>", ""), "金蛋", "")
		sGold = strings.TrimSpace(sGold)

		// 进度
		sPercent := strings.TrimSpace(rePercent.FindStringSubmatch(li)[1])

		// 输出
		log.Printf("\n %02d: \n %-10s %-20s \n %-10s %-20s \n", i+1, sId, sName, sGold, sPercent)
	}

	return nil
}
