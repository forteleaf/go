// 중고장터 태블릿 관련 판매
// http://cafe.naver.com/joonggonara.cafe?iframe_url=/ArticleList.nhn%3Fsearch.clubid=10050146%26search.boardtype=L%26search.menuid=749%26search.questionTab=A%26search.marketBoardTab=D%26search.specialmenutype=%26userDisplay=50
package main

import (
	"database/sql"
	"log"

	"github.com/PuerkitoBio/goquery"
	// _ "github.com/go-sql-driver/mysql"
)

type BoardNaverTablet struct {
	Num      int
	Category string
	Title    string
	User     string
	Date     string
	Hit      int
}

func NaverTabletList() []BoardNaverTablet {
	arr := []BoardNaverTablet{}
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		s.Find("selector")

	})
}
func connectDB() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	if err != nil {
		log.Fatal(err)
	}
	trans, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return trans, err
}
