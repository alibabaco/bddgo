package bddgo

import (
	"strings"
	"testing"
)

func TestCorrectRequest(t *testing.T) {
	httpRequestText := `POST /test HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 27

field1=value1&field2=value2`
	err := ParseRequest(strings.NewReader(httpRequestText))

	if err != nil {
		t.Errorf("error while parsing http request : %q", err)
	}
}

func TestWrongHeaders(t *testing.T) {
	httpRequestText := `POST /test HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 40
field1=value1&field2=value2`
	err := ParseRequest(strings.NewReader(httpRequestText))

	if err == nil {
		t.Errorf("expected `malformed MIME header line` but the request was parsed incorrectly!")
	}
}

func TestWrongEOF(t *testing.T) {
	httpRequestText := `GET /test HTTP/1.1`
	err := ParseRequest(strings.NewReader(httpRequestText))

	if err == nil {
		t.Errorf("expected `unexpected EOF` but the request was parsed incorrectly!")
	}
}
