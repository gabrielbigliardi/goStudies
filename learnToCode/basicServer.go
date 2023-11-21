package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func aurora(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Rorica Maria")
}

func main() {
	http.HandleFunc("/dog", aurora)

	//handling error
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
	if err != nil {
		log.Fatal("Server error", err)
	}
}
