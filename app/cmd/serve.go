package cmd

import (
	"fmt"

	"github.com/blockfint/di-example-go/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.InitializeApplication()
		if err != nil {
			panic(fmt.Errorf("Error initializing App: %v\n", err))
		}

		app.Server.Serve()

		return nil
	},
}
