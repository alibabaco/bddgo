package bddgo

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func ParseRequest(r io.Reader) error {
	buf := bufio.NewReader(r)

	req, err := http.ReadRequest(buf)

	if err != nil {
		return err
	}

	if contentLengthStr := req.Header.Get("Content-Length"); contentLengthStr != "" {
		b := new(bytes.Buffer)
		io.Copy(b, req.Body)
		req.Body.Close()
		req.Body = ioutil.NopCloser(b)
	}

	return nil
}
