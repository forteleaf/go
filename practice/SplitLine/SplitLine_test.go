package SplitLine

import "fmt"

func ExampleSplitLine() {
	raw := "Hello, WOrld,\r\nWe are the world \nhello"
	for _, x := range SplitLine(raw) {
		fmt.Println(x)
	}

}
