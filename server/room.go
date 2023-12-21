package server

import (
	"encoding/json"
	"fmt"
)

const (
	X    int = -1
	None int = 0
	O    int = 1
)

type Room struct {
	Clients []Client
	Board   []int
}

func NewRoom(clients []Client) Room {
	fmt.Println("New Room")
	room := Room{
		Clients: clients,
		Board:   make([]int, 9),
	}
	go room.ReceiveMessageLoop(clients[0], -1)
	go room.ReceiveMessageLoop(clients[1], 1)
	return room
}

func (room *Room) ReceiveMessageLoop(client Client, index int) {
	for {
		data := <-client.ReceiveCh
		fmt.Printf("[Room.Message] <%d> %s\n", index, string(data))

		var message CellClickReceiveMessage
		if json.Unmarshal(data, &message) == nil && message.MessageType == "CellClick" {
			fmt.Printf("X = %d, Y = %d\n", message.X, message.Y)
			if room.Board[message.Y*3+message.X] == None {
				room.Board[message.Y*3+message.X] = index
				room.SendBoardUpdateMessage()
			}
		}
	}
}

func (room *Room) SendBoardUpdateMessage() {
	sendMessage := BoardUpdateSendMessage{
		MessageType: "BoardUpdate",
		Cells:       room.Board,
	}
	jsonData, _ := json.Marshal(sendMessage)
	for _, client := range room.Clients {
		client.SendText(string(jsonData))
	}
}
