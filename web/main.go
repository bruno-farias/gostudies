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

	post := Post{Id: 1, Title: "First Post", Body: "Some content"}

	if title := r.FormValue("title"); title != "" {
	    post.Title = title
	}

	template := template.Must(template.ParseFiles("templates/index.html"))
	if error := template.ExecuteTemplate(w, "index.html", post); error != nil {
	    http.Error(w, error.Error(), http.StatusInternalServerError)
	}
    })

    fmt.Println(http.ListenAndServe(":8080", nil))

}
