package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	// io.WriteString(conn, fmt.Sprint("FROM Server: Hello World!"))

	scanner := bufio.NewScanner(conn)
	fmt.Printf("%T", scanner)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}

func main() {
	listner, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}

}
