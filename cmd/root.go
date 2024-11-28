package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "otelkit",
	Short: "A CLI tool for simplifying OpenTelemetry setup and utilities",
	Long: `OtelKit is a comprehensive CLI tool designed to assist with OpenTelemetry workflows.
It provides utilities for validating configurations, generating boilerplate, and more.`,
	// The following line can be uncommented to provide a default action when no subcommand is provided.
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
