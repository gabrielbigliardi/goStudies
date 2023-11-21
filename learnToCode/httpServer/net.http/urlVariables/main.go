package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		r.Method,
		r.URL,
		r.Form,
	}
	for v, i := range data.Submissions {
		fmt.Println("valor:", v)
		fmt.Println("index:", i)
	}
	fmt.Println("\n", data.Method)
	tpl.ExecuteTemplate(w, "index.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
