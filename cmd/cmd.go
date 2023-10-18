package main

import (
	"github.com/spf13/cobra"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server"
)

func main() {
	var rootCmd = &cobra.Command{}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start-server",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Start()
		},
	})

	rootCmd.Execute()
}
