package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

func main() {
	u, _ := url.Parse(packtPubBaseURL)
	cookieJar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookieJar.SetCookies(u, cookies)

	userinfo := url.Values{}
	userinfo.Add("email", "forteleaf@gmail.com")
	userinfo.Add("password", "shuria40")
	userinfo.Add("op", "Login")

	client := &http.Client{
		Jar: cookieJar,
	}
	req, _ := http.NewRequest("POST", packtPubFreeURL, strings.NewReader(userinfo.Encode()))
	// res, _ := client.PostForm(packtPubBaseURL, userinfo)
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//header 를 하니까 cookie 가 제대로 생성된다.

	// body io.reader
	res, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	// https://www.packtpub.com/freelearning-claim/19350/21478
	fmt.Println(res.Status)
	resp, _ := client.Get("https://www.packtpub.com/freelearning-claim/19350/21478")

	// req1, _ := http.NewRequest("PUT", "https://www.packtpub.com/freelearning-claim/19350/21478", bytes.NewBufferString(userinfo.Encode()))
	// req1.Header.Add("User-Agent", "Mozilla/5.0")
	// req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// resp, _ := client.Get("https://www.packtpub.com/freelearning-claim/19350/21478")
	// resp, _ := client.Do(req1)
	fmt.Println("============================" + resp.Status)

	// res, _ := http.PostForm(packtPubBaseURL, userinfo)
	//
	// if res.StatusCode == 200 {
	// 	cookies := res.Cookies()
	// 	for k, v := range cookies {
	// 		fmt.Println(k, v)
	// 	}
	// 	fmt.Println("접속에 성공했습니다.")
	// }
	// defer res.Body.Close()
	// fmt.Print(res.Status)
	// http.Get("https://www.packtpub.com/freelearning-claim/25854/21478")
	// fmt.Print(string(body))
}
