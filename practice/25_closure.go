package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println(a + b)
		return
	}(3, 5)

	sum := func(a, b int) int {
		return a + b
	}
	r := sum(1, 2)
	fmt.Println(r)

}
