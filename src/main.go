package main

import(
/*	"fmt"
	"net/http"
	"log"
	"databases/sql"
	_ "github.com/go-sql-driver/mysql"
	"encording/json"

	"dojo/src/main/handler"
*/
//"fmt"
"net/http"
//"log"
//"database/sql"
_ "github.com/go-sql-driver/mysql"
//"encoding/json"

"TechDoJo/handler"
)

/*type user struct{
	ID int
	Name string
	Token string
}
*/

func main(){
	//http.HandleFunc("/get",handler.GetHandler)
//  http://localhost:8080/getAll
	http.HandleFunc("/getAll", handler.GetAllHandler)
	http.ListenAndServe(":8080", nil)

}