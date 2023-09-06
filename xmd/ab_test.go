package xmd

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCache_Sync(t *testing.T) {

	log.Println(1 << 26)
	log.Println(1 << 15)

	q1 := `forever.pceggs.com=UserID=p/++L9QfTzRyXvPBefzxxA==&Time=xjTmrbIR/RoOtgVaheK1NYIjLehU7tXX&Date=xjTmrbIR/RreMea1lb8sPA==&Status=KAyeeDyZo6Y=; re.pceggs.com=computerid=dbsmaFoiSqhyFnXs2sCuJg==&sign=295A14AF17B2AEC9A15B35262036422D; .ADWASPX7A5C561934E_PCEGGS=C48F4656F52532AA448AB89DC70D0AF52B1B96638F837EA93DF35C2F959E40651F4142CF636457D95BD03BB3EEDCFD9EE04F9C6A210F08F7A1A64808C30F189D782D014D09465B25DBF9B9910103EF1FFFA0BFF367F29D9FF166BEAB94F351FB45D4A52AA3D36FD1A73A2C1E05ABA6332F5F804562F5746461F5611B295CD929A1704BCF; Hm_lvt_f8f6a0064a3e891522bdf044119d462a=1690530552,1692150515; ckurl.pceggs.com=ckurl=http://www.pceggs.com/game/gameindex/gameindex.aspx?gameid=4; Hm_lpvt_f8f6a0064a3e891522bdf044119d462a=1692150527`
	printCookie(q1)

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

		nv := v
		if strings.Contains(k, "lpvt") {
			nv = strconv.Itoa(int(time.Now().Unix()))
		}

		log.Printf("%40s : %s \n", k, nv)

		ns = append(ns, strings.Join([]string{k, nv}, "="))
	}

	log.Printf("New Cookies String: \n %s \n", strings.Join(ns, "; "))
}
