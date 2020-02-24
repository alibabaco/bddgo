package bddgo

import (
	"os"
	"strings"
	"testing"
)

type request struct {
	text     string
	testName string
}

var requests []request

func setup() {
	requests = append(requests, request{`POST /test HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 27

field1=value1&field2=value2`, "TestCorrectRequest"})

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

func TestCorrectRequest(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestCorrectRequest")
	err := ParseRequest(strings.NewReader(httpRequest.text))

	if err != nil {
		t.Errorf("error while parsing http request : %q", err)
	}
}

func TestWrongHeaders(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongHeaders")
	err := ParseRequest(strings.NewReader(httpRequest.text))

	if err == nil {
		t.Errorf("expected `malformed MIME header line` but the request was parsed incorrectly!")
	}
}

func TestWrongEOF(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongEOF")
	err := ParseRequest(strings.NewReader(httpRequest.text))

	if err == nil {
		t.Errorf("expected `unexpected EOF` but the request was parsed incorrectly!")
	}
}

func TestCorrectMultiRequests(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestCorrectMultiRequests")
	err := ParseRequest(strings.NewReader(httpRequest.text))

	if err != nil {
		t.Errorf("error whlie parsing three consecutive http requets : %q", err)
	}
}

func TestWrongMultiRequests(t *testing.T) {
	httpRequest := filterRequestsByName(requests, "TestWrongMultiRequests")
	err := ParseRequest(strings.NewReader(httpRequest.text))

	if err == nil {
		t.Errorf("expected `unexpected EOF` but the requets where parsed incorrectly")
	}
}
