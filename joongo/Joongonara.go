package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

/*
검색하기
http://m.cafe.naver.com/ArticleSearchList.nhn?search.query=키보드+적축
&search.menuid=&search.searchBy=0&search.sortBy=date&search.clubid=10050146

글 번호 뽑아오기

*/

// 홈페이지 읽어오기
func slurp(url string) (string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func re_groups(re *regexp.Regexp, text string, group int) []string {
	result := []string{}
	found := re.FindAllStringSubmatch(text, -1)
	for _, v := range found {
		result = append(result, v[group])
	}
	return result
}

// 검색어 : 키보드 적축
func list_naver() []string {
	s, err := slurp("http://m.cafe.naver.com/ArticleSearchList.nhn?search.query=키보드+적축&search.menuid=&search.searchBy=0&search.sortBy=date&search.clubid=10050146")
	fmt.Printf("%d", s)
	if err != nil {
		return nil
	}
	return re_groups(regexp.MustCompile("<ul "),
	return re_groups(regexp.MustCompile("<option value=\".+\">.+: (.+)</option>"),
		s,
		1)
}

func main() {
  // list_naver -> slurp ([]string )-> 
	fmt.Println("Naver:", strings.Join(list_naver(), ", "))
}
