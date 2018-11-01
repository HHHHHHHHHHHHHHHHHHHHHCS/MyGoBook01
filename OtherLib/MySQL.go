package OtherLib

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"math/rand"
	"time"
)

var (
	username = "root"
	password = "123123"
	address  = "127.0.0.1:3306"
	dbname   = "testdb01"
	options  = "charset=utf8"
)

func Main_MySQL() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, address, dbname, options)
	db, err := sql.Open("mysql", dataSourceName)
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("insert scores set name=?,score=?,other=?")
	checkErr(err)
	res, err := stmt.Exec(randomName(), r.Intn(100), time.Now())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("update scores set score=? where name=?")
	checkErr(err)
	res, err = stmt.Exec(r.Intn(100), "bob")
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("select  * from scores")
	checkErr(err)
	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	for k, v := range cols {
		values[k] = new(interface{})
		fmt.Println(k, v)
	}
	for rows.Next() {
		rows.Scan(values...)
		PrintRow(values)
	}

	stmt,err=db.Prepare("delete from scores where id=?")
	checkErr(err)
	res,err=stmt.Exec(id)
	checkErr(err)
	affect,err=res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()

}

func randomName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var name string
	for i := 0; i <= r.Intn(3)+3; i++ {
		name += string(r.Intn(26) + 97)
	}
	return name
}

func PrintRow(colsdata []interface{}) {
	for _, val := range colsdata {
		switch v := (*val.(*interface{})).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2016-01-02 15:05:05.999"))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
