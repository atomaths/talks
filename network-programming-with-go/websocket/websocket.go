package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

const MAX_CLIENT = 100

type Client struct {
	conn   *websocket.Conn
	userId string
}

var (
	chConn = make(chan Client, MAX_CLIENT)
	chMsg  = make(chan string, MAX_CLIENT)
)

func ChatServer(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	n, _ := ws.Read(buf)
	userId := string(buf[0:n])

	chMsg <- userId + "님이 들어오셨습니다."

	newClient := Client{ws, userId}
	chConn <- newClient

	for n, e := ws.Read(buf); e == nil; n, e = ws.Read(buf) {
		chMsg <- "<b>" + userId + "</b>: " + string(buf[0:n])
	}

	chMsg <- userId + "님이 나가셨습니다."
	ws.Close()
}

func WriteToClient() {
	clients := make([]Client, 0, MAX_CLIENT)

	for {
		select {
		case newClient := <-chConn:
			clients = append(clients, newClient)
		case msg := <-chMsg:
			fmt.Println(string(msg))
			for _, v := range clients {
				v.conn.Write([]byte(msg))
				// TODO: Write 실패 시 client 제거
			}
		}
	}
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	go WriteToClient()

	http.HandleFunc("/", IndexPage)
	http.Handle("/chat", websocket.Handler(ChatServer))
	if err := http.ListenAndServe(":9009", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
