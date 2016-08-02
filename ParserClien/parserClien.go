package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

type ClienSoldBoard struct {
	link     string
	category string
	title    string
	user     string
	reply    int
}

// clientSoldList get link, category, title, user, reply num
func clienSoldList() []ClienSoldBoard {
	arr := []ClienSoldBoard{}
	doc, err := goquery.NewDocument("http://m.clien.net/cs3/board?bo_style=lists&bo_table=sold")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find(".scalable").Attr("onclick")
		// slice bounds out of range
		//link = link[15 : len(link)-1]
		category := s.Find(".lst_category").Text()
		title := s.Find(".lst_tit").Text()
		user := s.Find(".lst_user a").Text()
		reply, _ := strconv.Atoi(s.Find(".lst_reply").Text())
		// fmt.Println(link, category, title, user, reply)
		// struct 에 집어 넣기
		arr = append(arr, ClienSoldBoard{link, category, title, user, reply})
	})
	arr = arr[2:]
	fmt.Println(arr)
	return arr
}

func mariaDB() {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// prepared statement
	stmt, err := db.Prepare("insert into ")

}

func main() {
	_ = clienSoldList()
}
