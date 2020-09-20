package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"

)



func GetAllHandler(w http.ResponseWriter, r *http.Request){

	type User struct{
		ID int `json:"id"`
		Name string `json:"name"`
		Token string `json:"token"`
	}
	
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
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Token)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name, user.Token)
		fmt.Fprintln(w ,user.ID, user.Name, user.Token)
		
	outputJson, err := json.Marshal(&user)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w,string(outputJson))
	
	}

	

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}