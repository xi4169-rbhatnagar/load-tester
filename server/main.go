package main

import (
	"fmt"
	"log/slog"
	"os"
	"server/config"
	"server/server"
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
