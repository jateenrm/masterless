package main

import "net/http"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Quote struct {
quote string 
category string
}

var db *sql.DB
var err error

func InsertQuoteEndpoint(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
        		http.ServeFile(res, req, "insert.html")
       		return
	}

	quote := req.FormValue("quote")
	category := req.FormValue("category")

	var test string

    err := db.QueryRow("SELECT quote FROM quotation WHERE quote=?", quote).Scan(&test)

switch {
case err == sql.ErrNoRows:	
	_, err = db.Exec("INSERT INTO quotation(quote, category) VALUES(?, ?)", quote, category)
        if err != nil {
            http.Error(res, "Unable to Add Quote to Quotation Table.", 500)    
            return
        }

        res.Write([]byte("Quote Added Successfully!"))
        return
case err != nil: 
        http.Error(res, "Server error, unable to process you request.", 500)    
        return
default: 
        http.Redirect(res, req, "/", 301)
}
}
func FetchQuoteEndpoint(res http.ResponseWriter, req *http.Request) {
var test2 string
q := `
SELECT quote FROM quotation;
ORDER BY RAND()
LIMIT 1
`
err := db.Query((q), quote).Scan(&test2)
switch {
case err == sql.ErrNoRows:
	res.Write([]byte("Quotation:" + test2))
        return
case err != nil:
        http.Error(res, "Server error, unable to process you request.", 500)
        return
default:
        http.Redirect(res, req, "/", 301)
}
}
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}


func main() {
	db, err = sql.Open("mysql", "y2jat:RN6AX7ZG8@tcp(dbjateen.c8h9zfmu5hnn.ap-south-1.rds.amazonaws.com:3306)/mydb?charset=utf8")
		if err != nil {
		panic(err.Error())
		}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
http.HandleFunc("/new", InsertQuoteEndpoint)
http.HandleFunc("/quote", FetchQuoteEndpoint)
http.HandleFunc("/", homePage)
http.ListenAndServe(":8080", nil)
}
