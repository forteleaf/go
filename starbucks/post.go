package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	// get
	firstURL = "http://first.wifi.olleh.com/starbucks/index_en_new.html"
	// post
	secondURL = "https://first.wifi.olleh.com/starbucks/starbucks_en.php"
	postURL   = "https://first.wifi.olleh.com/starbucks/issue.php"
	// postURL = "http://www.istarbucks.co.kr/wireless/wireless.asp"
)

func getMacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("error")
	}
	mac := interfaces[3].HardwareAddr.String()
	macmap := strings.Split(interfaces[3].HardwareAddr.String(), ":")
	// macmap = [68 a8 6d 0b 08 82]
	mac = strings.Join(macmap, "")
	// 68a86d0b0882
	return mac
}

func main() {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	mac := getMacAddr()

	firstinfo := url.Values{}
	firstinfo.Add("firstflag", "starbucks")
	// 211.195.238.25 is 웨스턴돔 스타벅스
	firstinfo.Add("ip", "211.195.238.25")
	firstinfo.Add("mac", mac)
	firstinfo.Add("url", "http://client.net")

	userinfo := url.Values{}
	userinfo.Add("firstflag", "albert")
	userinfo.Add("lang", "en")
	userinfo.Add("devicecode", "pc")
	userinfo.Add("cust_email_addr", "89363@naver.com")
	userinfo.Add("cust_hp_cp", "l")
	userinfo.Add("ip", "211.195.238.25")
	userinfo.Add("branchflag", "스타벅스-웨스턴돔")
	userinfo.Add("mac", mac)

	// 일단 아무사이트에 접속을 한다.
	client.Get("http://first.wifi.olleh.com/webauth/index.html?ip=211.195.238.25&mac=68a86d0b0882&url=http://clien.net&ssid=olleh_starbucks&ap_mac=0a300d890032&apmodel=MW-6900AP")
	req, _ := http.NewRequest("POST", secondURL, strings.NewReader(firstinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.8,en-US;q=0.6,en;q=0.4")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("두번째 접속에 성공했습니다.")
	}

	req, _ = http.NewRequest("POST", postURL, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ = client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("접속에 성공하였습니다.")
	}

	//firstflag:starbucks
	//branchflag:스타벅스-뻥뻥
	//lang:en
	//devicecode:pc
	//userNm:albert						// name
	//cust_email_addr:896661810@qq.com 	// email
	//cust_hp_cp:l 						// 로밍, 로컬
	//ip:220.88.122.67					// IP
	//mac:68a86d0b0882					// mac address

}
