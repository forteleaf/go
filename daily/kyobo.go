package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// 20160815 의 링크 http://www.kyobobook.co.kr/prom/2015/book/151116_dailyCheck.jsp?orderClick=c1j

// javascript:goSubmit('001');
const (
	kyobologinurl = "https://www.kyobobook.co.kr/login/login.laf"
)

func main() {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	userinfo := url.Values{}
	userinfo.Add("mb_id", "illuadel")
	userinfo.Add("mb_password", "shuria40")
	userinfo.Add("Kc", "GNHHNOlogin")
	req, _ := http.NewRequest("POST", kyobologinurl, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("접속성공")
	}

	res, _ = client.Get("http://www.kyobobook.co.kr/prom/2015/book/151116_dailyCheck_insert.jsp?eventGb=001")
	if res.StatusCode == 200 {
		fmt.Println("출석체크 성공")
	}

}
