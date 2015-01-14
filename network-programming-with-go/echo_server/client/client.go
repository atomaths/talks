package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:9009")
	if e != nil {
		panic(e.Error())
	}

	buf := make([]byte, 512)

	for {
		_, err := conn.Write([]byte("hi\n"))
		if err != nil {
			println(err.Error())
		}

		n, err := conn.Read(buf)
		fmt.Printf("%s", buf[0:n])

		time.Sleep(3 * 1e9)
	}
}
