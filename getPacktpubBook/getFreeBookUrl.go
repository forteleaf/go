package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

var cookies []*http.Cookie

//url 가져오기
func getfreeBookURL(body string) string {
	pattern := regexp.MustCompile(`/freelearning-claim/\d+/\d+`)
	return pattern.FindString(body)
}

// body 가져오기
func getBody() (*bytes.Buffer, error) {
	resp, err := http.Get(packtPubBaseURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(body)
	return buf, nil
}

func main() {
	body, _ := getBody()
	freeBookURL := getfreeBookURL(body)
	fmt.Println(freeBookURL)
}
