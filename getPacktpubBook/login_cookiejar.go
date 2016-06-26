package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

func main() {
	data := url.Values{"email": "forteleaf@gmail.com",
		"password": "shuria40",
		"op":       "Login",
	}

	res, _ := http.PostForm(packtPubBaseURL, data)

	if res.StatusCode == 200 {
		cookies := res.Cookies()
		for k, v := range cookies {
			fmt.Println(k, v)
		}
		fmt.Println("접속에 성공했습니다.")
	}
	defer res.Body.Close()
	// fmt.Print(string(body))
}
