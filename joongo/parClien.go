package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

type BoardClien struct {
	num      int // 글 번호
	category string
	title    string
	user     string
	date     string
	link     string
	hit      int
}

// error 처리
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func clienSoldList() []BoardClien {
	arr := []BoardClien{}
	for i := 0; i < 3; i++ {
		doc, err := goquery.NewDocument("http://clien.net/cs2/bbs/board.php?bo_table=sold&page=" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("tbody tr.mytr").Each(func(i int, s *goquery.Selection) {
			num, _ := strconv.Atoi(s.Find("td").First().Text())
			link, _ := s.Find(".post_subject a").Attr("href")
			category := s.Find(".post_category a").Text()
			category = category[1 : len(category)-1]
			title := s.Find(".post_subject a").Text()
			user, ok := s.Find(".post_name img").Attr("src")
			if !ok {
				user = s.Find(".member").Text()
			}
			// 이상하게 time.parse 가 안됨. string 으로 저장을 하고, 이를 그냥 넣는게 속편함
			date, _ := s.Find("td").Last().Prev().Children().Attr("title")
			hit, _ := strconv.Atoi(s.Find("td").Last().Text())
			arr = append(arr, BoardClien{num, category, title, user, date, link, hit})
		})
	}
	return arr
}

// connectDB 에 연결해서 트랜잭션 시작
func connectDB() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	checkError(err)
	trans, err := db.Begin()
	checkError(err)
	return trans, err
}

// mariaDb 에 insert
func mariaDB(data []BoardClien) {
	tranx, _ := connectDB()
	defer tranx.Rollback()

	// prepared statement
	stmt, err := tranx.Prepare("insert into joongo values (concat('clin_',?),'',?,?,?,'',?,?)" +
		"on DUPLICATE KEY UPDATE category=?, title =?, hit=?")
	// stmt, err := tranx.Prepare("begin tran")
	checkError(err)
	defer stmt.Close()

	for k, _ := range data {
		_, err := stmt.Exec(data[k].num, data[k].category, data[k].title, data[k].user, data[k].date, data[k].hit,
			data[k].category, data[k].title, data[k].hit)
		checkError(err)
	}
	err = tranx.Commit()
	checkError(err)
}

func main() {
	arr := clienSoldList()
	mariaDB(arr)
}
