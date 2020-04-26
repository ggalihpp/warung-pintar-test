package message

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"gitlab.com/ggalihpp/warung_pintar/repository"
)

var (
	upgrader = websocket.Upgrader{}
)

func sendMessage(c echo.Context) error {
	body := new(messageRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if body.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username cannot be empty")
	}

	if err := repository.SendMessage(body.Username, body.Message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	broadcastMessageV2(body.Username, body.Message)

	return c.JSON(http.StatusOK, "OK")
}

func getMessages(c echo.Context) error {
	msgs := repository.GetMessages()

	return c.JSON(http.StatusOK, msgs)
}

func wsHandler(c echo.Context) error {
	currentGorillaConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	username := c.Param("USERNAME")
	currentConn := webSocketConnection{Conn: currentGorillaConn, Username: username}
	connections[username] = &currentConn

	fmt.Printf("%s connected \n", username)

	go handleIO(&currentConn)
	return nil
}
