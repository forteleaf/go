package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// fmt.Println(now.String())

	hh, mm, ss := now.Clock()

	fmt.Println(hh, mm, ss)
	fmt.Println(now.String())
	fmt.Println(now.Format(time.RFC1123))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	for i := 0; i < 10; i++ {
		now = now.Add(1 * time.Second)
		fmt.Println(now)
	}
}
