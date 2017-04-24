package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	packtPubBaseURL = "https://www.packtpub.com"
	packtPubFreeURL = "https://www.packtpub.com/packt/offers/free-learning"
)

type Packtpubconfig struct {
	Date string
	URL  string
}

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
	} else {
		fmt.Println("실패")
	}
}

// checkPacktpubConfig is check today URL
func checkPacktpubConfig() {
	var config Packtpubconfig

	file, err := os.Open("packtpub.ini")
	// doesn't exist pckypub.ini
	if err != nil {
		makeConfigFile()
	} else {
		// exist packtpub.ini but not today config
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if config.Date != time.Now().String()[:10] {
			makeConfigFile()
		}
	}
}

// makeConfigFile is make packypub.ini file
func makeConfigFile() {
	file, err := os.Create("packtpub.ini")
	chkError(err)
	defer file.Close()

	getFreeBookURL, _ := GetFreeBookURL()
	getFreeBookURL = packtPubBaseURL + getFreeBookURL
	fmt.Println("주소를 얻어왔습니다\n", getFreeBookURL)
	today := time.Now().String()[:10]
	aa := &Packtpubconfig{
		Date: today,
		URL:  getFreeBookURL,
	}
	encoder, _ := json.Marshal(aa)
	file.Write(encoder)
}

// error chkeck
func chkError(err error) {
	if err != nil {
		log.Fatal()
	}
}

// readConfig is read today url
func readConfig() (string, error) {
	var config Packtpubconfig
	file, err := os.Open("packtpub.ini")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config.URL, err
}

func main() {
	var email string
	var pwd string
	flag.StringVar(&email, "email", "none", "")
	flag.StringVar(&pwd, "pwd", "none", "")
	flag.Parse()
	if email == "none" || pwd == "none" {
		fmt.Println("not exist email, password\nprogram exit")
		os.Exit(0)
	}
	// check today URL file
	checkPacktpubConfig()

	addr, err := readConfig()
	chkError(err)
	Auth(email, pwd, addr)
}
