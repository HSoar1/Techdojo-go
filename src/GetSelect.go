package main

import (
	"database/sql"
	"fmt"
	"os"
	"bufio"

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
	var ID *bufio.Scanner = bufio.NewScanner(os.Stdin) 
	fmt.Print("検索するIDを入力してください\n")
	ID.Scan()
	rows, err := db.Query("SELECT * FROM users WHERE id = " + ID.Text())
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