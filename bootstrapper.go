package bddgo

import (
	"bufio"
	"io"
	"net/http"
)

func ServeSingleRequest(
	reader *bufio.Reader,
	writer http.ResponseWriter,
	handler http.Handler) error {

	req, err := http.ReadRequest(reader)
	if err != nil {
		return err
	}

	//		contentLengthStr := req.Header.Get("Content-Length")
	//		if contentLengthStr != "" {
	//			b := new(bytes.Buffer)
	//			io.Copy(b, req.Body)
	//			req.Body.Close()
	//			req.Body = ioutil.NopCloser(b)
	//		}

	handler.ServeHTTP(writer, req)
	return nil
}

func ServeFromReader(
	reader io.Reader,
	writer http.ResponseWriter,
	handler http.Handler) error {

	buf := bufio.NewReader(reader)

	for {
		err := ServeSingleRequest(buf, writer, handler)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
