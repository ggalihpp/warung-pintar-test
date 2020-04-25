package message

type messageRequestBody struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}
