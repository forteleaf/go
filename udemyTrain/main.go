package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	fmt.Println("haha")
	res, err := http.Get("http://clien.net")
	if err != nil {
		log.Fatal(err.Error())
	}
	body, _:= ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
