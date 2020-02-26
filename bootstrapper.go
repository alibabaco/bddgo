package bddgo

import (
	"bufio"
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
