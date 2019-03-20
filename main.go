package main

import (
	"bufio"
	"fmt"
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	fmt.Println("This is not printed until the Client closes the connection at which point scanner.Scan returns false")
}
