package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage_then_ReturnNull(t *testing.T) {
	msgs := GetMessages()

	assert.Equal(t, []Message(nil), msgs)
}

func TestSendMessage_then_MessagesIsFill(t *testing.T) {
	SendMessage("test", "test message")

	msgs := GetMessages()

	assert.Equal(t, 1, len(msgs))
}
