package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	// Completed bool   `json:"completed"`
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

	fmt.Printf("%T\n", response)
	fmt.Printf("%s\n", response.Status)
	fmt.Printf("%T\n", response.Body)

	if response.StatusCode == http.StatusOK {

		todoItems := Todo{}
		decoder := json.NewDecoder(response.Body)
		// esse metodo permite dizer para aplicação que nao quer aceitar uma resposta
		// q contem dado q vc nao espera
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&todoItems)
		if err != nil {
			log.Fatal("Decode error:", err)
		}

		fmt.Printf("%+v\n", todoItems)

	}
}
