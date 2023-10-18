package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/config"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server/server"
)

func main() {
	port, err := getEnvInt(config.EnvKeyPort, config.DefaultPort)
	if err != nil {
		slog.Error(fmt.Sprintf("error getting port from env: %v", err))
		os.Exit(1)
	}

	s := server.New(port)
	err = s.ListenAndServe()
	if err != nil {
		slog.Error(fmt.Sprintf("error starting the server: %v", err))
		os.Exit(1)
	}
}
