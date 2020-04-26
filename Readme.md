# Guidelines

## Prerequisites

Golang version 1.11.x or later that support _GO MODULES_
check the golang version with `go version` command

## Installation

note: This installation guide assumed your os is unix based.

clone the project

```
git clone https://github.com/ggalihpp/warung-pintar-test.git
```

go to project folder

```
cd warung-pintar-test
```

build the project with GO MODULE enabled

```
GO111MODULE=on go build .
```

_warung-pintar-test_ binary will created run with

```
./warung-pintar-test
```

## Development

To run the project directly use `run` command (If run for the first time, activate go module)

```
go run main.go
```

## ROUTES

| Endpoint               | Method | Description                                                        | Parameter                                        | cURL Example                                                                                                                                                                                                                                                                                                              |
| ---------------------- | ------ | ------------------------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /messages/             | GET    | Get all stored messages                                            |                                                  | curl --request GET \ --url http://localhost:3000/messages/                                                                                                                                                                                                                                                                |
| /messages/             | POST   | Send (store) messages                                              | JSON Body: {"username":STRING, "message":STRING} | curl --request POST \ --url http://localhost:3000/messages/ \ --header 'content-type: application/json' \ --data '{ "message":"EHEHEHE", "username":"ggalihpp" }'                                                                                                                                                         |
| /messages/ws/:USERNAME | GET    | Connecting to message websocket and listen to any incoming message | Params -> USERNAME                               | curl --include \ --no-buffer \ --header " Connection: Upgrade" \ --header " Upgrade: websocket" \ --header " Host: example.com:80" \ --header " Origin: http://example.com:80" \ --header " Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \ --header " Sec-WebSocket-Version: 13" \ http://localhost:3000/messages/ws/ggalihpp |
| /                      | GET    | Chatting homepage                                                  |                                                  |                                                                                                                                                                                                                                                                                                                           |
