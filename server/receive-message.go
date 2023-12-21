package server

type CellClickReceiveMessage struct {
	MessageType string `json:"messageType"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
}
