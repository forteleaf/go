package main

import "fmt"

func main() {
	a := [5]int{32, 29, 29, 33, 33}

	for _, value := range a {
		fmt.Println(value)
	}
}
