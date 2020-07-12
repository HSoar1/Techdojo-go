package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct{
	ID int
	Name string
	Token string
}

func main(){
	db, err := sql.Open("mysql", "root:rootpass@tcp(dojo_mysql_1:3306)/api_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var user user
		err := rows.Scan(&user.ID, &user.Name, &user.Token)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name, user.Token)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}