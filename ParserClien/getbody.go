package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getBody(url string) (*bytes.Buffer, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(body)
	return buf, nil
}

func clien_title() {
	s, err := getBody("http://clien.net/cs2/bbs/board.php?bo_table=sold")
	if err != nil {
		fmt.Print(err.Error())
	}
	newFile, err := os.Create("clien.txt")
	if err != nil {
		log.Fatal(err)
	}
	// 파일저장하기
	newFile.Write(s.Bytes())
}

func main() {
	clien_title()
}
