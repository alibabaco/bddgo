package bddgo

import (
	"bufio"
	"io"
	"net/http"
)

func ParseRequest(
	reader io.Reader,
	w http.ResponseWriter,
	handler http.Handler) error {

	buf := bufio.NewReader(reader)

	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			break
		}
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

		handler.ServeHTTP(w, req)
	}
	return nil
}
