package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go run main.go",
	Short:   "di-example-go",
	Example: "go run main.go",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(migrateDBCmd)
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(fmt.Errorf("Error executing: %v\n", err))
	}
}
