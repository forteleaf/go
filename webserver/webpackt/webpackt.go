package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	title string
	Books []Mybook
}

type Mybook struct {
	Num       int
	Thumbnail string
	Name      string
	Pdf       string
	Epub      string
	Mobi      string
	Codefiles string
}

// getMybook is download my packtpub list
func getMybook() (*Data, error) {
	db, err := sql.Open("mysql", "packt:packt1234@tcp(305401.iptime.org:3306)/packtpub")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from mybook order by num asc")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// select 데이터를 받아 오는 방법
	var list Data
	var item Mybook
	for rows.Next() {
		err = rows.Scan(&item.Num, &item.Thumbnail, &item.Name, &item.Pdf, &item.Epub, &item.Mobi, &item.Codefiles)
		if err != nil {
			panic(err)
		}
		list.Books = append(list.Books, Mybook{item.Num, item.Thumbnail, item.Name, item.Pdf, item.Epub, item.Mobi, item.Codefiles})
	}

	return &list, err
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	list, err := getMybook()
	if err != nil {
		panic(err)
	}
	t, err := template.ParseFiles("list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// list 넘기기
	err = t.Execute(w, list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe(":8080", nil)
}

// mariadb 를 통해서 db 값을 읽어오기
// 이를 struct 에 저장하기
