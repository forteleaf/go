package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	test_01 = "http://www.iheadlinenews.co.kr/news/photo/201405/1881_2099_1848.jpg"
	test_02 = "https://oxfordportal.blob.core.windows.net/emotion/recognition1.jpg"
)

func main() {
	client := &http.Client{}

	// form := url.Values{}
	// form.Add("url", test_01)
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, test_02))
	byteData_NewBuffer := bytes.NewBuffer(byteData)

	req, err := http.NewRequest(
		"POST",
		"https://api.projectoxford.ai/emotion/v1.0/recognize",
		// strings.NewReader(form.Encode()),
		byteData_NewBuffer,
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", "7480b7f07d6e4669b1fde35f69a8036e")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(body))
}
