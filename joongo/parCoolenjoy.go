package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

// http://www.coolenjoy.net/login_check?mb_id=illuadel&mb_password=404040
// login : mb_id=illuadel&mb_password=404040

const (
	LoginURL  = "http://www.cooln.net/login_check"
	marketURL = "http://www.cooln.net/bbs/mart2"
)

type BoardCoolen struct {
	num      int
	link     string
	category string
	title    string
	writer   string
	price    string
	date     string
	hit      int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	checkError(err)
	trans, err := db.Begin()
	checkError(err)
	return trans, err
}

// mariaDb 에 insert
func mariaDB(data []BoardCoolen) {
	tranx, _ := connectDB()
	defer tranx.Rollback()

	//"insert into joongo values (0,'clien',?,?,?,?,'x',?,?) on DUPLICATE KEY UPDATE category=?, hit=?"
	stmt, err := tranx.Prepare("insert into joongo values (concat('cool_',?),'',?,?,?,?,?,?)" +
		"on DUPLICATE KEY UPDATE category=?, price=?, hit=?")
	checkError(err)
	defer stmt.Close()

	for k, _ := range data {
		_, err := stmt.Exec(data[k].num, data[k].category, data[k].title, data[k].writer, data[k].price, data[k].date, data[k].hit,
			data[k].category, data[k].price, data[k].hit)
		checkError(err)
	}
	err = tranx.Commit()
	checkError(err)
}

// login and take market return
func login_coolenjoy() *http.Response {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	userinfo := url.Values{}
	userinfo.Add("mb_id", "illuadel")
	userinfo.Add("mb_password", "404040")
	req, _ := http.NewRequest("POST", LoginURL, strings.NewReader(userinfo.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		fmt.Println("접속성공")
	}

	res, _ = client.Get(marketURL)
	if res.StatusCode == 200 {
		fmt.Println("read market")
	} else {
		fmt.Println("fail")
		os.Exit(1)
	}
	return res
}

// goquery
func SoldListCoolen() []BoardCoolen {
	arr := []BoardCoolen{}
	doc, err := goquery.NewDocumentFromResponse(login_coolenjoy())
	checkError(err)
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		category := strings.Trim(s.Find(".td_num font").Text(), " ")
		link, _ := s.Find(".td_subject a").Attr("href")
		lastslash := strings.LastIndex(link, "/")
		num, _ := strconv.Atoi(link[lastslash+1:]) // link 파일에서 번호 추출
		title := s.Find(".td_subject a").Text()
		idx := strings.Index(title, "댓글")
		if idx != -1 {
			title = strings.TrimSpace(title[:idx])
		} else {
			title = "비공개"
		}
		// fmt.Println("[" + title + "]")
		// price := strings.Join(strings.Split(strings.Trim(s.Find(".td_won").Text(), "원"), ","), "")
		price := s.Find(".td_won").Text()
		writer := strings.TrimSpace(s.Find(".sv_member").Text())
		now := time.Now()
		// date := now.String()[:10] + " " + s.Find(".td_date").Text()
		date := now.String()
		// fmt.Println(now.String())
		fmt.Println(date)
		hit, _ := strconv.Atoi(s.Find(".td_hit").Text())
		arr = append(arr, BoardCoolen{num, link, category, title, writer, price, date, hit})
	})
	// fmt.Println(arr[2])
	return arr[2:]
}

func main() {
	arr := SoldListCoolen()
	mariaDB(arr)
}
