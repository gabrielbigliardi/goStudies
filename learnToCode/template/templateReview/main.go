package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatal("error", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "about.html", nil)
	if err != nil {
		log.Fatal("error", err)
	}

	//creating a file
	f1, err := os.Create("createMain.html")
	if err != nil {
		log.Fatal("another error")
	}
	defer f1.Close()

	f2, err := os.Create("createAbout.html")
	if err != nil {
		log.Fatal("another error")
	}
	defer f2.Close()

	err = tpl.ExecuteTemplate(f1, "index.html", nil)
	if err != nil {
		log.Fatal("error", err)
	}

	err = tpl.ExecuteTemplate(f2, "about.html", nil)
	if err != nil {
		log.Fatal("error", err)
	}
}
