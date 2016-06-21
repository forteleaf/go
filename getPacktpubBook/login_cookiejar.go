package main

import (
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
)

// https://www.packtpub.com/freelearning-claim/17371/21478

func main() {
	// setting cookie
	cookieJar, _ := cookiejar.New()
	client := http.Client{Jar: cookieJar}

	resp, err := http.Postform("https://www.packtpub.com/packt/offers/free-learning", url.Values{
		"op", "Login",
		"email":    "forteleaf@gmail.com",
		"password": "shuria40",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	req, err := http.NewRequest(nil, "https://www.packtpub.com/packt/offers/free-learning", values)
	fmt.Println(resp.Body.Close)
}
