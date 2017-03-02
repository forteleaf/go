package main

import (
	"log"
	"net/rpc"
	"time"
)

const (
	APIADDR = "https://newsleaf.tistory.com/api" // APIADDR is url
	USERID  = "forteleaf@hotmail.com"            // USEID is id
	APIID   = "2280790"                          // APIID is id
	APIPWD  = "888XATTM"                         // APIPWD is password
)

type TimeArgs struct {
}

type TimeReply struct {

	Time time.Time
}
type Content struct {
	categories  []string
	title       string
	description string
	mt_keywords string
}

func main() {
	var content Content
	content.categories = "카테고리"
	content.description = "내용이 들어가는 곳"
	content.title = "제목"
	content.mt_keywords = "테그,보이는,곳"

	client, err := rpc.Dial("tcp", APIADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	client.Call("metaWeblog.newPost", args, reply)
}
