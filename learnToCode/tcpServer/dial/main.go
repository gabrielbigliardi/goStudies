// package main

// import (
// 	"net"
// )

// func main() {
// 	// NOT WORKING
// 	var tcp = "tcp"
// 	con, err := net.Dial(tcp, "127.0.0.1:8070")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer con.Close()

// bs, err := io.ReadAll(con)
// if err != nil {
// 	log.Fatal("error on read", err)
// }
// fmt.Println(string(bs))
//

// }

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
