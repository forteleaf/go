package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("naver_news.go")
	if err != nil {
		return err
	}
	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		// 읽은 num 값 사용
	}
}
