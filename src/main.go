package main

import(
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"TechDoJo/handler"

)

func main(){
	//http.HandleFunc("/get",handler.GetHandler)
//  http://localhost:8080/getAll
	http.HandleFunc("/getAll", handler.GetAllHandler)
	http.ListenAndServe(":8080", nil)

}