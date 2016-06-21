package SplitLineRe

import "fmt"

func ExampleSplitLineRe() {
	raw := "Hello, WOrld,\r\nWe are the world\nhello"
	for _, x := range SplitLineRe(raw) {
		fmt.Println(x)
	}
}
