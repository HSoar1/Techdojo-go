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
	fmt.Print("削除するユーザーのIDを入力してください\n")
	ID.Scan()
	result, err := db.Exec("DELETE FROM users WHERE id = " + ID.Text())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}