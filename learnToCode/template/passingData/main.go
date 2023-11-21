package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))

}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", 42)
	if err != nil {
		log.Fatalln("erro na saida", err)
	}
}
