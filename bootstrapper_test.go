package bddgo

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sendRequest(t *testing.T, handler http.Handler, request string) string {
	writer := httptest.NewRecorder()
	reader := strings.NewReader(request)
	buf := bufio.NewReader(reader)

	err := ServeSingleRequest(buf, writer, handler)
	if err != nil {
		t.Errorf("error while parsing http request : %q", err)
	}

	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func TestSimpleRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	})

	body := sendRequest(t, handler, "GET / HTTP/1.1\r\n\r\n")
	assert.Equal(t, body, "Hello")
}

func TestURLEncoded(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintf(w, "%s %s", r.FormValue("foo"), r.FormValue("baz"))
	})

	body := sendRequest(t, handler, `POST / HTTP/1.1
Host: localhost:8090
User-Agent: curl/7.58.0
Accept: */*
Content-Length: 15
Content-Type: application/x-www-form-urlencoded

foo=bar&baz=qux`)

	assert.Equal(t, "bar qux", body)
}

func TestMultipart(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here r is *http.Request
		parseErr := r.ParseMultipartForm(1024 * 10000) // maxMemory 32MB
		if parseErr != nil {
			http.Error(w, "failed to parse multipart message", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%s %s", r.FormValue("foo"), r.FormValue("baz"))
	})

	body := sendRequest(t, handler, `POST / HTTP/1.1
Host: example.org
User-Agent: curl/7.58.0
Accept: */*
Content-Length: 236
Content-Type: multipart/form-data; boundary=------------------------1b9c1df10c236cdc

--------------------------1b9c1df10c236cdc
Content-Disposition: form-data; name="foo"

bar
--------------------------1b9c1df10c236cdc
Content-Disposition: form-data; name="baz"

qux
--------------------------1b9c1df10c236cdc--
`)

	assert.Equal(t, "bar qux", body)
}
