package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

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
		// primeiramente decodificamos o body da requisição
		err := decoder.Decode(&todoItems)
		if err != nil {
			log.Fatal("Decode error:", err)
		}

		// depois para codificar para json novamente \/
		todo, err := json.Marshal(todoItems)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T", todo)
		fmt.Println(todo)
		fmt.Println(string(todo))

	}
}
