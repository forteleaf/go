package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c := time.Tick(10 * time.Second)
	// 바로 시작 하지 않고, 10초 뒤에 시작한다.
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
}
