package helpers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessWithJSON(t *testing.T) {
	assert := assert.New(t)
	var SuccessWithJSONTests = []struct {
		message string
		code    int
	}{
		{"{\"doorId\": \"1\", \"price\": \"500$\"}", 200},
	}

	for _, test := range SuccessWithJSONTests {
		r := httptest.NewRecorder()
		SuccessWithJSON(r, test.message, test.code)
		assert.Equal("application/json; charset=utf-8", r.Header().Get("Content-Type"), "Content-Type header is not correct.")
		assert.Equal(200, r.Code, "HTTP code is not correct.")
		assert.JSONEq("{\"doorId\": \"1\", \"price\": \"500$\"}", r.Body.String(), "Answer body is not correct.")
	}
}

func TestErrorWithText(t *testing.T) {
	assert := assert.New(t)
	var ErrorWithTextTests = []struct {
		err     error
		message string
		code    int
	}{
		{errors.New("Some error message"), "Bad request", http.StatusBadRequest},
	}

	for _, test := range ErrorWithTextTests {
		r := httptest.NewRecorder()
		ErrorWithText(r, test.err, test.message, test.code)
		assert.Equal("text/plain", r.Header().Get("Content-Type"), "Content-Type header is not correct.")
		assert.Equal(http.StatusBadRequest, r.Code, "HTTP code is not correct.")
		assert.Equal("Bad request", r.Body.String(), "Answer body is not correct.")
	}
}
