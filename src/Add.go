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
	var name *bufio.Scanner = bufio.NewScanner(os.Stdin) 
	fmt.Print("追加するIDを入力してください\n")
	ID.Scan()
	fmt.Print("追加する名前を入力してください\n")
	name.Scan()
	result, err := db.Exec("INSERT IGNORE INTO users (id, username, token) VALUES ( " + ID.Text() + ", '" + name.Text() + "', '" + name.Text() + "stoken')")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}