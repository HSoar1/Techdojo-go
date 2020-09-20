package handler

import (
	"database/sql"
	"fmt"
	"os"
	"bufio"

	_ "github.com/go-sql-driver/mysql"
)

/*type user struct{
	ID int
	Name string
	Token string
}
*/

func UpdateHandler(){
	db, err := sql.Open("mysql", "root:rootpass@tcp(dojo_mysql_1:3306)/api_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var ID *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var name *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("名前を変更するユーザーのIDを入力してください\n")
	ID.Scan()
	fmt.Print("変更するユーザー名前を入力してください\n")
	name.Scan()
	result, err := db.Exec("UPDATE users SET username = '" + name.Text() + "', token = '" + name.Text() + "stoken' WHERE id = " + ID.Text())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}