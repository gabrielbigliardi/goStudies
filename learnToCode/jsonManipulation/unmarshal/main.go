package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todos struct {
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

	// for key, value := range response.Header {
	// 	fmt.Printf("Key: %s, Value: %v\n", key, value)
	// }

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// data := string(bodyBytes)
		// fmt.Println(data)

		todoItems := Todos{}
		json.Unmarshal(bodyBytes, &todoItems)

		fmt.Printf("%+v\n", todoItems)

	}
}
