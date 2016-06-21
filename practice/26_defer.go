package main

import "fmt"

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("hello")
	}()
}

func main() {
	HelloWorld()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 여기서 알아야 하는 것은
// LIFO 랑 실행되는 순서가 같다.
