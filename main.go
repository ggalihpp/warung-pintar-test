package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var loggerConfiguration middleware.LoggerConfig

func init() {
	///// LOAD ENV FILE
	//specify env file
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	// load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	loggerConfiguration = middleware.LoggerConfig{
		Format: `{"type":"backend","time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency":"${latency_human}"}` + "\n",
	}

}

func main() {
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(loggerConfiguration),
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	SetupHandlers(e)

	servicePort := os.Getenv("PORT")
	e.Start(servicePort)
}
