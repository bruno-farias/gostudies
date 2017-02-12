package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type Post struct {
    Id int
    Title string
    Body string
}



func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/index.html"))
	if error := template.ExecuteTemplate(w, "index.html", nil); error != nil {
	    http.Error(w, error.Error(), http.StatusInternalServerError)
	}
    })

    fmt.Println(http.ListenAndServe(":8080", nil))

}
