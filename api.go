package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/ggalihpp/warung_pintar/message"
)

func SetupHandlers(e *echo.Echo) {
	e.File("/", "public/index.html")

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	// Message
	msgGrp := e.Group("/messages")
	message.Handler(msgGrp)
}
