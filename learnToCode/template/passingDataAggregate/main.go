package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type person struct {
	name string
	age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

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

	// sliceOfPerson := []person{forrest, jenny}
	mapOfPerson := map[string]person{
		"forrest": forrest, "jenny": jenny,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "jenny.html", mapOfPerson)
	if err != nil {
		panic(err)
	}
}
