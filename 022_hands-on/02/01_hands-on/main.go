package main

import (
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
	io.WriteString(conn, "Você se conectoru com sucesso!")
}
