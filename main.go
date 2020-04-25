package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	SetupHandlers(e)

	servicePort := ":3000"
	e.Start(servicePort)
}
