package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	ClienSold = "http://www.clien.net/cs2/bbs/board.php?bo_table=sold"
)

func getBody(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), nil
}

func title_group(re *regexp.Regexp, body string, group int) []string {
	result := []string{}
	found := re.FindAllStringSubmatch(text, -1)
	for _, v := range found {
		result = append(result, v[group])
	}
	return result
}

func list_ClienSold() {
	body, err := getBody(ClienSold)
	if err != nil {
		log.Fatal(err)
	}
	title_group(regexp.MustCompile(mytr), body, 1)

}

func main() {
	list_ClienSold()
}
