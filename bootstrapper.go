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

	handler.ServeHTTP(writer, req)
	return nil
}
