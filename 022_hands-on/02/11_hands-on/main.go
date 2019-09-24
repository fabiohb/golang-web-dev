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
	go serve(conn)
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
	}

	writeResponse(conn, "Você se conectou com sucesso!")
}

func writeResponse(conn net.Conn, body string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain; charset=utf-8\r\n")

	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
