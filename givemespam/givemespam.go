package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// GetWords is Level 2
func GetWords() []string {
	var words []string
	doc, _ := goquery.NewDocument("https://docs.google.com/spreadsheets/d/116Ajw6c7JGGYY0pYBFAYh_tooM5SLC7yAtb4yUIp40Y/gviz/tq?tqx=out:html&tq&grid=0")
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		band := s.Find("td").Text()
		words = append(words, band)
	})
	return words[1:] // 첫째 배열은 nil 값
}

// GetTitle is Level 1
func GetTitle() string {
	list := GetWords()
	result := make([]string, 3) // 3개의 배열
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		ss := rand.Intn(len(list) - 1)
		word := strings.Split(list[ss], "")
		result[i] = strings.Join(word, " ")
	}
	return strings.Join(result, " ")
}

// ReadLines read a This Source!!!
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// GetSource 는 소스를 첨가할 수 있지!
func GetSource() string {
	lines, err := ReadLines("givemespam.go")
	if err != nil {
		log.Printf("readlines: %s", err)
		os.Exit(0)
	}
	var sentence string
	for _, line := range lines {
		//fmt.Println(i, line)
		sentence += line + "\n"
	}
	return sentence
}

// SendMail is 메일보내기!!
func SendMail(title string, pwd string) {
	to := "spamworldcup@stibee.com" // 받는 사람
	auth := smtp.PlainAuth("", "forteleaf@gmail.com", pwd, "smtp.gmail.com")

	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		GetSource() +
		"스팸이 먹고 싶다!!!\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "forteleaf@gmail.com", []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Print("input Password >>> ")
	var pwd string
	fmt.Scanln(&pwd)

	title := GetTitle()
	SendMail(title, pwd)
}
