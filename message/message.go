package message

import "github.com/labstack/echo"

// Handler - Handler for message
func Handler(e *echo.Group) {
	e.POST("/", sendMessage)
	e.GET("/", getMessages)
	e.GET("/ws/:USERNAME", wsHandler)
}
