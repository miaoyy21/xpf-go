package xmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func hDo(user UserBase, method string, url RequestURL, s interface{}, t interface{}) error {
	buf := new(bytes.Buffer)

	// JSON Encode
	if err := json.NewEncoder(buf).Encode(s); err != nil {
		return err
	}

	// Sync
	req, err := http.NewRequest(method, string(url), buf)
	if err != nil {
		return err
	}

	// Sync Header
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("User-Agent", user.agent)

	if strings.Contains(string(url), "app") {
		req.Header.Set("Origin", OriginAPP)
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		req.Header.Set("Cookie", buildCookie(user.cookies.Bet))
	} else {
		req.Header.Set("Origin", OriginWWW)
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", buildCookie(user.cookies.Prize))

		req.Header.Set("Referer", string(URLPrizeIndexList))
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}

	// Response
	http.DefaultClient.Timeout = 3 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// JSON Decode
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return err
	}

	return nil
}

func hDoPrizeHTML(user UserBase) ([]byte, error) {
	// Sync
	req, err := http.NewRequest("GET", string(URLPrizeIndexList), nil)
	if err != nil {
		return nil, err
	}

	// Sync Header
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", buildCookie(user.cookies.Prize))
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", string(URLPrizeIndexList))
	req.Header.Set("User-Agent", user.agent)
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// Response
	http.DefaultClient.Timeout = 10 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func buildCookie(q string) string {
	qs := strings.Split(q, "; ")

	ns := make([]string, 0)
	for _, s := range qs {
		s0 := strings.Split(s, "=")
		k, v := s0[0], strings.Join(s0[1:], "=")

		nv := v
		if strings.Contains(k, "lpvt") {
			nv = strconv.Itoa(int(time.Now().Unix() - 1))
		}

		ns = append(ns, strings.Join([]string{k, nv}, "="))
	}

	return strings.Join(ns, "; ")
}
