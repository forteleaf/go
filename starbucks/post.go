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
	// postURL = "https://first.wifi.olleh.com/starbucks/issue.php"
	postURL = "http://www.istarbucks.co.kr/wireless/wireless.asp"
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

	userinfo := url.Values{}
	userinfo.Add("firstlag", "albert")
	userinfo.Add("lang", "en")
	userinfo.Add("devicecode", "pc")
	userinfo.Add("cust_email_addr", "89363@naver.com")
	userinfo.Add("cust_hp_cp", "l")
	userinfo.Add("ip", "220.88.122.67")
	userinfo.Add("branchflag", "스타벅스-일산후곡")
	userinfo.Add("mac", getMacAddr())

	req, _ := http.NewRequest("POST", postURL, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
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
