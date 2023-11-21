package main

import (
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("error on listen ")
	}

	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			panic(err)
		}

		//handles only one connection
		io.Copy(con, con)
		con.Close()
	}
}
