package server

type BoardUpdateSendMessage struct {
	MessageType string `json:"messageType"`
	Cells       []int  `json:"cells"`
}
