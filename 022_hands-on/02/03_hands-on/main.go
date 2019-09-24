package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		accept(l)
	}
}

func accept(l net.Listener) {
	conn, err := l.Accept()
	if err != nil {
		log.Println(err)
	}

	// write to connection
	go process(conn)
}

func process(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	io.WriteString(conn, "Você se conectou com sucesso!")
}
