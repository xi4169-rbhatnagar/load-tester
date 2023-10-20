package server

import (
	"fmt"
	"net/http"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/persistence"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/types"
)

type Server struct {
	http.Server
	qStorer types.UserRequestStore
}

func New(port int) *Server {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	qStorer := persistence.NewRedisStore("127.0.0.1:6379", "", 0)
	h := registerRoutes(qStorer, http.DefaultServeMux)
	return &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}
