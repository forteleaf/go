package main

import (
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
<head>
<title>Hello World</title>
</head>
<body>
Hello World!
</body>
</html>`,
	)
}
func main() {
	http.HandleFunc("/Hello", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
