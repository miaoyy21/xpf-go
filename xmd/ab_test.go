package xmd

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestCache_Sync(t *testing.T) {

	log.Println(1 << 26)

	h5 := strings.ReplaceAll(string(tests), "\n", "")

	// 获取夺宝列表
	lis := regexp.MustCompile(`<li class="spzs">(.*?)</li>`).FindAllString(h5, -1)
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
}

func printCookie(q string) {
	qs := strings.Split(q, "; ")

	log.Println("Cookies Pairs: ")
	ns := make([]string, 0)
	for _, s := range qs {
		s0 := strings.Split(s, "=")
		k, v := s0[0], strings.Join(s0[1:], "=")

		log.Printf("%40s : %s \n", k, v)
		ns = append(ns, strings.Join([]string{k, v}, "="))
	}

	log.Printf("New Cookies String: \n %s \n", strings.Join(ns, "; "))
}
