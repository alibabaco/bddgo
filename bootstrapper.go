package bddgo

import (
	"bufio"
	//"bytes"
	"io"
	//"io/ioutil"
	"net/http"
)

func ParseRequest(
	reader io.Reader,
	w http.ResponseWriter,
	mux *http.ServeMux) error {

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

		mux.ServeHTTP(w, req)
	}
	return nil
}
