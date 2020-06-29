package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:@/dojo_api_1")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil{
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values{
		scanArgs[i] = &values[i];
	}

	for rows.Next(){
		err = rows.Scan(ScanArgs...)
		if err != nil{
			panic(err.Error())
		}

		var value String
		for i, col := range values{
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("----------------------------------------")
	}
}