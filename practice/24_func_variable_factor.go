package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for key, value := range n {
		total += value
		fmt.Println(key, value)
	}
	return total
}

func main() {
	r := sum(1, 2, 3, 4, 5, 4, 4, 5, 5, 2)
	fmt.Println(r)
}
