package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func slurp(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	return string(body), nil
}

// return al specfic groups
func re_groups(re *regexp.Regexp, text string, group int) []string {
	result := []string{}
	found := re.FindAllStringSubmatch(text, -1)
	for _, v := range found {
		result = append(result, v[group])
	}
	return result
}

func list_naver() []string {
	s, err := slurp("http://www.naver.com")
	if err != nil {
		return nil
	}
	return re_groups(regexp.MustCompile("<option value=\".+\">.+: (.+)</option>"),
		s,
		1)
}

func removeDuplicate(a []string) []string {
	result := []string{}
	found := make(map[string]bool)
	for _, v := range a {
		if !found[v] {
			found[v] = true
			result = append(result, v)
		}
	}
	return result
}

func list_daum() []string {
	s, err := slurp("http:www.daum.net")
	if err != nil {
		return nil
	}
	return removeDuplicate(re_groups(regexp.MustCompile("<span class=\"txt_issue\">\n.+\n(<.+>)?(.+?)(<.+>)?\n"),
		s,
		2))
}

func main() {
	const interval = 1
	fmt.Println("Refreshes every", interval, "minutes.")
	for {
		now := time.Now()
		fmt.Println(now)
		fmt.Println("네이버 :", strings.Join(list_naver(), ","))
		fmt.Println("다음 :", strings.Join(list_daum(), ", "))
		time.Sleep(interval * time.Minute)
	}
}
