package server

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Ws        *websocket.Conn
	ReceiveCh chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	client := &Client{
		Ws:        ws,
		ReceiveCh: make(chan []byte, 256),
	}
	go client.readLoop()
	return client
}

func (client Client) SendText(text string) {
	fmt.Println("[SEND] " + text)
	w, _ := client.Ws.NextWriter(websocket.TextMessage)
	w.Write([]byte(text))
	if err := w.Close(); err != nil {
		return
	}
}

func (client Client) readLoop() {
	for {
		_, data, err := client.Ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("[RECEIVE] " + string(data))
		client.ReceiveCh <- data
	}
}
