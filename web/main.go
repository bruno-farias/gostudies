package main

import (
    "fmt"
    "net/http"
    "html/template"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

type Post struct {
    Id    int
    Title string
    Body  string
}

var db, err = sql.Open("mysql", "homestead:secret@(127.0.0.1:33060)/go?charset=utf8")

func main() {

    //stmt, err := db.Prepare("insert into posts (title,body) values (?,?);")
    //checkErr(err)
	//
    //_ , err = stmt.Exec("My first Post", "My first content")
    //checkErr(err)
	//
    //db.Close()


    rows, err := db.Query("SELECT * FROM posts;")
    items := []Post{}
    checkErr(err)

    for rows.Next()  {
	post := Post{}

	rows.Scan(&post.Id, &post.Title, &post.Body)
	items = append(items, post)
	fmt.Println(items)
    }

    db.Close()


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	post := Post{Id: 1, Title: "First Post", Body: "Some content"}

	if title := r.FormValue("title"); title != "" {
	    post.Title = title
	}

	template := template.Must(template.ParseFiles("templates/index.html"))
	if error := template.ExecuteTemplate(w, "index.html", items); error != nil {
	    http.Error(w, error.Error(), http.StatusInternalServerError)
	}
    })

    fmt.Println(http.ListenAndServe(":8080", nil))

}

func checkErr(err error) {
    if err != nil {
	panic(err)
    }
}