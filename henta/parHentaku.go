package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Actor struct {
	name string
	link string
}

// get actor name
func parsingActors(pagenum string) {
	// https://hentaku.net/list.php?sort=&page=1
	doc, err := goquery.NewDocument("https://hentaku.net/list.php?sort=&page=" + pagenum)
	if err != nil {
		log.Fatal(err)
	}

	link, _ := doc.Find("div .list").Attr("href")
	link = "https://hentaku.net" + link
	name := doc.Find("div .list ")

}
func main() {

}
