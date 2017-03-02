package main

import "fmt"

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) (z int) {
		z = a*x + b
		return
	}
}
func main() {
	f := calc()
	fmt.Println(f(1))
	// result
	// 8
}
