package main

import "fmt"

// var 이름 []자료형
// var name []int

func main() {
	a := []int{1, 2, 3}
	var b []int

	// 슬라이스는 주소값을 참조만 한다
	b = a
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
