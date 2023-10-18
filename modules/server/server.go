package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	http.Server
}

func New(port int) *Server {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	h := registerRoutes(http.DefaultServeMux)
	return &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}
