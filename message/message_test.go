package message

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/messages")

	// Assertions
	if assert.NoError(t, getMessages(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSendMessageEmptyBody_then_BadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/messages")

	err := sendMessage(c)
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.Equal(t, "Username cannot be empty", he.Message)
		}
	}
}
