package main

import "fmt"

func main() {
	fmt.Println(divide(1, 0))
}
func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return a / b
}
