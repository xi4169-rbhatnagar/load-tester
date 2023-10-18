package server

import (
	"fmt"
	"log/slog"

	"github.com/xi4169-rbhatnagar/load-tester/modules/server/config"
)

func Start() error {
	port, err := getEnvInt(config.EnvKeyPort, config.DefaultPort)
	if err != nil {
		slog.Error(fmt.Sprintf("error getting port from env: %v", err))
		return err
	}

	s := New(port)
	slog.Info(fmt.Sprintf("Starting a server on port %d...", port))
	if err = s.ListenAndServe(); err != nil {
		slog.Error(fmt.Sprintf("error starting the server: %v", err))
		return err
	}
	fmt.Println("exitted the server")
	return nil
}
