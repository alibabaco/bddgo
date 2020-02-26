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

func TestPost(t *testing.T) {
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

/*
var requests []request
var responseWriter *httptest.ResponseRecorder
var handler func(w http.ResponseWriter, r *http.Request)
var mux *http.ServeMux

func setup() {
	requests = append(requests, request{, "TestCorrectRequest"})

	requests = append(requests, request{`POST /test HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 40
field1=value1&field2=value2`, "TestWrongHeaders"})

	requests = append(requests, request{`GET /test HTTP/1.1`, "TestWrongEOF"})

	requests = append(requests, request{`GET /api/v1 HTTP/1.1

POST /test HTTP/1.1
Host: foo.example
Content-Type: application/x-www-form-urlencoded
Content-Length: 27

field1=value1&field2=value2GET /api/v1/helloworld HTTP/1.2

`, "TestCorrectMultiRequests"})

	requests = append(requests, request{`GET /api/v1 HTTP/1.1

POST /test HTTP/1.1
Host: foo.example
Content-Type: application/x-www-form-urlencoded
Content-Length: 27

field1=value1&field2=value2GET /api/v1/helloworld HTTP/1.2
`, "TestWrongMultiRequests"})

	handler =
	mux = http.NewServeMux()
	mux.HandleFunc("/", handler)
}

func filterRequestsByName(requests []request, testName string) request {
	for _, req := range requests {
		if req.testName == testName {
			return req
		}
	}
	panic("didn't find a test with name : " + testName)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestWrongHeaders(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongHeaders")

	err := ParseRequest(strings.NewReader(httpRequest.text), responseWriter, mux)
	if err == nil {
		t.Errorf("expected `malformed MIME header line` but the request was parsed incorrectly!")
	}
}

func TestWrongEOF(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongEOF")
	err := ParseRequest(strings.NewReader(httpRequest.text), responseWriter, mux)
	if err == nil {
		t.Errorf("expected `unexpected EOF` but the request was parsed incorrectly!")
	}
}

func TestCorrectMultiRequests(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestCorrectMultiRequests")
	err := ParseRequest(strings.NewReader(httpRequest.text), responseWriter, mux)
	if err != nil {
		t.Errorf("error whlie parsing three consecutive http requets : %q", err)
	}
}

func TestWrongMultiRequests(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongMultiRequests")
	err := ParseRequest(strings.NewReader(httpRequest.text), responseWriter, mux)
	if err == nil {
		t.Errorf("expected `unexpected EOF` but the requets where parsed incorrectly")
	}
}

func TestResponseBody(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestCorrectRequest")
	responseWriter = httptest.NewRecorder()
	err := ParseRequest(strings.NewReader(httpRequest.text), responseWriter, mux)
	if err != nil {
		t.Errorf("error while parsing http request : %q", err)
	}

	body := responseWriter.Body.String()
	expectedBody := "<html><body>Hello World!</body></html>"
	if strings.TrimSpace(body) != expectedBody {
		t.Errorf("expected response body : %q but recieved %q", expectedBody, body)
	}
}
*/
