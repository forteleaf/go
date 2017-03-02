package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var myTime time.Time
	rows, err := db.Query("select current_timestamp()")

	if rows.Next() {
		if err = rows.Scan(&myTime); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%v", myTime)
}
