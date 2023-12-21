package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github/pikachu0310/unity-websocket-lecture/server"
	"time"
)

var Clients []server.Client

func HandleConnectionRequest(c echo.Context) error {
	upgrader := &websocket.Upgrader{}

	ws, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)
	client := server.NewClient(ws)
	Clients = append(Clients, *client)
	if len(Clients) == 2 {
		server.NewRoom(Clients)
		Clients = make([]server.Client, 0)
	}
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			i++
			//client.SendText("Hello From Server " + strconv.Itoa(i))
		}
	}()

	return nil
}

func main() {
	Clients = make([]server.Client, 0)

	e := echo.New()
	e.GET("/ws", HandleConnectionRequest)

	fmt.Println(e.Start(":1729"))
}
