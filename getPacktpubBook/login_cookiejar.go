package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com/"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

func main() {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	userinfo := url.Values{}
	userinfo.Add("email", "forteleaf@gmail.com")
	userinfo.Add("password", "shuria40")
	userinfo.Add("op", "Login")
	userinfo.Add("form_id", "packt_user_login_form")

	req, _ := http.NewRequest("POST", packtPubFreeURL, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)

	fmt.Println(client.Jar)
	res, _ = client.Get("https://www.packtpub.com/freelearning-claim/16784/21478")
	// req, _ = http.NewRequest("GET", "https://www.packtpub.com/freelearning-claim/16784/21478", nil)
	// res, _ = client.Do(req)
	// https://www.packtpub.com/freelearning-claim/16784/21478
	fmt.Println("===============" + res.Status)
	fmt.Println(client.Jar)
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
