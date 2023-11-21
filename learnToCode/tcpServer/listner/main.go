package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("error on listen ", err)
	}

	defer lis.Close()

	for {
		con, err := lis.Accept()
		if err != nil {
			log.Fatal("error on accept", err)
		}
		defer con.Close()
		io.WriteString(con, fmt.Sprint("FROM Server: Hello World!"))
		fmt.Fprintln(con, "How is your day?")
		fmt.Fprintf(con, "%v", "How is your day?")

	}

}
