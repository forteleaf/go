package main

import "fmt"

func main() {
	for _, r := range "가갛힛" {
		fmt.Println(string(r), r)
	}
	fmt.Println(len("가나다"))
}
