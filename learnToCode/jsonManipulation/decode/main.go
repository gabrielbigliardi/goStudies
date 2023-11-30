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

	fmt.Printf("%T\n", response)
	fmt.Printf("%s\n", response.Status)
	fmt.Printf("%T\n", response.Body)

	if response.StatusCode == http.StatusOK {

		todoItems := Todo{}
		json.NewDecoder(response.Body).Decode(&todoItems)
		// a vantagem de usar NewDecoder é que tem alguns metodos adicionais em relacao ao Unmarshal

		fmt.Printf("%+v\n", todoItems)

	}
}
