package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) { // io.ReadWriter 가능
	fmt.Println("Connected from:", conn.RemoteAddr().String())
	defer conn.Close()

	buf := make([]byte, 512)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed:", conn.RemoteAddr().String())
			}
			return
		}

		conn.Write(buf[0:n])
		fmt.Printf("%s", buf[0:n])
	}
}

func startServer() {
	listener, e := net.Listen("tcp", ":9009") // *net.TCPListener
	if e != nil {
		log.Fatalln(e)
	}
	defer listener.Close()

	for {
		conn, e := listener.Accept() // *net.TCPConn
		if e != nil {
			log.Println("Aceept error:", e)
			continue
		}

		echo(conn)
		// go echo(conn)
	}
}

func main() {
	startServer()
}
