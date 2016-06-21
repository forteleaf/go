package main

import (
	"fmt"
	"reflect"
)

func Pointer1() {
	num := 1
	var numPtr *int = &num
	fmt.Println(&numPtr, "type", reflect.TypeOf(numPtr))
	fmt.Println(&num)
}
func hello(n *int) {
	*n = 2
}
func main() {
	var n int = 1
	hello(&n)
	fmt.Println(n)
	fmt.Println(&n)
}
