package server

import (
	"net/http"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/types"
)

func registerRoutes(qStorer types.UserRequestStore, handler http.Handler) http.Handler {
	http.HandleFunc("/submit", HandleSubmit(qStorer))
	return handler
}
