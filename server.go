package bddgo

import (
	"net"
	"net/http"
)

func ServeHandler(handler http.Handler, socketFilename string) (err error) {
	server := http.Server{
		Handler: handler,
	}

	unixListener, err := net.Listen("unix", socketFilename)
	if err != nil {
		return
	}
	server.Serve(unixListener)
	return
}
