package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 10전 데이터 지우기

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() (*sql.Tx, error) {
	db, err := sql.Open("mysql", "sold:sold1234@tcp(305401.iptime.org:3306)/sold")
	checkError(err)
	trans, err := db.Begin()
	checkError(err)
	return trans, err
}

// mariaDb 에 insert
func main() {
	tranx, _ := connectDB()
	defer tranx.Rollback()

	db, err := tranx.Exec("delete from joongo where date < date(date_sub(now(), interval 10 day))")
	checkError(err)
	fmt.Println(db.RowsAffected())

	err = tranx.Commit()
	checkError(err)
}
