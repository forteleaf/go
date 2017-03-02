// 홈페이지에서 내 책 목록을 가져와서 mybook db 에 저장하기

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

type mybook struct {
	thumbnail string
	name      string
	pdf       string
	epub      string
	mobi      string
	codefiles string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func connectDB() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "packt:packt1234@tcp(305401.iptime.org:3306)/packtpub")
	checkError(err)
	trans, err := db.Begin()
	checkError(err)
	return trans, err
}
func mariaDB(data mybook) {
	tranx, _ := connectDB()
	defer tranx.Rollback()

	stmt, err := tranx.Prepare("insert into mybook values (0,?,?,?,?,?,?)")
	checkError(err)
	defer stmt.Close()

	stmt.Exec(data.thumbnail, data.name, data.pdf, data.epub, data.mobi, data.codefiles)
	err = tranx.Commit()
	checkError(err)
}

func main() {
	/* read html
	f, err := os.Open("mybook.html")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(f)
	*/
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	userinfo := url.Values{}
	userinfo.Add("email", "forteleaf@gmail.com")
	userinfo.Add("password", "shuria40")
	userinfo.Add("op", "Login")
	userinfo.Add("form_id", "packt_user_login_form")

	req, _ := http.NewRequest("POST", "https://www.packtpub.com/", strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("Login 성공")
	}
	res, _ = client.Get("https://www.packtpub.com/account/my-ebooks")
	if res.StatusCode == 200 {
		fmt.Println("내 목록 읽어오기 성공")
	} else {
		fmt.Println("실패")
		os.Exit(1)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".product-line").Each(func(i int, s *goquery.Selection) {
		thumbnail, _ := s.Find(".imagecache-thumbview").Attr("data-original")
		name := strings.TrimSpace(s.Find("div.title").Text()) // name
		var pdf, epub, mobi, codefile string
		// var pdf, epub, mobi string
		s.Find(".download-container a").Each(func(i int, ss *goquery.Selection) {
			links, _ := ss.Attr("href") // thumbnail
			if strings.Contains(links, "pdf") {
				pdf = links // pdf
			}
			if strings.Contains(links, "epub") {
				epub = links
			}
			if strings.Contains(links, "mobi") {
				mobi = links
			}
			if strings.Contains(links, "code") {
				codefile = links
			}
		})
		// fmt.Println(thumbnail, name)
		var mybook mybook
		mybook.thumbnail = thumbnail
		mybook.name = name
		mybook.pdf = pdf
		mybook.epub = epub
		mybook.mobi = mobi
		mybook.codefiles = codefile

		mariaDB(mybook)

	})
}
