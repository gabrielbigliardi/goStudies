package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

type person struct {
	name string
	age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func sliceToString(slice []string) string {
	// Use strings.Join para concatenar os elementos do slice
	result := strings.Join(slice, ", ")
	return result
}

func main() {
	forrest := person{
		"Forrest Gump",
		35,
	}

	jenny := person{
		"Jenny Curran",
		34,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "forest.html", forrest)
	if err != nil {
		panic(err)
	}

	mapOfPerson := map[string]person{
		"forrest": forrest, "jenny": jenny,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "jenny.html", mapOfPerson)
	if err != nil {
		panic(err)
	}

	// passing a slice to a file

	sliceOfPerson := []string{"teste1", "teste2"}

	file, err := os.Create(sliceToString(sliceOfPerson))
	if err != nil {
		log.Fatal("error on creating file", err)
	}

	_, err = file.WriteString("teste")
	if err != nil {
		log.Fatal("error on writing on file", err)
	}
}
