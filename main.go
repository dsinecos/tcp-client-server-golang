package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Panicln(err)
	}

	// Close the listening socket after the function returns
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicf("Error in connection %v", err)
		}

		// Write to TCP Connection socket
		io.WriteString(conn, "\n Hello from TCP Server")
		fmt.Fprintln(conn, "How is your day")
		fmt.Fprintf(conn, "%v", "Well, I hope \n")

		conn.Close()
	}
}
