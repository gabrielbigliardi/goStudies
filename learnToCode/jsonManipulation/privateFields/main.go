package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Completed bool `json:"completed"`
	// por estar com letra minuscula, o atributo eh privado
	itle   string
	Id     int `json:"id"`
	UserId int
}

// ESSE CODIGO RETORNARA UM ERRO, O EXERCICIO TRATA-SE DE TRATAR O ERRO
// O erro se da ao campo completed nao estar presente
// tratamento feito pelo metodo DisallowUnknownFields

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		todoItems := Todo{}
		decoder := json.NewDecoder(response.Body)

		err := decoder.Decode(&todoItems)
		if err != nil {
			log.Fatal("Decode error:", err)
		}

		fmt.Printf("%+v\n", todoItems)

	}
}
