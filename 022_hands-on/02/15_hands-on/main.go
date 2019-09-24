package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	var i int
	var method, uri string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			method = xs[0]
			uri = xs[1]
			fmt.Println("METHOD:", method)
			fmt.Println("URI:", uri)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Teste</title>
		</head>
		<body>
			<h1>"Você se conectou com sucesso!"</h1>
		</body>
		</html>
	`
	writeResponse(conn, body)
}

func writeResponse(conn net.Conn, body string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html; charset=utf-8\r\n")

	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
