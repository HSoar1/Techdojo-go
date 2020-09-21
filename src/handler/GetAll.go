package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"bytes"
)



func GetAllHandler(w http.ResponseWriter, r *http.Request){

	type User struct{
		ID int `json:"id"`
		Name string `json:"name"`
		Token string `json:"token"`
	}
	
	var output []User

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

	output = append(output, user)
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		panic(err.Error())
	}

	var result bytes.Buffer
	json.Indent(&result,outputJson,"","  ")
	fmt.Fprintln(w,result.String())	

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}