// 실행을 할때 옵션을 넣어서 실행하기

package main

import (
	"flag"
	"fmt"
)

type user struct {
	email    string
	password string
}

func main() {
	var email string
	var password string

	// 저장할곳, 옵션명, 값이 없을 때 default, 이건 뭐지

	flag.StringVar(&email, "email", "none password", "param value")
	flag.StringVar(&password, "password", "none password", "pwd value")
	flag.Parse()

	fmt.Println(email, password)
	// fmt.Printf("param = %s\n", param,)
}
