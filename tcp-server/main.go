package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type requestHeaders map[string]string

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

	request := requestHeaders{}

	for scanner.Scan() {
		ln := scanner.Text()
		header := strings.Fields(ln)
		if len(header) > 0 {
			request[header[0]] = header[1]
		}

		if len(header) > 0 && header[0] == "Host:" {
			body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>%s</strong></body></html>`, fmt.Sprintf("URI : %s%s", header[1], request["GET"]))

			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
		}
	}
}
