package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type List struct {
	Link string // 링크번호
	Pic  string // 사진링크
}

// 배우정보
type Profile struct {
	NameKOR string // 이름 한글
	NameENG string // 이름 영문
	NameJPN string // 이름 일어
	Birth   string // 생일
	Height  int    // 신장
	BWH     string // 신체 사이즈
	Cup     string // 컵 사이즈
}

// 작품
type Product struct {
	Pum     string    // 품번
	Title   string    // 제목
	Release time.Time // 발매일
	Tag     string    // 태그
	Actor   string    // 출연진
}

func TotalPage() {

}

// Total 페이지를 받아와서 페이지를 하나씩 열면서 배우 링크정보와 이미지 주소를 가져옴
func ActorsList(TotalPage int) []List {
	// 1~마지막 페이지까지 스캔
	// closure ParList
	ParList := func(listPageNum int) []List {
		slist := []List{}

		doc, err := goquery.NewDocument("https://hentaku.net/list.php?sort=&page=" + strconv.Itoa(listPageNum))
		if err != nil {
			log.Fatal(err)
		}
		// 총 페이지 갯수
		// page := doc.Find("div.page a li").Last().Text()

		doc.Find("div.list").Each(func(i int, s *goquery.Selection) {
			// link, img 두개만 가져오기
			link, _ := s.Find("a").Attr("href")
			link = "https://hentaku.net" + link
			img, _ := s.Find("img").Attr("src")
			slist = append(slist, List{link, img})
		})
		return slist
	}

	fmt.Println("배우목록 가져오기를 시작합니다. ")
	t0 := time.Now()
	totallist := []List{} // totallist : 모든 리스트를 저장하는 곳
	for i := 1; i <= TotalPage; i++ {
		list := ParList(i)
		for j := 0; j < len(list); j++ {
			totallist = append(totallist, List{list[j].Link, list[j].Pic})
		}
		// fmt.Printf("리스트의 갯수 %v", len(list))
	}
	t1 := time.Now()
	fmt.Printf("걸린시간은 %.3f 시간입니다.", t1.Sub(t0).Seconds())
	fmt.Println("모든 배우들", totallist)
	return totallist
}

func main() {
	ActorsList(27)
}
