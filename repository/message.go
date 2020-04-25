package repository

import "time"

// MessageStore - to store all the message data
var MessageStore []Message

//// HELPER ////

// SendMessage - represent sending message function
func SendMessage(username, msg string) error {
	M := Message{
		Username: username,
		Message:  msg,
	}

	return M.store()
}

// GetMessages - Get all stored messages
func GetMessages() []Message {
	return MessageStore
}

//////////////////

// Message - Model represent message data
type Message struct {
	Message  string    `json:"message"`
	Username string    `json:"username"`
	SendDate time.Time `json:"send_date"`
}

func (m *Message) store() error {
	m.SendDate = time.Now()

	MessageStore = append([]Message{*m}, MessageStore...)

	return nil
}
