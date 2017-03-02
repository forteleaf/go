/*
	yotube -> face rectangles -> send pic via MS API -> response emotion score
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func test_listHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by youtself
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	list := make(map[string]string)
	list["name"] = "hwang"
	list["age"] = "18"

	t, err := template.ParseFiles("template/main.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/", listHandler)        // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
