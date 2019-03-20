package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9080")

	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "TCP Client connection")

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(bs))

}
