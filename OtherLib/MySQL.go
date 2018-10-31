package OtherLib

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var (
	username = "username"
	password = "password"
	address  = "127.0.0.1:3306"
	dbname   = "testdb01"
	options  = "charset=utf8"
)

func Main_MySQL() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, address, dbname, options)
	db, err := sql.Open("mysql", dataSourceName)
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("INSERT  city SET name=?,score=?,other=?")
	res, err := stmt.Exec("bob", 0, "201X-X-X")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
