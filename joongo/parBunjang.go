package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

type Bunjang struct {
	Num      string // 글번호
	Thumnail string
	Title    string
	Date     string
	Price    string
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

func InsertBunjang(data []Bunjang) {
	tranx, _ := connectDB()
	defer tranx.Rollback()
	// | sitenum   | varchar(100) | NO   | PRI | NULL    |       |
	// | thumbnail | varchar(255) | YES  |     | NULL    |       |
	// | category  | varchar(10)  | YES  |     | NULL    |       |
	// | title     | varchar(255) | YES  |     | NULL    |       |
	// | writer    | varchar(50)  | YES  |     | NULL    |       |
	// | price     | varchar(10)  | YES  |     | NULL    |       |
	// | date      | datetime     | YES  |     | NULL    |       |
	// | hit       | int(5)       | YES  |     | NULL    |       |
	stmt, err := tranx.Prepare("insert into joongo values (concat('bunj_',?),?,?,?,?,?,?,?)" +
		"on DUPLICATE KEY UPDATE price=?")
	checkError(err)
	defer stmt.Close()

	for k, _ := range data {
		_, err := stmt.Exec(data[k].Num, data[k].Thumnail, "판매", data[k].Title, "", data[k].Price, data[k].Date, "",
			data[k].Price)
		checkError(err)
	}
	err = tranx.Commit()
	checkError(err)
}
func bunjang() []Bunjang {
	arr := []Bunjang{}
	now := time.Now()

	doc, err := goquery.NewDocument("http://m.bunjang.co.kr/categories/600?order=date")
	if err != nil {
		log.Fatal(err)
	}
	now = now.Add((2 * time.Second) - (time.Minute))
	doc.Find(".goodslist li").Each(func(i int, s *goquery.Selection) {
		/// link = /products/59283082/?ref=%25ED%2599%2588&ad_ref=
		link, _ := s.Find("a").First().Attr("href")
		num := link[10:18]
		thumbnail, _ := s.Find("a .thumb").Attr("data-src") //사진
		title := s.Find(".txtinfo em").Text()               //제목
		price := s.Find("a p strong span").Text()           //가격

		// date := now.Format(time.RFC822)
		date := now.Format("2006-01-02 13:04:05") //시간
		// fmt.Println(link)
		// fmt.Println(thumbnail) fmt.Println("[" + title + "]")
		// fmt.Println("[" + price + "]")
		// fmt.Println(date)
		// fmt.Println(now.Format("2006-01-02 15:04:05"))
		arr = append(arr, Bunjang{num, thumbnail, title, date, price})
	})
	fmt.Println(arr[2:])
	return arr[2:]
}

func main() {
	arr := bunjang()
	InsertBunjang(arr)
}
