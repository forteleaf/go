package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	email   = "forteleaf@hotmail.com" // USEID is id
	api_id  = "2280790"               // APIID is id
	api_pwd = "888XATTM"              // APIPWD is password
)

type arg struct {
	BLOGID   string
	USER     string
	PASSWORD string
	Content  map[string]string
	ok       bool
}

type Content struct {
	title       string
	description string
}
e
func main() {
	client, err := rpc.Dial("tcp", "https://newsleaf.tistory.com/api")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var reply string
	// cont := Content{
	// 	"제목",
	// 	"내용",
	// 	"키워드",
	// 	{"카테그리", "고리고리"},
	// }
	// cont.categories = {"ddd", "테그테그"}
	// cont.title = "제목"
	// cont.mt_keywords = "키워드"
	// cont.description = "내용이다요"
	err = client.Call("metaWeblog.newPost", &arg{
		"forteleaf@gmail.com", "2280790", "888XATTM",
		{"title": "XML-RPC를 이용한 글 올리기 테스트", "description": "글에 들어갈 내용입니다."},
		true,
	}, &reply)

	if err != nil {
		log.Fatalf("Something went wrong: %v", err.Error())
	}
	fmt.Printf("The 'reply' pointer value has been changed to: %s", reply)

}
