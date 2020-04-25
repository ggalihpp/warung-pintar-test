package message

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

var connections = map[string]*webSocketConnection{}

type socketPayload struct {
	Message string
}

type socketResponse struct {
	From    string
	Message string
}

type webSocketConnection struct {
	*websocket.Conn
	Username string
}

func handleIO(currentConn *webSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, fmt.Sprintf("%s has joined the chat...", currentConn.Username))

	for {
		payload := socketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, fmt.Sprintf("%s has left the chat...", currentConn.Username))
				removeConn(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, payload.Message)
	}
}

func removeConn(currentConn *webSocketConnection) {
	delete(connections, currentConn.Username)
}

func broadcastMessage(currentConn *webSocketConnection, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(socketResponse{
			From:    currentConn.Username,
			Message: message,
		})
	}
}

func broadcastMessageV2(username, message string) {
	for _, eachConn := range connections {
		eachConn.WriteJSON(socketResponse{
			From:    username,
			Message: message,
		})
	}
}
