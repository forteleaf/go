// coolenjoy 장터 XML 파일 불러오기

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	Coolenjoy = "http://www.coolenjoy.net/rss?bo_table=mart2"
)

type ListItem struct {
	Items []Item `xml:"channel>item"`
}
type Item struct {
	// XMLName xml.Name `xml:"rss"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Context string `xml:"description"`
	Writer  string `xml:"creator"`
	Date    string `xml:"date"`
}

func getBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	return body, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func coolenjoySoldList() {
	s, err := getBody(Coolenjoy)
	checkError(err)
	var list ListItem
	err = xml.Unmarshal(s, &list)
	checkError(err)

	for f, v := range list.Items {
		fmt.Print(f, v.Context, "\n")
	}
}

func main() {
	coolenjoySoldList()

}
