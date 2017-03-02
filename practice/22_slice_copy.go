package main

import "fmt"

// var s = make([]int, 5, 10)
// 슬라이스의 길이가 5,
// 용량이 10

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3)

	copy(b, a)

	fmt.Println(a)
	fmt.Println(b)

}
