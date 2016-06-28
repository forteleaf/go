package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
	// What You Need to Know about Docker - Free eBook
)

func main() {
	var cookies []*http.Cookie
	u, _ := url.Parse(packtPubBaseURL)
	cookieJar, _ := cookiejar.New(nil)
	cookieJar.SetCookies(u, cookies)

	userinfo := url.Values{}
	userinfo.Add("email", "forteleaf@gmail.com")
	userinfo.Add("password", "shuria40")
	userinfo.Add("op", "Login")

	client := &http.Client{Jar: cookieJar}

	req, _ := http.NewRequest("POST", "packtPubFreeURL", bytes.NewBufferString(userinfo.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//header 를 하니까 cookie 가 제대로 생성된다.

	// body io.reader
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("200 접속 성공")
		fmt.Println("cookie 확인", cookieJar)
	}
	// https://www.packtpub.com/freelearning-claim/25854/21478
	//resp, _ := client.Get("https://www.packtpub.com/freelearning-claim/25854/21478")
	client.CheckRedirect(req, res)
	fmt.Println(resp.StatusCode)

	// client.Head("https://www.packtpub.com/freelearning-claim/25854/21478")
	// client.PostForm("https://www.packtpub.com/freelearning-claim/25854/21478", userinfo)
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
