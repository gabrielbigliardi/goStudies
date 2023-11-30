package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	Jumba     int    `json:"userId"`
	Dog       int    `json:"id"`
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
	// fmt.Printf("%s\n", response.Body)
	// fmt.Printf("%s\n", response.Header)

	// for key, value := range response.Header {
	// 	fmt.Printf("Key: %s, Value: %v\n", key, value)
	// }

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(bodyBytes) // ReadAll devolve um array de bytes

		data := string(bodyBytes) // string() converte o array da bytes em uma string
		fmt.Println(data)

		todoItems := Todo{}                   // tb eh possivel instanciar uma struct q possui o mesmo formato dos dados que era receber
		json.Unmarshal(bodyBytes, &todoItems) // e apontar para ela o conteudo convertido de bytes com o metodo json.Unmarshal()

		fmt.Println(todoItems)
		fmt.Println(todoItems.Dog)

		fmt.Printf("%+v\n", todoItems)

	}
}
