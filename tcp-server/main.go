package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()

	if err != nil {
		log.Println("Connection timeout not set")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n", ln)
	}

	fmt.Println("This is printed after the connection is closed due to the deadline")
}
