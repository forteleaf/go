package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	packtPubBaseURL = "https://www.packtpub.com/"
	packtPubFreeURL = "https://www.packtpub.com/packt/offers/free-learning"
)

// getBody is get []byte body
func getBody() (*bytes.Buffer, error) {
	resp, err := http.Get(packtPubBaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(body)
	return buf, nil
}

// GetFreeBookURL 은 책구매하는 주소
func GetFreeBookURL() (url string, err error) {
	fmt.Println("책주소를 얻어오는 중입니다.")
	doc, err := goquery.NewDocument(packtPubFreeURL)
	if err != nil {
		return "", err
	}
	doc.Find("a.twelve-days-claim").Each(func(i int, s *goquery.Selection) {
		node := s.First()
		url, _ = node.Attr("href")
	})
	return url, err
}

func inputInfo() (string, string) {
	var email, password string
	fmt.Println("Packtpub 매일 받아오기")
	fmt.Println("이메일을 입력하세요")
	fmt.Scan(&email)
	match, _ := regexp.MatchString(".+@.+\\..+", email)
	if !match {
		fmt.Println("이메일 형식이 틀렸습니다")
		os.Exit(0)
	}
	fmt.Println("비밀번호를 입력하세요")
	fmt.Scan(&password)
	return email, password
}

func Auth(email string, password string, FreeBookURL string) {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	userinfo := url.Values{}
	userinfo.Add("email", email)
	userinfo.Add("password", password)
	userinfo.Add("op", "Login")
	userinfo.Add("form_id", "packt_user_login_form")

	req, _ := http.NewRequest("POST", packtPubBaseURL, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("Login 성공")
	}

	res, _ = client.Get(FreeBookURL)
	if res.StatusCode == 200 {
		fmt.Println("책 구매하기 성공")
	}
}

func main() {
	email, password := inputInfo()
	getFreeBookURL, _ := GetFreeBookURL()
	getFreeBookURL = packtPubBaseURL + getFreeBookURL
	fmt.Println(email, password, getFreeBookURL)

	Auth(email, password, getFreeBookURL)
}
